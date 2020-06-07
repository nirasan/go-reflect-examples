package main

//go:generate go run main.go -i s.go -o s_gen.go -s Struct1

type Struct1 struct {
	Int1    int
	Float1  float64
	Bool1   bool
	String1 string
}
