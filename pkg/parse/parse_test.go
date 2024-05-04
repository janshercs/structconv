package parse_test

import (
	"testing"

	"github.com/janshercs/structconv/pkg/parse"
	"github.com/stretchr/testify/assert"
)

func TestGetStructFromPackage(t *testing.T) {
	obj, s := parse.GetStructFromPackage("./testfiles/bar", "Foo")
	assert.NotNil(t, obj)
	assert.NotNil(t, s)

	assert.Equal(t, "Foo", obj.Name())
	assert.Equal(t, "github.com/janshercs/structconv/pkg/parse/testfiles/bar.Foo", obj.Type().String())
	assert.Equal(t, 6, s.NumFields())
}
