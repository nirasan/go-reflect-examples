package main

import "fmt"

var v1 = 1
var (
	v2 = "2"
	v3 = false
)

const c1 = 1
const (
	c2 = "2"
	c3 = true
)

type t1 int

const (
	ct1 t1 = iota
	ct2
	ct3
)

type s1 struct {
	f1 int
	f2 string
}

type i1 interface{}

func main() {
	fmt.Println("hello world")
}
