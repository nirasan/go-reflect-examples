package main

import (
	"fmt"
	"os"

	"github.com/nirasan/go-reflect-examples/generator_example"
)

func main() {
	g := &generator_example.Generator{}
	b, err := g.Generate(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Print(string(b))
}
