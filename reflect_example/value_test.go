package reflect_example

import (
	"reflect"
	"testing"
)

func TestValue(t *testing.T) {
	t.Run("basic usage", func(t *testing.T) {
		// create int value
		value := 100

		// create Value struct
		rv := reflect.ValueOf(value)

		// get real value as int64
		if rv.Int() != 100 {
			t.Errorf("invalid value. %v", rv.Int())
		}

		// get real value as interface
		i := rv.Interface()
		if v, ok := i.(int); !ok || v != 100 {
			t.Errorf("invalid value. %v, %v", v, ok)
		}

		// get Type from Value
		if rv.Type().Name() != "int" {
			t.Errorf("invalid value. %v, %v", rv, rv.Type())
		}

		// get Kind from Value
		if rv.Kind() != reflect.Int {
			t.Errorf("invalid value. %v, %v", rv, rv.Kind())
		}
	})

	t.Run("zero value check", func(t *testing.T) {
		// zero value check
		var v int
		if !reflect.ValueOf(v).IsZero() {
			t.Errorf("value must zero. %v", v)
		}
		v = 1
		if reflect.ValueOf(v).IsZero() {
			t.Errorf("value must not zero. %v", v)
		}
	})

	t.Run("nil check", func(t *testing.T) {
		// nil check
		rv := reflect.ValueOf((*int)(nil))
		if !rv.IsNil() {
			t.Error("must nil")
		}
	})

	t.Run("update", func(t *testing.T) {
		// create int value
		value := 100

		// create Value struct from *int
		rv := reflect.ValueOf(&value)
		// value is pointer and value is not settable
		if rv.Kind() != reflect.Ptr || rv.CanSet() == true {
			t.Errorf("invalid value. %v", rv)
		}

		// get int value from *int
		rv = rv.Elem()
		// value is int and value is settable
		if rv.Kind() != reflect.Int || rv.CanSet() != true {
			t.Errorf("invalid value. %v", rv)
		}

		// set new value
		rv.SetInt(200)
		if value != 200 {
			t.Errorf("invalid value. %v", value)
		}
	})
}
