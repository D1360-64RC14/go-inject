package goinject

import "errors"

var (
	ErrNotAnInterface          = errors.New("goinject: abstract type must be an Interface")
	ErrNotAnStruct             = errors.New("goinject: concrete type must be an Struct or *Struct")
	ErrInterfaceNotImplemented = errors.New("goinject: concrete type must implement abstract type")
	ErrAlreadyRegistered       = errors.New("goinject: there's already a relation for abstract type")
	ErrNoConcreteTypeSupplied  = errors.New("goinject: there's no concrete type supplied for abstract type")
)
