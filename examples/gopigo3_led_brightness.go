// +build example
//
// Do not build by default.

package main

import (
	"time"

	"github.com/robtarr/gobot"
	"github.com/robtarr/gobot/drivers/gpio"
	"github.com/robtarr/gobot/platforms/dexter/gopigo3"
	"github.com/robtarr/gobot/platforms/raspi"
)

func main() {
	raspiAdaptor := raspi.NewAdaptor()
	gpg3 := gopigo3.NewDriver(raspiAdaptor)
	led := gpio.NewGroveLedDriver(gpg3, "AD_1_1")

	work := func() {
		brightness := uint8(0)
		fadeAmount := uint8(15)

		gobot.Every(100*time.Millisecond, func() {
			led.Brightness(brightness)
			brightness = brightness + fadeAmount
			if brightness == 0 || brightness == 255 {
				fadeAmount = -fadeAmount
			}
		})
	}

	robot := gobot.NewRobot("gopigo3pwm",
		[]gobot.Connection{raspiAdaptor},
		[]gobot.Device{gpg3, led},
		work,
	)

	robot.Start()
}
