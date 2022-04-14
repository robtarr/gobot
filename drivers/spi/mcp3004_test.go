package spi

import (
	"testing"

	"github.com/robtarr/gobot"
	"github.com/robtarr/gobotrr/gobot/drivers/aio"
	"github.com/robtarr/gobotrr/gobot/gobottest"
)

var _ gobot.Driver = (*MCP3004Driver)(nil)

// must implement the AnalogReader interface
var _ aio.AnalogReader = (*MCP3004Driver)(nil)

func initTestMCP3004Driver() *MCP3004Driver {
	d := NewMCP3004Driver(&TestConnector{})
	return d
}

func TestMCP3004DriverStart(t *testing.T) {
	d := initTestMCP3004Driver()
	gobottest.Assert(t, d.Start(), nil)
}

func TestMCP3004DriverHalt(t *testing.T) {
	d := initTestMCP3004Driver()
	d.Start()
	gobottest.Assert(t, d.Halt(), nil)
}

func TestMCP3004DriverRead(t *testing.T) {
	d := initTestMCP3004Driver()
	d.Start()

	// TODO: actual read test
}
