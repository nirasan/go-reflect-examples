package main

import (
	"flag"
	"go/ast"
	"go/parser"
	"go/token"
)

//go:generate go run main.go -- $GOFILE

func main() {
	flag.Parse()
	filename := flag.Arg(0)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		panic(err)
	}

	if err := ast.Print(fset, f); err != nil {
		panic(err)
	}
}
