package reflect_example

import (
	"reflect"
	"testing"
)

func TestType(t *testing.T) {
	t.Run("basic usage", func(t *testing.T) {
		// create Type struct
		value := 100
		rt := reflect.TypeOf(value)

		// get name of type
		if rt.Name() != "int" {
			t.Errorf("invalid name. %v", rt.Name())
		}

		// get kind of type
		if rt.Kind() != reflect.Int {
			t.Errorf("invalid kind. %v", rt.Kind())
		}

		// get size of type
		if rt.Size() != 8 {
			t.Errorf("invalid size. %v", rt.Size())
		}

		// get zero value from Type
		zero := reflect.Zero(rt)
		if zero.Int() != 0 {
			t.Errorf("zero value must 0. %v", zero)
		}
	})

	t.Run("panic if call miss match method", func(t *testing.T) {
		value := 100
		rt := reflect.TypeOf(value)
		defer func() {
			if err := recover(); err == nil {
				t.Errorf("error should have occurred")
			}
		}()
		// NumField occur panics if Kind is not Struct
		rt.NumField()
	})
}
