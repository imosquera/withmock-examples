package baz_test

import (
	"code.google.com/p/gomock/gomock"
	"github.com/imosquera/withmock-examples/baz"
	"github.com/imosquera/withmock-examples/baz/_mocks_"
	. "launchpad.net/gocheck"
	"path" //mock
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type HookSuite struct{}

var _ = Suite(&HookSuite{})

func (s *HookSuite) SetupTest(c *C) {
}

func (s *HookSuite) TearDownTest(c *C) {
}

func (s *HookSuite) TestBaz(c *C) {
	mockCtrl := gomock.NewController(c)
	defer mockCtrl.Finish()

	baz_mocks.SetController(mockCtrl)

	path.MOCK().SetController(mockCtrl)
	path.EXPECT().Join("basepath", "newpath/newfile.gz")

	quxxer := baz_mocks.NewQuxxer()
	quxxer.EXPECT().DoSomething()

	baz := baz.Baz{QuxA: quxxer}
	baz.SplitPaths("basepath")

}
