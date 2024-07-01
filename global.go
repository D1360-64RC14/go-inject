package goinject

import (
	"reflect"
)

var DefaultContainer DIContainer = NewBaseContainer()

// RegisterType of an abstract type to a concrete type inside the DI container
// that can be injected later.
//
// The Abstract type must be an interface, and the Concrete type must be a struct
// type that must implement the interface.
//
//	type BookRepository interface {
//		Get() []int
//		Save(string) int
//	}
//
//	type MySQLBookRepository struct {}
//
//	goinject.RegisterType[BookRepository, MySQLBookRepository]()
func RegisterType[Abstract any, Concrete any]() {
	DefaultContainer.RegisterType(
		reflect.TypeFor[Abstract](),
		reflect.TypeFor[Concrete](),
	)
}

// Register an abstract type to a concrete instance inside the DI container that
// can be injected later.
//
// The Abstract type must be an interface, and the object instance must be the
// type of a struct that implements the interface.
//
// Always specify the Abstract type, or the Concrete type will be inferred, and
// a panic will occur.
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

// Inject into the given variable reference the instance of some pre-registered
// Concrete type or instance from the DI container.
//
// If the Concrete type isn't instantiated yet, it will be instantiated. Or if
// it implements the [InitializableDependency] interface, the
// [InitializableDependency.Initialize] method will be called.
//
//	var bookRepo BookRepository
//
//	goinject.Inject(&bookRepo)
func Inject[Abstract any](obj *Abstract) {
	*obj = DefaultContainer.Inject(reflect.TypeFor[Abstract]()).(Abstract)
}
