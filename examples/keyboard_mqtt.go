// +build example
//
// Do not build by default.

package main

import (
	"github.com/robtarr/gobot"
	"github.com/robtarr/gobot/platforms/keyboard"
	"github.com/robtarr/gobot/platforms/mqtt"
)

func main() {
	keys := keyboard.NewDriver()
	mqttAdaptor := mqtt.NewAdaptor("tcp://iot.eclipse.org:1883", "conductor")

	work := func() {
		keys.On(keyboard.Key, func(data interface{}) {
			key := data.(keyboard.KeyEvent)

			switch key.Key {
			case keyboard.ArrowUp:
				mqttAdaptor.Publish("rover/frente", []byte{})
			case keyboard.ArrowRight:
				mqttAdaptor.Publish("rover/derecha", []byte{})
			case keyboard.ArrowDown:
				mqttAdaptor.Publish("rover/atras", []byte{})
			case keyboard.ArrowLeft:
				mqttAdaptor.Publish("rover/izquierda", []byte{})
			}
		})
	}

	robot := gobot.NewRobot("keyboardbot",
		[]gobot.Connection{mqttAdaptor},
		[]gobot.Device{keys},
		work,
	)

	robot.Start()
}
