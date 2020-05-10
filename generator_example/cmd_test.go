package generator_example

import (
	"os/exec"
	"testing"

	. "github.com/nirasan/go-reflect-examples/test_tools"
)

func TestCmd(t *testing.T) {
	out, err := exec.Command("go", "generate", "./testdata/cmd/").Output()
	Must(t, err, nil)
	Must(t, string(out), `package generator

import "fmt"

func (s *S3) String() string { return fmt.Sprintf("&%+v", s) }
func (s *S1) String() string { return fmt.Sprintf("&%+v", s) }
func (s *S2) String() string { return fmt.Sprintf("&%+v", s) }
`)
}
