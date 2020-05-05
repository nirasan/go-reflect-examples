package reflect_example

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	v := map[string]int{
		"key1": 1,
		"key2": 2,
	}

	rt := reflect.TypeOf(v)

	// get Kind
	if rt.Kind() != reflect.Map {
		t.Errorf("invalid kind. %v", rt.Kind())
	}

	// get key and value Kind
	if rt.Key().Kind() != reflect.String || rt.Elem().Kind() != reflect.Int {
		t.Errorf("invalid kind. %v, %v", rt.Key(), rt.Elem())
	}

	rv := reflect.ValueOf(v)

	// loop using MapKeys
	keys := rv.MapKeys()
	if len(keys) != 2 {
		t.Errorf("length of keys must 2. %v", keys)
	}
	for _, key := range keys {
		val := rv.MapIndex(key)
		if key.String() == "key1" && val.Int() != 1 {
			t.Errorf("invalid element. %v, %v", key, val)
		}
		if key.String() == "key2" && val.Int() != 2 {
			t.Errorf("invalid element. %v, %v", key, val)
		}
	}

	// loop using MapRange
	iter := rv.MapRange()
	for iter.Next() {
		key := iter.Key()
		val := iter.Value()
		if key.String() == "key1" && val.Int() != 1 {
			t.Errorf("invalid element. %v, %v", key, val)
		}
		if key.String() == "key2" && val.Int() != 2 {
			t.Errorf("invalid element. %v, %v", key, val)
		}
	}

	// update
	rv = reflect.ValueOf(&v).Elem()
	rv.SetMapIndex(reflect.ValueOf("key2"), reflect.ValueOf(200))
	if v["key2"] != 200 {
		t.Errorf("failed to update. %v", v)
	}

	// create map from Type
	rv = reflect.MakeMap(reflect.TypeOf(map[string]int{}))
	if rv.Kind() != reflect.Map || rv.Len() != 0 {
		t.Errorf("invalid value. %v", rv)
	}
}
