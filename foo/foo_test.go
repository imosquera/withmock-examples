package foo

import (
	"code.google.com/p/gomock/gomock"
	baz "github.com/imosquera/withmock-examples/baz" //mock
	. "launchpad.net/gocheck"
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

func (s *HookSuite) TestBar(c *C) {
	mockCtrl := gomock.NewController(c)
	defer mockCtrl.Finish()

	baz.MOCK().SetController(mockCtrl)
	bazzer := baz.MOCK().NewBazzer()
	bazzer.EXPECT().GetBaz(gomock.Any())
	bar := Bar{}
	bar.bazzer = bazzer
	bar.GetBar()
}
