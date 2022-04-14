// +build example
//
// Do not build by default.

package main

import (
	"fmt"

	"github.com/robtarr/gobot"
	"github.com/robtarr/gobot/drivers/gpio"
	"github.com/robtarr/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()
	stepper := gpio.NewStepperDriver(r, [4]string{"7", "11", "13", "15"}, gpio.StepperModes.DualPhaseStepping, 2048)

	work := func() {
		//set spped
		stepper.SetSpeed(15)

		//Move forward one revolution
		if err := stepper.Move(2048); err != nil {
			fmt.Println(err)
		}

		//Move backward one revolution
		if err := stepper.Move(-2048); err != nil {
			fmt.Println(err)
		}
	}

	robot := gobot.NewRobot("stepperBot",
		[]gobot.Connection{r},
		[]gobot.Device{stepper},
		work,
	)

	robot.Start()
}
