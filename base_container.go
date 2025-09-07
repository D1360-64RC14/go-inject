package goinject

import (
	"fmt"
	"reflect"
	"sync"
)

type abstractType reflect.Type
type concreteType reflect.Type

type BaseContainer struct {
	relations map[abstractType]concreteType
	instances map[abstractType]any

	mx sync.Mutex
}

func NewBaseContainer() *BaseContainer {
	return &BaseContainer{
		relations: make(map[abstractType]concreteType),
		instances: make(map[abstractType]any),
	}
}

func (i *BaseContainer) Register(abstractType reflect.Type, concreteInstance any) {
	i.RegisterType(abstractType, reflect.TypeOf(concreteInstance))

	i.mx.Lock()
	defer i.mx.Unlock()

	i.instances[abstractType] = concreteInstance
}

func (i *BaseContainer) RegisterType(abstractType reflect.Type, concreteType reflect.Type) {
	if abstractType == nil || abstractType.Kind() != reflect.Interface {
		panic(ErrNotAnInterface)
	}

	if concreteType == nil {
		panic(ErrNotAnStruct)
	}

	isConcreteStruct := concreteType.Kind() == reflect.Struct
	isConcreteStructPtr := concreteType.Kind() == reflect.Pointer && concreteType.Elem().Kind() == reflect.Struct

	if !isConcreteStruct && !isConcreteStructPtr {
		panic(ErrNotAnStruct)
	}

	if !concreteType.Implements(abstractType) {
		panic(fmt.Errorf("%w: %s (concrete type), %s (abstract type)", ErrInterfaceNotImplemented, concreteType.Name(), abstractType.Name()))
	}

	i.mx.Lock()
	defer i.mx.Unlock()

	if _, ok := i.relations[abstractType]; ok {
		panic(fmt.Errorf("%w: %s (abstract type)", ErrAlreadyRegistered, abstractType.Name()))
	}

	if concreteType.Kind() == reflect.Pointer {
		concreteType = concreteType.Elem()
	}

	i.relations[abstractType] = concreteType
}

func (i *BaseContainer) Inject(abstractType reflect.Type) any {
	if abstractType == nil || abstractType.Kind() != reflect.Interface {
		panic(ErrNotAnInterface)
	}

	i.mx.Lock()
	defer i.mx.Unlock()

	concreteType, ok := i.relations[abstractType]
	if !ok {
		panic(fmt.Errorf("%w: %s (abstract type)", ErrNoConcreteTypeSupplied, abstractType.Name()))
	}

	instance, ok := i.instances[abstractType]
	if !ok {
		instance = reflect.New(concreteType).Interface()
		i.instances[abstractType] = instance

		if dInstance, ok := instance.(InitializableDependency); ok {
			dInstance.InitializeDependency()
		}
	}

	return instance
}
