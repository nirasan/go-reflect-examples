package reflect_example

import (
	"fmt"
	"reflect"
	"testing"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

func (p *Person) SetAge(age int) {
	p.Age = age
}

func TestStruct(t *testing.T) {
	v := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       20,
		Gender:    "Male",
	}
	t.Run("kind", func(t *testing.T) {
		rt := reflect.TypeOf(v)
		if rt.Kind() != reflect.Struct {
			t.Errorf("invalid kind. %v", rt.Kind())
		}
	})

	t.Run("field type", func(t *testing.T) {
		rt := reflect.TypeOf(v)
		if rt.NumField() != 4 {
			t.Errorf("NumField must 4. %v", rt.NumField())
		}
		f := rt.Field(0)
		if f.Name != "FirstName" || f.Type.Kind() != reflect.String {
			t.Errorf("invalid field. %v", f)
		}
		tag := f.Tag
		if tag.Get("json") != "first_name" {
			t.Errorf("invalid tag. %v", tag)
		}
		if f, ok := rt.FieldByName("Age"); ok {
			if f.Name != "Age" || f.Type.Kind() != reflect.Int {
				t.Errorf("invalid field. %v", f)
			}
		}
	})

	t.Run("field value", func(t *testing.T) {
		rv := reflect.ValueOf(&v).Elem()
		if rv.NumField() != 4 {
			t.Errorf("NumField must 4. %v", rv.NumField())
		}
		f := rv.Field(0)
		if f.String() != "John" {
			t.Errorf("invalid field. %v", f)
		}
		f.SetString("JOHN")
		if v.FirstName != "JOHN" {
			t.Errorf("failed to update field. %v", v)
		}
	})

	t.Run("method", func(t *testing.T) {
		rv := reflect.ValueOf(&v)
		if rv.NumMethod() != 2 {
			t.Errorf("NumMethod must 2. %v", rv.NumMethod())
		}
		// call FullName()
		m := rv.Method(0)
		ret := m.Call([]reflect.Value{})
		if ret[0].String() != "JOHN Doe" {
			t.Errorf("failed to call method. %v, %v", m, ret)
		}
	})
}
