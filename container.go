package goinject

import (
	"reflect"
)

// DIContainer declares what must be implemented by the struct type to be
// used by the global functions.
type DIContainer interface {
	// RegisterType of an abstract type to a concrete type inside the DI
	// container, to be injected later.
	//
	// The Abstract type must be an interface, and the Concrete type must be a
	// struct type that implements the interface. Anything different from this
	// must panic.
	RegisterType(abstractType reflect.Type, concreteType reflect.Type)

	// Register an abstract type to a concrete instance inside the DI
	// container, to be injected later.
	//
	// The Abstract type must be an interface, and the object instance must be
	// the type of a struct that implements the interface. Anything different
	// from this must panic.
	Register(abstractType reflect.Type, concreteInstance any)

	// Inject the instance of the registered Concrete type from the DI container.
	//
	// The Abstract type must be an interface.
	//
	// The Concrete type must be instantiated if it isn't already. The
	// [InitializableDependency.InitializeDependency] method must be be called
	// if the Concrete type implements the [InitializableDependency] interface.
	Inject(abstractType reflect.Type) any
}

// InitializableDependency declares the
// [InitializableDependency.InitializeDependency] contract that can be called
// during the registration process by the [BaseContainer].
type InitializableDependency interface {
	// InitializeDependency after the instance is created, from the
	// [DIContainer.Inject] method.
	InitializeDependency()
}
