// +build example
//
// Do not build by default.

package main

import (
	"fmt"
	"time"

	"github.com/robtarr/gobot"
	"github.com/robtarr/gobot/drivers/i2c"
	"github.com/robtarr/gobot/platforms/intel-iot/edison"
)

func main() {
	e := edison.NewAdaptor()
	blinkm := i2c.NewBlinkMDriver(e)

	work := func() {
		gobot.Every(3*time.Second, func() {
			r := byte(gobot.Rand(255))
			g := byte(gobot.Rand(255))
			b := byte(gobot.Rand(255))
			blinkm.Rgb(r, g, b)
			color, _ := blinkm.Color()
			fmt.Println("color", color)
		})
	}

	robot := gobot.NewRobot("blinkmBot",
		[]gobot.Connection{e},
		[]gobot.Device{blinkm},
		work,
	)

	robot.Start()
}
