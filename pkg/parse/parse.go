package parse

import (
	"go/types"
	"log"

	"golang.org/x/tools/go/packages"
)

func GetPackage(pkgPath string) *packages.Package {
	pkgs, err := packages.Load(
		&packages.Config{Mode: packages.NeedTypes | packages.NeedImports | packages.NeedSyntax}, // uses bitwise operations, smart! learn here: https://blog.learngoprogramming.com/golang-const-type-enums-iota-bc4befd096d3
		pkgPath,
	)
	if err != nil {
		panic(err)
	}

	if len(pkgs) > 1 {
		log.Printf("found more than 1 packages with the pattern %q, returning only the first one.", pkgs)
	}

	return pkgs[0]
}

func GetStructDef(pkg *packages.Package, structName string) (types.Object, *types.Struct) {
	obj := pkg.Types.Scope().Lookup(structName)
	if obj == nil {
		return nil, nil
	}
	s, ok := obj.Type().Underlying().(*types.Struct)
	if !ok {
		return nil, nil
	}
	return obj, s
}

func GetStructFromPackage(pkgName, structName string) (types.Object, *types.Struct) {
	return GetStructDef(GetPackage(pkgName), structName)
}
