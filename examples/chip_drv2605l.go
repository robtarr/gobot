// +build example
//
// Do not build by default.

package main

import (
	"time"

	"github.com/robtarr/gobot"
	"github.com/robtarr/gobot/drivers/i2c"
	"github.com/robtarr/gobot/platforms/chip"
)

func main() {
	board := chip.NewAdaptor()
	haptic := i2c.NewDRV2605LDriver(board)

	work := func() {
		gobot.Every(3*time.Second, func() {
			pause := haptic.GetPauseWaveform(50)
			haptic.SetSequence([]byte{1, pause, 1, pause, 1})
			haptic.Go()
		})
	}

	robot := gobot.NewRobot("DRV2605LBot",
		[]gobot.Connection{board},
		[]gobot.Device{haptic},
		work,
	)

	robot.Start()
}
