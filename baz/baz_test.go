package baz

import (
	"code.google.com/p/gomock/gomock"
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

func (s *HookSuite) BazTest(c *C) {
	mockCtrl := gomock.NewController(c)
	defer mockCtrl.Finish()

	path.MOCK().SetController(mockCtrl)
	path.EXPECT().Join("basepath", "newpath/newfile.gz")

	baz := Baz{}
	baz.SplitPaths("basepath")

}
