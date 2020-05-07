package ast_examples

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	. "github.com/nirasan/go-reflect-examples/test_tools"
)

func TestParse(t *testing.T) {
	t.Run("parse ast from string", func(t *testing.T) {
		code := `
package main
import "fmt"
func main() {
    fmt.Println("Hello, World!")
}`
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, "", code, 0)
		Must(t, err, nil)

		// print ast
		err = ast.Print(fset, f)
		Must(t, err, nil)

		// get top level declaration
		Must(t, len(f.Decls), 2)
		// one of declaration is `import`
		Must(t, f.Decls[0].(*ast.GenDecl).Tok, token.IMPORT)
		// one of declaration is `func`
		Must(t, f.Decls[1].(*ast.FuncDecl).Name.Name, "main")
	})

	t.Run("parse ast from file", func(t *testing.T) {
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, "./testdata/parse/file1.go", nil, 0)
		Must(t, err, nil)

		// print ast
		err = ast.Print(fset, f)
		Must(t, err, nil)
	})

	t.Run("parse ast from dir", func(t *testing.T) {
		fset := token.NewFileSet()
		pkgs, err := parser.ParseDir(fset, "./testdata/parse/", nil, 0)
		Must(t, err, nil)

		// print ast
		err = ast.Print(fset, pkgs)
		Must(t, err, nil)

		// check files
		Must(t, len(pkgs), 1)
		Must(t, len(pkgs["main"].Files), 2)

		// print file1.go
		err = ast.Print(fset, pkgs["main"].Files["testdata/parse/file1.go"])
		Must(t, err, nil)
	})

}
