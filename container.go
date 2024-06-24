package goinject

import (
	"reflect"
)

type DIContainer interface {
	RegisterType(abstractType reflect.Type, concreteType reflect.Type)
	Register(abstractType reflect.Type, concreteInstance any)
	Inject(abstractType reflect.Type) any
}

type InitializableDependency interface {
	Initialize()
}
