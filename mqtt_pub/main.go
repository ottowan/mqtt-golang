package main

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {

	opts := mqtt.NewClientOptions().AddBroker("tcp://broker.hivemq.com:1883").
		SetClientID("ottowan03")

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	topic := "home/airconditioner/livingroom/temperature/unit"
	// c.Publish(topic, 1, false, "Example Payload")
	if token := c.Publish(topic, 1, false, "Example Payload"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}

}
