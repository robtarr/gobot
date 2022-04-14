package firmata

import (
	"strings"
	"testing"

	"github.com/robtarr/gobot"
	"github.com/robtarr/gobotrr/gobot/gobottest"
)

var _ gobot.Adaptor = (*TCPAdaptor)(nil)

func initTestTCPAdaptor() *TCPAdaptor {
	a := NewTCPAdaptor("localhost:4567")
	return a
}

func TestFirmataTCPAdaptor(t *testing.T) {
	a := initTestTCPAdaptor()
	gobottest.Assert(t, strings.HasPrefix(a.Name(), "TCPFirmata"), true)
}
