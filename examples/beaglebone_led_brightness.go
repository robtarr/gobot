// +build example
//
// Do not build by default.

package main

import (
	"time"

	"github.com/robtarr/gobot"
	"github.com/robtarr/gobot/drivers/gpio"
	"github.com/robtarr/gobot/platforms/beaglebone"
)

func main() {
	beagleboneAdaptor := beaglebone.NewAdaptor()
	led := gpio.NewLedDriver(beagleboneAdaptor, "P9_14")

	work := func() {
		brightness := uint8(0)
		fadeAmount := uint8(5)

		gobot.Every(100*time.Millisecond, func() {
			led.Brightness(brightness)
			brightness = brightness + fadeAmount
			if brightness == 0 || brightness == 255 {
				fadeAmount = -fadeAmount
			}
		})
	}

	robot := gobot.NewRobot("pwmBot",
		[]gobot.Connection{beagleboneAdaptor},
		[]gobot.Device{led},
		work,
	)

	robot.Start()
}
