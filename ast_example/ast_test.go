package ast_example

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	. "github.com/nirasan/go-reflect-examples/test_tools"
)

func TestAst(t *testing.T) {
	t.Run("get declarations", func(t *testing.T) {
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, "./testdata/ast/file1.go", nil, 0)
		Must(t, err, nil)

		err = ast.Print(fset, f)
		Must(t, err, nil)

		// import declaration
		importDecl := f.Decls[0].(*ast.GenDecl)
		Must(t, importDecl.Tok, token.IMPORT)
		Must(t, len(importDecl.Specs), 1)
		Must(t, importDecl.Specs[0].(*ast.ImportSpec).Path.Value, `"fmt"`)

		// var declaration
		varDecl1 := f.Decls[1].(*ast.GenDecl)
		Must(t, varDecl1.Tok, token.VAR)
		Must(t, len(varDecl1.Specs), 1)
		Must(t, varDecl1.Specs[0].(*ast.ValueSpec).Names[0].Name, "v1")
		Must(t, varDecl1.Specs[0].(*ast.ValueSpec).Values[0].(*ast.BasicLit).Value, "1")

		// multiple var declaration
		varDecl2 := f.Decls[2].(*ast.GenDecl)
		Must(t, varDecl2.Tok, token.VAR)
		Must(t, len(varDecl2.Specs), 2)
		Must(t, varDecl2.Specs[0].(*ast.ValueSpec).Names[0].Name, "v2")
		Must(t, varDecl2.Specs[0].(*ast.ValueSpec).Values[0].(*ast.BasicLit).Value, `"2"`)

		// const declaration
		constDecl1 := f.Decls[3].(*ast.GenDecl)
		Must(t, constDecl1.Tok, token.CONST)
		Must(t, len(constDecl1.Specs), 1)
		Must(t, constDecl1.Specs[0].(*ast.ValueSpec).Names[0].Name, "c1")
		Must(t, constDecl1.Specs[0].(*ast.ValueSpec).Values[0].(*ast.BasicLit).Value, "1")

		// multiple const declaration
		constDecl2 := f.Decls[4].(*ast.GenDecl)
		Must(t, constDecl2.Tok, token.CONST)
		Must(t, len(constDecl2.Specs), 2)
		Must(t, constDecl2.Specs[0].(*ast.ValueSpec).Names[0].Name, "c2")
		Must(t, constDecl2.Specs[0].(*ast.ValueSpec).Values[0].(*ast.BasicLit).Value, `"2"`)

		// struct declaration
		structDecl := f.Decls[7].(*ast.GenDecl)
		Must(t, structDecl.Tok, token.TYPE)
		Must(t, structDecl.Specs[0].(*ast.TypeSpec).Name.Name, "s1")
		Must(t, structDecl.Specs[0].(*ast.TypeSpec).Type.(*ast.StructType).Fields.List[0].Names[0].Name, "f1")
		Must(t, structDecl.Specs[0].(*ast.TypeSpec).Type.(*ast.StructType).Fields.List[1].Names[0].Name, "f2")

		// interface declaration
		interfaceDecl := f.Decls[8].(*ast.GenDecl)
		Must(t, interfaceDecl.Tok, token.TYPE)
		Must(t, interfaceDecl.Specs[0].(*ast.TypeSpec).Name.Name, "i1")

		// func declaration
		funcDecl := f.Decls[9].(*ast.FuncDecl)
		Must(t, funcDecl.Name.Name, "main")
	})

	t.Run("inspect nodes", func(t *testing.T) {
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, "./testdata/ast/file2.go", nil, 0)
		Must(t, err, nil)

		ast.Inspect(f, func(node ast.Node) bool {
			switch v := node.(type) {
			case *ast.GenDecl:
				if v.Tok == token.TYPE {
					s := v.Specs[0].(*ast.TypeSpec)
					t.Logf("struct: %s", s.Name.Name)
					for _, f := range s.Type.(*ast.StructType).Fields.List {
						t.Logf("field: %s", f.Names[0].Name)
					}
				}
			}
			return true
		})
	})

	t.Run("parse comment", func(t *testing.T) {
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, "./testdata/ast/file3.go", nil, parser.ParseComments)
		Must(t, err, nil)

		err = ast.Print(fset, f)
		Must(t, err, nil)

		ast.Inspect(f, func(node ast.Node) bool {
			switch v := node.(type) {
			case *ast.GenDecl:
				if v.Doc != nil {
					if v.Tok == token.TYPE {
						Must(t, v.Doc.List[0].Text, "// struct comment")
						Must(t, v.Specs[0].(*ast.TypeSpec).Name.Name, "s5")
					} else if v.Tok == token.CONST {
						Must(t, v.Doc.List[0].Text, "// const comment")
						Must(t, v.Specs[0].(*ast.ValueSpec).Names[0].Name, "c5")
					}
				}
			}
			return true
		})
	})

}
