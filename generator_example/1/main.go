package main

import (
	"bytes"
	"go/format"
	"io"
	"os"
)

func main() {
	if err := Generate("world", os.Stdout); err != nil {
		panic(err)
	}
}

func Generate(name string, w io.Writer) error {
	var b bytes.Buffer

	/* コードの作成 */

	// このコメントは生成されたコードであることを明示するためのもの。
	// 人間が読むだけではなく各種ツールも利用するもので、例えば Linter なら生成されたコードのチェックを行わないなど形で利用される。
	// フォーマットは `^// Code generated .* DO NOT EDIT\.$` という正規表現で定義されている。
	b.WriteString("// Code generated \"Generate\"; DO NOT EDIT.\n") // 生成されたコードであることを明示する

	b.WriteString("package main\n")
	b.WriteString("import \"fmt\"\n")
	b.WriteString("func main() {\n")
	b.WriteString("fmt.Println(\"hello " + name + "\")\n")
	b.WriteString("}\n")

	/* コードの文法チェックと整形 */
	src, err := format.Source(b.Bytes())
	if err != nil {
		return err
	}

	/* コードの出力 */
	if _, err := w.Write(src); err != nil {
		return err
	}

	return nil
}
