package main

import (
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func messageHandler(c mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func connLostHandler(c mqtt.Client, err error) {
	fmt.Printf("Connection lost, reason: %v\n", err)

	//Perform additional action...
}

func main() {

	//mqtt.DEBUG = log.New(os.Stderr, "DEBUG ", log.Ltime)
	//create a ClientOptions
	opts := mqtt.NewClientOptions().
		AddBroker("tcp://broker.hivemq.com:1883").
		SetClientID("ottowan01").
		SetDefaultPublishHandler(messageHandler).
		SetConnectionLostHandler(connLostHandler)

	topic := "home/airconditioner/livingroom/temperature/unit"

	//topic := "home/airconditioner/livingroom/temperature"
	//topic := "home/#"
	opts.OnConnect = func(c mqtt.Client) {
		fmt.Printf("Client connected, subscribing to: " + topic + "\n")

		//Subscribe here, otherwise after connection lost,
		//you may not receive any message
		if token := c.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}

	}

	//create and start a client using the above ClientOptions
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for {
		//Lazy...
		time.Sleep(500 * time.Millisecond)
	}
}
