package reflect_example

import (
	"reflect"
	"testing"
)

type Type1 struct{}

type Type2 struct{}

func (t *Type1) Func1() string {
	return "func1"
}

type Interface1 interface {
	Func1() string
}

func TestInterface(t *testing.T) {
	i := reflect.TypeOf((*Interface1)(nil)).Elem()

	t1 := reflect.TypeOf(&Type1{})
	if t1.Implements(i) != true {
		t.Errorf("Type1 must implements Interface1")
	}

	t2 := reflect.TypeOf(&Type2{})
	if t2.Implements(i) != false {
		t.Errorf("Type1 must not implements Interface1")
	}
}
