package convert

import (
	"github.com/janshercs/structconv/pkg/match/testfiles/blah"
	"github.com/janshercs/structconv/pkg/match/testfiles/boo"
)

func Convert(v blah.Foo) boo.Foo {
	return boo.Foo{
		Foo:  v.Foo,
		Bar:  v.Bar,
		This: v.This,
		That: v.That,

		Amount: ConvertAmount(v.Amount),
	}
}

func ConvertAmount(v *blah.Amount) *boo.Amount {
	return &boo.Amount{}
}
