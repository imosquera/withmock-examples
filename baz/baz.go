package baz

import (
	"path"
)

type Qux struct{}

func (q *Qux) DoSomething() string {
	return "did something"
}

type Quxxer interface {
	DoSomething()
}

type Baz struct {
	QuxA Quxxer
}

func (b *Baz) GetBaz() string {
	return "blank string"
}

func (b *Baz) SplitPaths(p string) []string {
	b.QuxA.DoSomething()
	newpath := path.Join(p, "newpath/newfile.gz")
	return []string{newpath}
}

type Bazzer interface {
	SplitPaths(string) []string
	GetBaz(q Qux) string
}
