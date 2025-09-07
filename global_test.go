package goinject_test

import (
	"errors"
	"testing"

	goinject "github.com/d1360-64rc14/go-inject"
)

func TestRegisterType(t *testing.T) {
	t.Run("Normal execution", func(t *testing.T) {
		err := recoverPanic(func() {
			goinject.RegisterType[TestA, *TestAImpl]()
		})

		if err != nil {
			t.Errorf("unexpected error: '%v'", err)
		}
	})

	t.Run("Already registered", func(t *testing.T) {
		err := recoverPanic(func() {
			goinject.RegisterType[TestA, *TestAImpl]()
		})

		if !errors.Is(err, goinject.ErrAlreadyRegistered) {
			t.Errorf("expected error: '%v', got '%v'", goinject.ErrAlreadyRegistered, err)
		}
	})

	t.Run("Wrong abstract type", func(t *testing.T) {
		err := recoverPanic(func() {
			goinject.RegisterType[string, *TestBImpl]()
		})

		if !errors.Is(err, goinject.ErrNotAnInterface) {
			t.Errorf("expected error: '%v', got '%v'", goinject.ErrNotAnInterface, err)
		}
	})

	t.Run("Wrong concrete type", func(t *testing.T) {
		err := recoverPanic(func() {
			goinject.RegisterType[TestC, string]()
		})

		if !errors.Is(err, goinject.ErrNotAnStruct) {
			t.Errorf("expected error: '%v', got '%v'", goinject.ErrNotAnStruct, err)
		}
	})
}

func TestRegister(t *testing.T) {
	t.Run("Normal execution", func(t *testing.T) {
		inst := &TestCImpl{}

		err := recoverPanic(func() {
			goinject.Register[TestC](inst)
		})

		if err != nil {
			t.Errorf("unexpected error: '%v'", err)
		}
	})

	t.Run("Already registered", func(t *testing.T) {
		inst := &TestCImpl{}

		err := recoverPanic(func() {
			goinject.Register[TestC](inst)
		})

		if !errors.Is(err, goinject.ErrAlreadyRegistered) {
			t.Errorf("expected error: '%v', got '%v'", goinject.ErrAlreadyRegistered, err)
		}
	})

	t.Run("Wrong abstract type", func(t *testing.T) {
		inst := ""

		err := recoverPanic(func() {
			goinject.Register[string](inst)
		})

		if !errors.Is(err, goinject.ErrNotAnInterface) {
			t.Errorf("expected error: '%v', got '%v'", goinject.ErrNotAnInterface, err)
		}
	})
}

func TestInject(t *testing.T) {
	t.Run("Normal execution", func(t *testing.T) {
		err := recoverPanic(func() {
			_ = goinject.Inject[TestA]()
		})

		if err != nil {
			t.Errorf("unexpected error: '%v'", err)
		}
	})

	t.Run("Wrong abstract type", func(t *testing.T) {
		err := recoverPanic(func() {
			_ = goinject.Inject[*TestAImpl]()
		})

		if !errors.Is(err, goinject.ErrNotAnInterface) {
			t.Errorf("expected error: '%v', got '%v'", goinject.ErrNotAnInterface, err)
		}
	})

	t.Run("Injected object", func(t *testing.T) {
		var inst TestA

		err := recoverPanic(func() {
			inst = goinject.Inject[TestA]()
		})

		if inst == nil {
			t.Errorf("expected instance, got nil")
		}

		if err != nil {
			t.Errorf("unexpected error: '%v'", err)
		}
	})

	t.Run("Not registered type", func(t *testing.T) {
		err := recoverPanic(func() {
			_ = goinject.Inject[TestE]()
		})

		if !errors.Is(err, goinject.ErrNoConcreteTypeSupplied) {
			t.Errorf("expected error: '%v', got '%v'", goinject.ErrNoConcreteTypeSupplied, err)
		}
	})
}

func TestInjectAt(t *testing.T) {
	t.Run("Normal execution", func(t *testing.T) {
		var inst TestA

		err := recoverPanic(func() {
			goinject.InjectAt(&inst)
		})

		if err != nil {
			t.Errorf("unexpected error: '%v'", err)
		}
	})

	t.Run("Wrong abstract type", func(t *testing.T) {
		var inst *TestAImpl

		err := recoverPanic(func() {
			goinject.InjectAt(&inst)
		})

		if !errors.Is(err, goinject.ErrNotAnInterface) {
			t.Errorf("expected error: '%v', got '%v'", goinject.ErrNotAnInterface, err)
		}
	})

	t.Run("Injected object", func(t *testing.T) {
		var inst TestA

		err := recoverPanic(func() {
			goinject.InjectAt(&inst)
		})

		if inst == nil {
			t.Errorf("expected instance, got nil")
		}

		if err != nil {
			t.Errorf("unexpected error: '%v'", err)
		}
	})

	t.Run("Not registered type", func(t *testing.T) {
		var inst TestE

		err := recoverPanic(func() {
			goinject.InjectAt(&inst)
		})

		if !errors.Is(err, goinject.ErrNoConcreteTypeSupplied) {
			t.Errorf("expected error: '%v', got '%v'", goinject.ErrNoConcreteTypeSupplied, err)
		}
	})
}
