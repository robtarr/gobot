/*
Package ardrone provides the Gobot adaptor and driver for the Parrot Ardrone.

Installing:

	go get -d -u github.com/robtarr/gobot/... && go install github.com/robtarr/gobot/platforms/parrot/ardrone

Example:

	package main

	import (
		"time"

		"github.com/robtarr/gobotrr/gobot"
		"github.com/robtarr/gobotrr/gobot/platforms/parrot/ardrone"
	)

	func main() {
		ardroneAdaptor := ardrone.NewAdaptor()
		drone := ardrone.NewDriver(ardroneAdaptor)

		work := func() {
			drone.TakeOff()
			drone.On(drone.Event("flying"), func(data interface{}) {
				gobot.After(3*time.Second, func() {
					drone.Land()
				})
			})
		}

		robot := gobot.NewRobot("drone",
			[]gobot.Connection{ardroneAdaptor},
			[]gobot.Device{drone},
			work,
		)

		robot.Start()
	}

For more information refer to the ardrone README:
https://github.com/hybridgroup/gobot/tree/master/platforms/parrot/ardrone/README.md
*/
package ardrone // import "github.com/robtarr/gobotrr/gobot/platforms/parrot/ardrone"
