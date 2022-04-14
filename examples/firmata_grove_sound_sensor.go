// +build example
//
// Do not build by default.

package main

import (
	"fmt"
	"os"

	"github.com/robtarr/gobot"
	"github.com/robtarr/gobot/drivers/aio"
	"github.com/robtarr/gobot/platforms/firmata"
)

func main() {
	board := firmata.NewAdaptor(os.Args[1])
	sensor := aio.NewGroveSoundSensorDriver(board, "3")

	work := func() {
		sensor.On(aio.Data, func(data interface{}) {
			fmt.Println("sensor", data)
		})
	}

	robot := gobot.NewRobot("sensorBot",
		[]gobot.Connection{board},
		[]gobot.Device{sensor},
		work,
	)

	robot.Start()
}
