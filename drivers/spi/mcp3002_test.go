package spi

import (
	"testing"

	"github.com/robtarr/gobot"
	"github.com/robtarr/gobotrr/gobot/drivers/aio"
	"github.com/robtarr/gobotrr/gobot/gobottest"
)

var _ gobot.Driver = (*MCP3002Driver)(nil)

// must implement the AnalogReader interface
var _ aio.AnalogReader = (*MCP3002Driver)(nil)

func initTestMCP3002Driver() *MCP3002Driver {
	d := NewMCP3002Driver(&TestConnector{})
	return d
}

func TestMCP3002DriverStart(t *testing.T) {
	d := initTestMCP3002Driver()
	gobottest.Assert(t, d.Start(), nil)
}

func TestMCP3002DriverHalt(t *testing.T) {
	d := initTestMCP3002Driver()
	d.Start()
	gobottest.Assert(t, d.Halt(), nil)
}

func TestMCP3002DriverRead(t *testing.T) {
	d := initTestMCP3002Driver()
	d.Start()

	// TODO: actual read test
}
