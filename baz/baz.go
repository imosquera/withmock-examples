package baz

import (
	"path"
)

type Qux struct {
}

type Baz struct {
	q Qux
}

func (b *Baz) GetBaz() string {
	return "blank string"
}

func (b *Baz) SplitPaths(p string) []string {
	newpath := path.Join(p, "newpath/newfile.gz")
	return []string{newpath}
}

type Bazzer interface {
	SplitPaths(string) []string
	GetBaz(q Qux) string
}
