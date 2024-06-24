package goinject_test

import (
	"errors"
	"reflect"
	"testing"

	goinject "github.com/d1360-64rc14/go-inject"
)

type TestA interface {
	TestA()
}
type TestB interface {
	TestB()
}
type TestC interface {
	TestC()
}
type TestD interface {
	TestD()
}
type TestE interface {
	TestE()
}

type TestAImpl struct{}
type TestBImpl struct{}
type TestCImpl struct {
	Executed bool
}
type TestDImpl struct{}
type TestEImpl struct{}

func (a *TestAImpl) Initialize() {}
func (a *TestAImpl) TestA()      {}
func (b *TestBImpl) TestB()      {}
func (a *TestCImpl) Initialize() {
	a.Executed = true
}
func (a *TestCImpl) TestC()      {}
func (a *TestDImpl) Initialize() {}
func (a *TestDImpl) TestD()      {}
func (a *TestEImpl) Initialize() {}
func (a *TestEImpl) TestE()      {}
func (a *TestEImpl) TestA()      {}

func TestBaseInjector(t *testing.T) {
	t.Run("RegisterType", func(t *testing.T) {
		t.Run("Normal execution", func(t *testing.T) {
			t.Parallel()

			i := goinject.NewBaseContainer()

			err := recoverPanic(func() {
				i.RegisterType(
					reflect.TypeFor[TestA](),
					reflect.TypeFor[*TestAImpl](),
				)
			})
			if err != nil {
				t.Errorf("unexpected error: '%v'", err)
				return
			}

			var test TestA

			err = recoverPanic(func() {
				test = i.Inject(reflect.TypeFor[TestA]()).(*TestAImpl)
			})
			if err != nil {
				t.Errorf("unexpected error: '%v'", err)
				return
			}

			if test == nil {
				t.Error("test shouldn't be nil")
				return
			}

			err = recoverPanic(func() {
				test.TestA()
			})
			if err != nil {
				t.Errorf("unexpected error: '%v'", err)
				return
			}
		})
		t.Run("Multiple registrations", func(t *testing.T) {
			t.Parallel()

			i := goinject.NewBaseContainer()

			err := recoverPanic(func() {
				i.RegisterType(
					reflect.TypeFor[TestC](),
					reflect.TypeFor[*TestCImpl](),
				)
				i.RegisterType(
					reflect.TypeFor[TestD](),
					reflect.TypeFor[*TestDImpl](),
				)
				i.RegisterType(
					reflect.TypeFor[TestE](),
					reflect.TypeFor[*TestEImpl](),
				)
			})
			if err != nil {
				t.Errorf("unexpected error: '%v'", err)
				return
			}

			var testC TestC
			var testD TestD
			var testE TestE

			err = recoverPanic(func() {
				testC = i.Inject(reflect.TypeFor[TestC]()).(TestC)
				testD = i.Inject(reflect.TypeFor[TestD]()).(TestD)
				testE = i.Inject(reflect.TypeFor[TestE]()).(TestE)
			})
			if err != nil {
				t.Errorf("unexpected error: '%v'", err)
				return
			}

			if testC == nil {
				t.Error("testC shouldn't be nil")
				return
			}
			if testD == nil {
				t.Error("testD shouldn't be nil")
				return
			}
			if testE == nil {
				t.Error("testE shouldn't be nil")
				return
			}
		})
	})

	t.Run("Calling Initialize method", func(t *testing.T) {
		t.Parallel()

		i := goinject.NewBaseContainer()

		err := recoverPanic(func() {
			i.RegisterType(
				reflect.TypeFor[TestC](),
				reflect.TypeFor[*TestCImpl](),
			)
		})
		if err != nil {
			t.Errorf("unexpected error: '%v'", err)
			return
		}

		var testC *TestCImpl

		err = recoverPanic(func() {
			testC = i.Inject(reflect.TypeFor[TestC]()).(*TestCImpl)
		})
		if err != nil {
			t.Errorf("unexpected error: '%v'", err)
			return
		}

		if testC == nil {
			t.Error("testC shouldn't be nil")
			return
		}

		if !testC.Executed {
			t.Error("method Initialize was not called")
			return
		}
	})

	t.Run("Register", func(t *testing.T) {
		t.Run("Normal execution", func(t *testing.T) {
			t.Parallel()

			i := goinject.NewBaseContainer()

			{
				sentTestAInstance := &TestAImpl{}

				err := recoverPanic(func() {
					i.Register(
						reflect.TypeFor[TestA](),
						sentTestAInstance,
					)
				})

				if err != nil {
					t.Errorf("unexpected error: '%v'", err)
					return
				}
			}

			{
				var gotTestAInstance TestA

				err := recoverPanic(func() {
					gotTestAInstance = i.Inject(reflect.TypeFor[TestA]()).(TestA)
				})

				if err != nil {
					t.Errorf("unexpected error: '%v'", err)
					return
				}

				err = recoverPanic(func() {
					gotTestAInstance.TestA()
				})

				if err != nil {
					t.Errorf("unexpected error: '%v'", err)
					return
				}
			}
		})
		t.Run("Multiple registrations", func(t *testing.T) {
			t.Parallel()

			i := goinject.NewBaseContainer()

			err := recoverPanic(func() {
				i.Register(reflect.TypeFor[TestC](), &TestCImpl{})
				i.Register(reflect.TypeFor[TestD](), &TestDImpl{})
				i.Register(reflect.TypeFor[TestE](), &TestEImpl{})
			})
			if err != nil {
				t.Errorf("unexpected error: '%v'", err)
				return
			}

			var testC TestC
			var testD TestD
			var testE TestE

			err = recoverPanic(func() {
				testC = i.Inject(reflect.TypeFor[TestC]()).(*TestCImpl)
				testD = i.Inject(reflect.TypeFor[TestD]()).(*TestDImpl)
				testE = i.Inject(reflect.TypeFor[TestE]()).(*TestEImpl)
			})
			if err != nil {
				t.Errorf("unexpected error: '%v'", err)
				return
			}

			if testC == nil {
				t.Error("testC shouldn't be nil")
				return
			}
			if testD == nil {
				t.Error("testD shouldn't be nil")
				return
			}
			if testE == nil {
				t.Error("testE shouldn't be nil")
				return
			}
		})
	})
}

func TestBaseInjectorRegisterErrors(t *testing.T) {
	testCases := []struct {
		desc             string
		abstractType     reflect.Type
		concreteInstance any
		err              error
	}{
		{
			desc:             "Right abstract and concrete types",
			abstractType:     reflect.TypeFor[TestA](),
			concreteInstance: &TestAImpl{},
			err:              nil,
		},
		{
			desc:             "Nil abstract type",
			abstractType:     nil,
			concreteInstance: &TestAImpl{},
			err:              goinject.ErrNotAnInterface,
		},
		{
			desc:             "Nil concrete type",
			abstractType:     reflect.TypeFor[TestA](),
			concreteInstance: nil,
			err:              goinject.ErrNotAnStruct,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			i := goinject.NewBaseContainer()

			err := recoverPanic(func() {
				i.Register(tC.abstractType, tC.concreteInstance)
			})

			if !errors.Is(err, tC.err) {
				t.Errorf("expected error '%v', got '%v'", tC.err, err)
			}
		})
	}
}

func TestBaseInjectorRegisterTypeErrors(t *testing.T) {
	testCases := []struct {
		desc         string
		abstractType reflect.Type
		concreteType reflect.Type
		err          error
	}{
		{
			desc:         "Right abstract and concrete types",
			abstractType: reflect.TypeFor[TestA](),
			concreteType: reflect.TypeFor[*TestAImpl](),
			err:          nil,
		},
		{
			desc:         "Wrong abstract type",
			abstractType: reflect.TypeFor[TestB](),
			concreteType: reflect.TypeFor[*TestAImpl](),
			err:          goinject.ErrInterfaceNotImplemented,
		},
		{
			desc:         "Abstract type is not an interface",
			abstractType: reflect.TypeFor[TestAImpl](),
			concreteType: reflect.TypeFor[*TestAImpl](),
			err:          goinject.ErrNotAnInterface,
		},
		{
			desc:         "Concrete type is not an Struct",
			abstractType: reflect.TypeFor[TestA](),
			concreteType: reflect.TypeFor[TestA](),
			err:          goinject.ErrNotAnStruct,
		},
		{
			desc:         "Concrete type is not an *Struct",
			abstractType: reflect.TypeFor[TestA](),
			concreteType: reflect.TypeFor[*TestA](),
			err:          goinject.ErrNotAnStruct,
		},
		{
			desc:         "Nil abstract type",
			abstractType: nil,
			concreteType: reflect.TypeFor[*TestA](),
			err:          goinject.ErrNotAnInterface,
		},
		{
			desc:         "Nil concrete type",
			abstractType: reflect.TypeFor[TestA](),
			concreteType: nil,
			err:          goinject.ErrNotAnStruct,
		},
		{
			desc:         "Nil abstract and concrete type",
			abstractType: nil,
			concreteType: nil,
			err:          goinject.ErrNotAnInterface,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			i := goinject.NewBaseContainer()

			err := recoverPanic(func() {
				i.RegisterType(tC.abstractType, tC.concreteType)
			})

			if !errors.Is(err, tC.err) {
				t.Errorf("expected error '%v', got '%v'", tC.err, err)
			}
		})
	}
}

func TestBaseInjectorRegisterTypeShared(t *testing.T) {
	testCases := []struct {
		desc         string
		abstractType reflect.Type
		concreteType reflect.Type
		err          error
	}{
		{
			desc:         "Right abstract and concrete types",
			abstractType: reflect.TypeFor[TestA](),
			concreteType: reflect.TypeFor[*TestAImpl](),
			err:          nil,
		},
		{
			desc:         "Relation already supplied",
			abstractType: reflect.TypeFor[TestA](),
			concreteType: reflect.TypeFor[*TestEImpl](),
			err:          goinject.ErrAlreadyRegistered,
		},
	}

	i := goinject.NewBaseContainer()

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := recoverPanic(func() {
				i.RegisterType(tC.abstractType, tC.concreteType)
			})

			if !errors.Is(err, tC.err) {
				t.Errorf("expected error '%v', got '%v'", tC.err, err)
			}
		})
	}
}

func TestBaseInjectorInject(t *testing.T) {
	i := goinject.NewBaseContainer()

	i.RegisterType(
		reflect.TypeFor[TestA](),
		reflect.TypeFor[*TestAImpl](),
	)

	testCases := []struct {
		desc         string
		abstractType reflect.Type
		err          error
	}{
		{
			desc:         "Existent interface",
			abstractType: reflect.TypeFor[TestA](),
			err:          nil,
		},
		{
			desc:         "Nonexistent interface",
			abstractType: reflect.TypeFor[TestB](),
			err:          goinject.ErrNoConcreteTypeSupplied,
		},
		{
			desc:         "Non interface type",
			abstractType: reflect.TypeFor[*TestAImpl](),
			err:          goinject.ErrNotAnInterface,
		},
		{
			desc:         "Nil type",
			abstractType: nil,
			err:          goinject.ErrNotAnInterface,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			err := recoverPanic(func() {
				i.Inject(tC.abstractType)
			})

			if !errors.Is(err, tC.err) {
				t.Errorf("expected error '%v', got '%v'", tC.err, err)
			}
		})
	}
}

func recoverPanic(f func()) (e error) {
	defer func() {
		if r := recover(); r != nil {
			e = r.(error)
		}
	}()

	f()

	return
}
