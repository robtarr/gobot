// +build example
//
// Do not build by default.

package main

import (
	"fmt"
	"time"

	"github.com/robtarr/gobot"
	"github.com/robtarr/gobot/drivers/i2c"
	"github.com/robtarr/gobot/platforms/chip"
)

func main() {
	board := chip.NewAdaptor()
	luxSensor := i2c.NewTSL2561Driver(board, i2c.WithTSL2561Gain16X)

	work := func() {
		gobot.Every(1*time.Second, func() {
			broadband, ir, err := luxSensor.GetLuminocity()

			if err != nil {
				fmt.Println("Err:", err)
			} else {
				light := luxSensor.CalculateLux(broadband, ir)
				fmt.Printf("BB: %v, IR: %v, Lux: %v\n", broadband, ir, light)
			}
		})
	}

	robot := gobot.NewRobot("tsl2561Bot",
		[]gobot.Connection{board},
		[]gobot.Device{luxSensor},
		work,
	)

	robot.Start()
}
