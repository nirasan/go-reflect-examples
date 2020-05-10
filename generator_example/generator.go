package generator_example

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"strings"
)

type Generator struct {
}

func (g *Generator) Generate(path string) ([]byte, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, path, nil, 0)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	for name, pkg := range pkgs {
		b.WriteString(fmt.Sprintf("package %s\n", name))
		b.WriteString("import \"fmt\"\n")
		for _, file := range pkg.Files {
			for _, d := range file.Decls {
				dd, ok := d.(*ast.GenDecl)
				if !ok || dd.Tok != token.TYPE {
					continue
				}
				s := dd.Specs[0].(*ast.TypeSpec).Name.Name
				r := strings.ToLower(s[:1])
				b.WriteString(fmt.Sprintf("func (%s *%s) String() string { return fmt.Sprintf(\"&%%+v\", %s) }\n", r, s, r))
			}
		}
	}

	return format.Source(b.Bytes())
}
