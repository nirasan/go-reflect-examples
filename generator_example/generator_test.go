package generator_example

import (
	"testing"

	. "github.com/nirasan/go-reflect-examples/test_tools"
)

func TestGenerator(t *testing.T) {
	g := &Generator{}
	b, err := g.Generate("./testdata/generator/")
	Must(t, err, nil)
	Must(t, string(b), `package generator

import "fmt"

func (s *S3) String() string { return fmt.Sprintf("&%+v", s) }
func (s *S1) String() string { return fmt.Sprintf("&%+v", s) }
func (s *S2) String() string { return fmt.Sprintf("&%+v", s) }
`)
}
