package reflect_example

import (
	"reflect"
	"testing"
)

func TestBool(t *testing.T) {
	var v bool

	// create Type
	rt := reflect.TypeOf(v)

	// kind is Bool
	if rt.Kind() != reflect.Bool {
		t.Errorf("invalid kind. %v", rt.Kind())
	}

	// create Value
	rv := reflect.ValueOf(v)

	// get bool from Value
	if rv.Bool() != false {
		t.Errorf("value must false. %v", rv)
	}

	// update
	rv = reflect.ValueOf(&v).Elem()
	rv.SetBool(true)
	if v != true {
		t.Errorf("failed to update. %v", v)
	}
}
