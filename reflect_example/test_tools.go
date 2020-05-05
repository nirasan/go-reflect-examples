package reflect_example

import (
	"reflect"
	"testing"
)

func Must(t *testing.T, l, r interface{}) {
	if !reflect.DeepEqual(l, r) {
		t.Errorf("%#v must equals to %#v", l, r)
	}
}
