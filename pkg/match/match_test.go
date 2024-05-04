package match_test

import (
	"log"
	"testing"

	"github.com/janshercs/structconv/pkg/match"
	"github.com/janshercs/structconv/pkg/parse"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMatchFields(t *testing.T) {
	_, a := parse.GetStructFromPackage("./testfiles/blah", "Foo")
	require.NotNil(t, a)

	_, b := parse.GetStructFromPackage("./testfiles/boo", "Foo")
	require.NotNil(t, b)

	fields := match.MatchFields(a, b)

	assert.Len(t, fields, 5)

	for _, field := range fields {
		log.Print("field name: ", field[0].Name())
		log.Print("field type: ", field[0].Type())
		log.Print("package type: ", field[0].Pkg().Path())
	}
}
