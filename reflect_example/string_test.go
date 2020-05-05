package reflect_example

import (
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	v := "hello world"

	rt := reflect.TypeOf(v)

	// get Kind
	if rt.Kind() != reflect.String {
		t.Errorf("invalid kind. %v", rt.Kind())
	}

	rv := reflect.ValueOf(&v).Elem()

	// get Kind
	if rv.Kind() != reflect.String {
		t.Errorf("invalid kind. %v", rv.Kind())
	}

	// get underlying value
	if rv.String() != "hello world" {
		t.Errorf("invalid string. %v", rv.String())
	}

	// get length
	if rv.Len() != 11 {
		t.Errorf("invalid length. %v", rv.Len())
	}

	// update underlying value
	rv.SetString("HELLO WORLD")
	if v != "HELLO WORLD" {
		t.Errorf("failed to update. %v", v)
	}
}
