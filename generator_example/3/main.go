package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
)

var (
	inputFilename  = flag.String("i", "", "input file name")
	outputFilename = flag.String("o", "", "output file name")
	targetStruct   = flag.String("s", "", "target struct name")
)

func main() {
	flag.Parse()

	output, err := os.OpenFile(*outputFilename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, *inputFilename, nil, 0)
	if err != nil {
		log.Fatal(err)
	}

	s := *targetStruct
	var buf bytes.Buffer

	ast.Inspect(f, func(node ast.Node) bool {
		// 型宣言は *ast.TypeSpec
		v, ok := node.(*ast.TypeSpec)
		// 型名が対象の構造体名でなければ終了
		if !ok || v.Name.Name != s {
			return true
		}
		// フィールド名のリストを取得
		fields := []string{}
		for _, vv := range v.Type.(*ast.StructType).Fields.List {
			for _, n := range vv.Names {
				fields = append(fields, n.Name)
			}
		}
		// パッケージ名
		p := f.Name.Name
		// レシーバー名
		r := strings.ToLower(s)[0]

		// コードの描画開始
		// 対象の構造体の String メソッドを宣言する
		buf.WriteString("// Code generated ; DO NOT EDIT.\n")
		buf.WriteString("package " + p + "\n")
		buf.WriteString("import \"fmt\"\n")
		buf.WriteString(fmt.Sprintf("func (%c %s) String() string {\n", r, *targetStruct))
		// String のフォーマットとフォーマットの引数を作成
		formats := []string{}
		values := []string{}
		for _, field := range fields {
			formats = append(formats, fmt.Sprintf("\\t%s:%%v", field))
			values = append(values, fmt.Sprintf("%c.%s", r, field))
		}
		buf.WriteString(fmt.Sprintf("return fmt.Sprintf(\"%s{\\n%s\\n}\", %s)\n", s, strings.Join(formats, "\\n"), strings.Join(values, ",")))
		buf.WriteString("}\n")
		return true
	})

	// コードの文法チェックと整形
	src, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	if _, err = output.Write(src); err != nil {
		log.Fatal(err)
	}
}
