package reflect_example

import (
	"reflect"
	"testing"
)

func TestArray(t *testing.T) {
	v := [3]int{1, 2, 3}

	t.Run("basic usage", func(t *testing.T) {
		// get Type
		rt := reflect.TypeOf(v)
		if rt.Kind() != reflect.Array {
			t.Errorf("invalid kind. %v", rt.Kind())
		}

		// get element Type
		rte := rt.Elem()
		if rte.Kind() != reflect.Int {
			t.Errorf("invalid kind. %v", rte.Kind())
		}

		// get Value
		rv := reflect.ValueOf(v)

		// get len and cap
		if rv.Len() != 3 || rv.Cap() != 3 {
			t.Errorf("invalid len or cap. %v, %v, %v", rv.Len(), rv.Cap(), rv)
		}

		// get element
		if rv.Index(0).Int() != 1 {
			t.Errorf("invalid element. %v", rv.Index(0))
		}
	})

	t.Run("update", func(t *testing.T) {
		v := []string{"hello", "world"}
		rv := reflect.ValueOf(v)
		rv.Index(0).SetString("HELLO")
		if v[0] != "HELLO" {
			t.Errorf("failed to update. %+v", v)
		}
	})

	t.Run("loop", func(t *testing.T) {
		rv := reflect.ValueOf(v)
		sum := 0
		for i := 0; i < rv.Len(); i++ {
			sum += int(rv.Index(i).Int())
		}
		if sum != 6 {
			t.Errorf("sum must 6. %v", sum)
		}
	})

	t.Run("slice of array", func(t *testing.T) {
		rv := reflect.ValueOf(&v).Elem()
		rv2 := rv.Slice(0, 2)
		if rv2.Len() != 2 || rv2.Cap() != 3 || rv2.Index(0).Int() != 1 || rv2.Index(1).Int() != 2 {
			t.Errorf("failed to slice. %v, %v, %v", rv2.Len(), rv2.Cap(), rv2)
		}
	})
}
