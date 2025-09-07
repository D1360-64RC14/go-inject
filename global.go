package goinject

import (
	"reflect"
)

// DefaultContainer holds the container used when calling the global functions.
// It can be reassigned to a different container if needed.
var DefaultContainer DIContainer = NewBaseContainer()

// RegisterType of an abstract type to a concrete type inside the DI container,
// to be injected later.
//
// The Abstract type must be an interface, and the Concrete type must be a
// struct type that implements the interface.
//
//	type BookRepository interface {
//		Get() []int
//		Save(string) int
//	}
//
//	type MySQLBookRepository struct {}
//
//	goinject.RegisterType[BookRepository, MySQLBookRepository]()
//
// The Concrete type will be instantiated when injected if it isn't already.
// If the Concrete type implements the [InitializableDependency] interface, the
// [InitializableDependency.InitializeDependency] method will be called to
// instantiate it.
func RegisterType[Abstract any, Concrete any]() {
	DefaultContainer.RegisterType(
		reflect.TypeFor[Abstract](),
		reflect.TypeFor[Concrete](),
	)
}

// Register an abstract type to a concrete instance inside the DI container,
// to be injected later.
//
// The Abstract type must be an interface, and the object instance must be the
// type of a struct that implements the interface.
//
// Always specify the Abstract type, or a struct type will be inferred,
// resulting in a panic.
//
//	type BookRepository interface {
//		Get() []int
//		Save(string) int
//	}
//
//	bookRepo := repository.NewMySQLBookRepository()
//
//	goinject.Register[BookRepository](bookRepo)
func Register[Abstract any](obj Abstract) {
	DefaultContainer.Register(
		reflect.TypeFor[Abstract](),
		obj,
	)
}

// Inject the instance of some pre-registered Concrete type from the DI container.
//
// The Concrete type will be instantiated if it isn't already. Or the
// [InitializableDependency.InitializeDependency] method will be called if it
// implements the [InitializableDependency] interface.
//
//	var bookRepo BookRepository = goinject.Inject[BookRepository](&bookRepo)
func Inject[Abstract any]() Abstract {
	return DefaultContainer.Inject(reflect.TypeFor[Abstract]()).(Abstract)
}

// InjectAt the given variable reference the instance of some pre-registered
// Concrete type from the DI container.
//
// The Concrete type will be instantiated if it isn't already. Or the
// [InitializableDependency.InitializeDependency] method will be called if it
// implements the [InitializableDependency] interface.
//
//	var bookRepo BookRepository
//
//	goinject.InjectAt(&bookRepo)
func InjectAt[Abstract any](obj *Abstract) {
	*obj = DefaultContainer.Inject(reflect.TypeFor[Abstract]()).(Abstract)
}
