package main

import "testing"

func TestStruct1_String(t *testing.T) {
	s := Struct1{
		Int1:    1,
		Float1:  2,
		Bool1:   true,
		String1: "aaa",
	}
	t.Log(s)
}
