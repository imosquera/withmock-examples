package baz

type Qux struct {
}
type Baz struct {
	q Qux
}

func (b *Baz) GetBaz() string {
	return "blank string"
}

type Bazzer interface {
	GetBaz(q Qux) string
}
