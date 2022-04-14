package bebop

import (
	"strings"
	"testing"

	"github.com/robtarr/gobot"
	"github.com/robtarr/gobotrr/gobot/gobottest"
)

var _ gobot.Driver = (*Driver)(nil)

func TestBebopDriverName(t *testing.T) {
	a := initTestBebopAdaptor()
	d := NewDriver(a)
	gobottest.Assert(t, strings.HasPrefix(d.Name(), "Bebop"), true)
	d.SetName("NewName")
	gobottest.Assert(t, d.Name(), "NewName")
}
