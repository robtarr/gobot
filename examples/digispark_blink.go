// +build example
//
// Do not build by default.

package main

import (
	"fmt"
	"time"

	"github.com/robtarr/gobot"
	"github.com/robtarr/gobot/drivers/gpio"
	"github.com/robtarr/gobot/platforms/digispark"
)

func main() {
	digisparkAdaptor := digispark.NewAdaptor()
	led := gpio.NewLedDriver(digisparkAdaptor, "0")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{digisparkAdaptor},
		[]gobot.Device{led},
		work,
	)

	err := robot.Start()
	if err != nil {
		fmt.Println(err)
	}
}
