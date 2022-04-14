// +build example
//
// Do not build by default.

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/robtarr/gobot"
	"github.com/robtarr/gobot/drivers/i2c"
	"github.com/robtarr/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor(os.Args[1])
	bmp280 := i2c.NewBMP280Driver(firmataAdaptor)

	work := func() {
		gobot.Every(1*time.Second, func() {
			t, _ := bmp280.Temperature()
			fmt.Println("Temperature", t)

			p, _ := bmp280.Pressure()
			fmt.Println("Pressure", p)

			a, _ := bmp280.Altitude()
			fmt.Println("Altitude", a)
		})
	}

	robot := gobot.NewRobot("bmp280bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{bmp280},
		work,
	)

	robot.Start()
}
