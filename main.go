package main

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/mqtt"
	"github.com/hybridgroup/gobot/platforms/raspi"
)

func main() {
	gbot := gobot.NewGobot()
	mqttAdaptor := mqtt.NewMqttAdaptor("go-iot", "tcp://test.mosquitto.org:1883", "raspy")
	r := raspi.NewRaspiAdaptor("raspi")
	dirA := gpio.NewLedDriver(r, "dirA", "3")
	dirB := gpio.NewLedDriver(r, "dirB", "5")
	pwmA := gpio.NewLedDriver(r, "pwmA", "7")
	pwmB := gpio.NewLedDriver(r, "pwmB", "15")

	work := func() {
		dirA.Off()
		dirB.Off()
		pwmA.Off()
		pwmB.Off()

		mqttAdaptor.On("go-iot", func(data []byte) {
			if string(data) == "su" {
				dirA.Off()
				dirB.Off()
				pwmA.On()
				pwmB.On()
			}
			if string(data) == "giu" {
				dirA.On()
				dirB.On()
				pwmA.Off()
				pwmB.Off()
			}
			if string(data) == "left" {
				dirA.Off()
				dirB.On()
				pwmA.Off()
				pwmB.On()
			}
			if string(data) == "right" {
				dirA.On()
				dirB.Off()
				pwmA.On()
				pwmB.Off()
			}
		})
	}

	car := gobot.NewRobot("mqttBot",
		[]gobot.Connection{r, mqttAdaptor},
		[]gobot.Device{dirA, dirB, pwmA, pwmB},
		work,
	)

	gbot.AddRobot(car)

	gbot.Start()
}
