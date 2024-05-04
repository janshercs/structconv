package match

import (
	"go/types"
)

// If both types are structs, return true.
// If both types are pointers, return true if both pointer underlying element are the same.
func CompareTypes(a, b types.Type) bool {
	aType := GetUnderlyingType(a)
	bType := GetUnderlyingType(b)

	switch at := aType.(type) {
	case *types.Struct:
		if _, ok := bType.(*types.Struct); ok {
			return true
		}

	case *types.Pointer:
		bPointer, ok := bType.(*types.Pointer)
		if !ok {
			return false
		}
		return CompareTypes(at.Elem(), bPointer.Elem())

	default:
		return aType == bType
	}
	return false
}

// MatchFields matches by field name and type
//
// if the type is a struct type, it will be deemed as matched.
func MatchFields(a, b *types.Struct) (matchedFields [][2]*types.Var) {
	if a == nil || b == nil {
		return
	}
	var fields = make(map[string]*types.Var)
	for i := range a.NumFields() {
		field := a.Field(i)
		fields[field.Name()] = field
	}

	for i := range b.NumFields() {
		field := b.Field(i)

		t, ok := fields[field.Name()]
		if !ok {
			continue
		}

		if CompareTypes(t.Type(), field.Type()) {
			matchedFields = append(matchedFields, [2]*types.Var{t, field})
		}
	}
	return
}

func GetUnderlyingType(t types.Type) types.Type {
	if t == t.Underlying() {
		return t
	}
	return GetUnderlyingType(t.Underlying())
}
