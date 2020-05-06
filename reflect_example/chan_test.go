package reflect_example

import (
	"reflect"
	"testing"

	. "github.com/nirasan/go-reflect-examples/test_tools"
)

func TestChan(t *testing.T) {
	v := make(chan int, 3)

	rt := reflect.TypeOf(v)

	Must(t, rt.Kind(), reflect.Chan)

	Must(t, rt.Elem().Kind(), reflect.Int)

	rv := reflect.ValueOf(v)

	rv.Send(reflect.ValueOf(100))

	ret, ok := rv.Recv()
	Must(t, ok, true)
	Must(t, ret.Int(), int64(100))

	v <- 200
	index, recv, ok := reflect.Select([]reflect.SelectCase{reflect.SelectCase{
		Dir:  reflect.SelectRecv,
		Chan: rv,
	}})
	Must(t, index, 0)
	Must(t, recv.Int(), int64(200))
	Must(t, ok, true)
}
