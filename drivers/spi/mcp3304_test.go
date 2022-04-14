package spi

import (
	"testing"

	"github.com/robtarr/gobot"
	"github.com/robtarr/gobotrr/gobot/drivers/aio"
	"github.com/robtarr/gobotrr/gobot/gobottest"
)

var _ gobot.Driver = (*MCP3304Driver)(nil)

// must implement the AnalogReader interface
var _ aio.AnalogReader = (*MCP3304Driver)(nil)

func initTestMCP3304Driver() *MCP3304Driver {
	d := NewMCP3304Driver(&TestConnector{})
	return d
}

func TestMCP3304DriverStart(t *testing.T) {
	d := initTestMCP3304Driver()
	gobottest.Assert(t, d.Start(), nil)
}

func TestMCP3304DriverHalt(t *testing.T) {
	d := initTestMCP3304Driver()
	d.Start()
	gobottest.Assert(t, d.Halt(), nil)
}

func TestMCP3304DriverRead(t *testing.T) {
	d := initTestMCP3304Driver()
	d.Start()

	// TODO: actual read test
}
