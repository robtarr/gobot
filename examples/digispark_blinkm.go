// +build example
//
// Do not build by default.

package main

import (
	"fmt"
	"time"

	"github.com/robtarr/gobot"
	"github.com/robtarr/gobot/drivers/i2c"
	"github.com/robtarr/gobot/platforms/digispark"
)

func main() {
	board := digispark.NewAdaptor()
	blinkm := i2c.NewBlinkMDriver(board)

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

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{board},
		[]gobot.Device{blinkm},
		work,
	)

	err := robot.Start()
	if err != nil {
		fmt.Println(err)
	}
}
