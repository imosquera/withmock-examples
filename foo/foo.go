package foo

import (
	"github.com/imosquera/withmock-examples/baz"
)

type Bar struct {
	bazzer baz.Bazzer
}

func (self *Bar) GetBar() string {
	q := baz.Qux{}
	self.bazzer.GetBaz(q)
	newStr := "1" + "0"
	return newStr
}

type Barer interface {
	GetBar()
}
