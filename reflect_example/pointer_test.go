package reflect_example

import (
	"reflect"
	"testing"
)

func TestPinter(t *testing.T) {
	v := 100

	rt := reflect.TypeOf(&v)

	// get Kind
	Must(t, rt.Kind(), reflect.Ptr)

	// get element Kind
	Must(t, rt.Elem().Kind(), reflect.Int)

	rv := reflect.ValueOf(&v)

	// get underlying value
	Must(t, rv.Elem().Int(), int64(100))

	// Indirect returns pointer element if Kind is pointer
	Must(t, rv.Elem(), reflect.Indirect(rv))
	Must(t, rv.Elem(), reflect.Indirect(reflect.Indirect(rv)))

	// nil check
	rv = reflect.ValueOf((*int)(nil))
	Must(t, rv.IsNil(), true)
}
