package bartest

type Foo struct {
	Foo    string
	Bar    int
	Baz    baz
	Amount *Amount
	This   *string
	That   *int
}

type baz string

type this func()

func thatsit() {}

type Amount struct {
	Amount   int
	Currency string
}
