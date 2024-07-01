package goinject

import (
	"reflect"
)

// DIContainer declares what must be implemented by the struct type to be
// used by the global functions.
type DIContainer interface {
	RegisterType(abstractType reflect.Type, concreteType reflect.Type)
	Register(abstractType reflect.Type, concreteInstance any)
	Inject(abstractType reflect.Type) any
}

// InitializableDependency declares the [InitializableDependency.Initialize]
// contract that can be called during the registration process by the [BaseContainer].
type InitializableDependency interface {
	Initialize()
}
