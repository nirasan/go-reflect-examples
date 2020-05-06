package reflect_example

import (
	"reflect"
	"testing"

	. "github.com/nirasan/go-reflect-examples/test_tools"
)

func TestFunction(t *testing.T) {
	v := func(s1, s2 string) string {
		return s1 + " " + s2
	}

	rt := reflect.TypeOf(v)

	Must(t, rt.Kind(), reflect.Func)

	Must(t, rt.NumIn(), 2)

	Must(t, rt.NumOut(), 1)

	rv := reflect.ValueOf(v)

	ret := rv.Call([]reflect.Value{reflect.ValueOf("a"), reflect.ValueOf("b")})
	Must(t, ret[0].String(), "a b")
}
