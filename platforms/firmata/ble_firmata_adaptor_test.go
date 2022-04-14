// +build !windows

package firmata

import (
	"strings"
	"testing"

	"github.com/robtarr/gobot"
	"github.com/robtarr/gobotrr/gobot/gobottest"
)

var _ gobot.Adaptor = (*BLEAdaptor)(nil)

func initTestBLEAdaptor() *BLEAdaptor {
	a := NewBLEAdaptor("DEVICE", "123", "456")
	return a
}

func TestFirmataBLEAdaptor(t *testing.T) {
	a := initTestBLEAdaptor()
	gobottest.Assert(t, strings.HasPrefix(a.Name(), "BLEFirmata"), true)
}
