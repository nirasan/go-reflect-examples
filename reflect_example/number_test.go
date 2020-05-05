package reflect_example

import (
	"reflect"
	"testing"
)

func TestNumberKind(t *testing.T) {
	t.Run("Kind", func(t *testing.T) {
		samples := []struct {
			Value interface{}
			Kind  reflect.Kind
		}{
			{int(1), reflect.Int},
			{int8(1), reflect.Int8},
			{int16(1), reflect.Int16},
			{int32(1), reflect.Int32},
			{int64(1), reflect.Int64},
			{uint(1), reflect.Uint},
			{uint8(1), reflect.Uint8},
			{uint16(1), reflect.Uint16},
			{uint32(1), reflect.Uint32},
			{uint64(1), reflect.Uint64},
			{float32(1), reflect.Float32},
			{float64(1), reflect.Float64},
			{complex64(1), reflect.Complex64},
			{complex128(1), reflect.Complex128},
		}
		for _, s := range samples {
			rv := reflect.ValueOf(s.Value)
			if rv.Kind() != s.Kind {
				t.Errorf("invalid sample. %+v", s)
			}
		}
	})

	t.Run("int", func(t *testing.T) {
		v := 100
		rt := reflect.TypeOf(v)
		if rt.Kind() != reflect.Int {
			t.Errorf("invalid kind. %v", rt.Kind())
		}
		rv := reflect.ValueOf(&v).Elem()
		if rv.Int() != 100 {
			t.Errorf("value must 100. %v", rv)
		}
		rv.SetInt(200)
		if v != 200 {
			t.Errorf("value must updated. %v", v)
		}
	})

	t.Run("uint", func(t *testing.T) {
		v := uint(100)
		rt := reflect.TypeOf(v)
		if rt.Kind() != reflect.Uint {
			t.Errorf("invalid kind. %v", rt.Kind())
		}
		rv := reflect.ValueOf(&v).Elem()
		if rv.Uint() != 100 {
			t.Errorf("value must 100. %v", rv)
		}
		rv.SetUint(200)
		if v != 200 {
			t.Errorf("value must updated. %v", v)
		}
	})

	t.Run("float", func(t *testing.T) {
		v := float64(100)
		rt := reflect.TypeOf(v)
		if rt.Kind() != reflect.Float64 {
			t.Errorf("invalid kind. %v", rt.Kind())
		}
		rv := reflect.ValueOf(&v).Elem()
		if rv.Float() != 100 {
			t.Errorf("value must 100. %v", rv)
		}
		rv.SetFloat(200)
		if v != 200 {
			t.Errorf("value must updated. %v", v)
		}
	})

	t.Run("complex", func(t *testing.T) {
		v := complex64(100)
		rt := reflect.TypeOf(v)
		if rt.Kind() != reflect.Complex64 {
			t.Errorf("invalid kind. %v", rt.Kind())
		}
		rv := reflect.ValueOf(&v).Elem()
		if rv.Complex() != 100 {
			t.Errorf("value must 100. %v", rv)
		}
		rv.SetComplex(200)
		if v != 200 {
			t.Errorf("value must updated. %v", v)
		}
	})
}
