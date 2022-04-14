// +build example
//
// Do not build by default.

package main

import (
	"time"

	"github.com/robtarr/gobot"
	"github.com/robtarr/gobot/drivers/gpio"
	"github.com/robtarr/gobot/drivers/i2c"
	"github.com/robtarr/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()
	gp := i2c.NewGrovePiDriver(r)
	led := gpio.NewLedDriver(gp, "D2")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		[]gobot.Device{gp, led},
		work,
	)

	robot.Start()
}
