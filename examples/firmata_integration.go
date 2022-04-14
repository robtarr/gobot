// +build example
//
// Do not build by default.

/*
 How to run
 Pass serial port to use as the first param:

	go run examples/firmata_integration.go /dev/ttyACM0
*/

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/robtarr/gobot"
	"github.com/robtarr/gobot/drivers/aio"
	"github.com/robtarr/gobot/drivers/gpio"
	"github.com/robtarr/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor(os.Args[1])
	led1 := gpio.NewLedDriver(firmataAdaptor, "3")
	led2 := gpio.NewLedDriver(firmataAdaptor, "4")
	button := gpio.NewButtonDriver(firmataAdaptor, "2")
	sensor := aio.NewAnalogSensorDriver(firmataAdaptor, "0")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led1.Toggle()
		})
		gobot.Every(2*time.Second, func() {
			led2.Toggle()
		})
		button.On(gpio.ButtonPush, func(data interface{}) {
			led2.On()
		})
		button.On(gpio.ButtonRelease, func(data interface{}) {
			led2.Off()
		})
		sensor.On(aio.Data, func(data interface{}) {
			fmt.Println("sensor", data)
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led1, led2, button, sensor},
		work,
	)

	robot.Start()
}
