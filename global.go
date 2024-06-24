package goinject

import (
	"reflect"
)

var DefaultContainer DIContainer = NewBaseContainer()

func RegisterType[Abstract any, Concrete InitializableDependency]() {
	DefaultContainer.RegisterType(
		reflect.TypeFor[Abstract](),
		reflect.TypeFor[Concrete](),
	)
}

func Register[Abstract any](obj Abstract) {
	DefaultContainer.Register(
		reflect.TypeFor[Abstract](),
		obj,
	)
}

func Inject[Abstract any](obj *Abstract) {
	*obj = DefaultContainer.Inject(reflect.TypeFor[Abstract]()).(Abstract)
}
