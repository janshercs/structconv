package generate_test

import (
	"testing"

	"github.com/janshercs/structconv/pkg/generate"
)

func TestDemo(t *testing.T) {
	generate.CreateAndFormat("./output/demo.gen.go")
}
