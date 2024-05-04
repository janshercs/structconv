package footest

type Foo struct {
	Foo    string
	Bar    int
	Baz    baz
	Amount *Amount
	This   *string
	That   *int
}

type Bar struct {
	baz string
}

type baz interface{}

type Amount struct {
	Amount   int
	Currency string
}
