package main

import (
	"fmt"
	"os"
	"time"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	logger "github.com/sirupsen/logrus"
)



func messageHandler(c mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}



func main() {


	clientID := os.Args[1]
	topic := os.Args[2]
	message := os.Args[3]

	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://broker.hivemq.com:1883")
	opts.SetClientID(clientID)
	opts.SetKeepAlive(time.Second * time.Duration(60))
	// If lost connection, reconnect again
	opts.SetConnectionLostHandler(func(client mqtt.Client, e error) {
		logger.Warn(fmt.Sprintf("Connection lost : %v", e))
	})


	// connect to broker
	client := mqtt.NewClient(opts)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		logger.Fatalf("Fail to connect broker, %v",token.Error())
	}

	// publish to topic
	token = client.Publish(topic, byte(0), false, message)
	if token.Wait() && token.Error() != nil {
		logger.Errorf("Fail to publish, %v",token.Error())
	}




}
