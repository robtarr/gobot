package spi

import (
	"testing"

	"github.com/robtarr/gobot"
	"github.com/robtarr/gobotrr/gobot/drivers/aio"
	"github.com/robtarr/gobotrr/gobot/gobottest"
)

var _ gobot.Driver = (*MCP3208Driver)(nil)

// must implement the AnalogReader interface
var _ aio.AnalogReader = (*MCP3208Driver)(nil)

func initTestMCP3208Driver() *MCP3208Driver {
	d := NewMCP3208Driver(&TestConnector{})
	return d
}

func TestMCP3208DriverStart(t *testing.T) {
	d := initTestMCP3208Driver()
	gobottest.Assert(t, d.Start(), nil)
}

func TestMCP3208DriverHalt(t *testing.T) {
	d := initTestMCP3208Driver()
	d.Start()
	gobottest.Assert(t, d.Halt(), nil)
}

func TestMCP3208DriverRead(t *testing.T) {
	d := initTestMCP3208Driver()
	d.Start()

	// TODO: actual read test
}
