// +build example
//
// Do not build by default.

package main

import (
	"github.com/robtarr/gobot"
	"github.com/robtarr/gobot/api"
	"github.com/robtarr/gobot/drivers/gpio"
	"github.com/robtarr/gobot/platforms/digispark"
)

func main() {
	master := gobot.NewMaster()
	api.NewAPI(master).Start()

	digisparkAdaptor := digispark.NewAdaptor()
	led := gpio.NewLedDriver(digisparkAdaptor, "0")

	robot := gobot.NewRobot("digispark",
		[]gobot.Connection{digisparkAdaptor},
		[]gobot.Device{led},
	)

	master.AddRobot(robot)

	master.Start()
}
