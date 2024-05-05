package generate

import (
	"fmt"
	"go/types"
	"io"
	"os"
	"os/exec"
	"text/template"
)

func GetAllStructsToConvert([][2]types.Var) []*types.Struct {
	return []*types.Struct{}
}

func Create(fileName string) {
	f, err := os.Open("./templ/out.tmpl")
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	templ := template.Must(template.New("out").Parse(string(b)))

	templ.Execute(createFile(fileName), TemplVar{
		PkgName:      "convert",
		SrcTypeName:  "blah.Foo",
		DestTypeName: "boo.Foo",
		Imports: []string{
			"github.com/janshercs/structconv/pkg/match/testfiles/blah",
			"github.com/janshercs/structconv/pkg/match/testfiles/boo",
		},
		BasicFields:  []string{"Foo", "Bar", "This", "That"},
		StructFields: []string{"Amount"},
	})
	f.Close()
}

func CreateAndFormat(fileName string) {
	Create(fileName)
	format(fileName)
}

func createFile(fileName string) *os.File {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	return file
}

type TemplVar struct {
	PkgName      string
	SrcTypeName  string
	DestTypeName string
	Imports      []string
	BasicFields  []string
	StructFields []string
}

func format(fileName string) {
	cmd := exec.Command("goimports", "-w", fileName)
	if errOut, err := cmd.CombinedOutput(); err != nil {
		panic(fmt.Errorf("failed to run %v with error: %s, %s", cmd, err, errOut))
	}
}

// TODO: How to recursively convert structs?
// TODO: How to collect all the struct defs?
// TODO: How to do struct pointers easily?
// TODO: Add struct tags

// OUTOFSCOPE: Anonymous structs
// OUTOFSCOPE: Anonymous structs
