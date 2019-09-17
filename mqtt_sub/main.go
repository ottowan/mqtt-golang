package main

import (
	"fmt"
	"os"
	"time"
	"strconv"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var count = 0
var lastTopic = ""
func messageHandler(c mqtt.Client, msg mqtt.Message) {

	count = count+1
	lastTopic = msg.Topic()
	//fmt.Println(count)
	//fmt.Println(len(msg))
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())

}

func connLostHandler(c mqtt.Client, err error) {
	fmt.Printf("Connection lost, reason: %v\n", err)

	//Perform additional action...
}


var topic = []string{
	"home/#",
	"home/garage/#",
	"home/garage/lighting/#",
	"home/garage/lighting/dim/#",
	"home/garage/lighting/dim/lumen",
	"HOME/#",
	"HOME/GARAGE/#",
	"HOME/GARAGE/LIGHTING/#",
	"HOME/GARAGE/LIGHTING/DIM/#",
	"HOME/GARAGE/LIGHTING/DIM/LUMEN",	
	"Home/#",
	"Home/Garage/#",
	"Home/Garage/Lighting/#",
	"Home/Garage/Lighting/Dim/#",
	"Home/Garage/Lighting/Dim/Lumen",
}

func main() {

	for i:=0; i< len(topic); i++{
		//fmt.Println(topic[i])

	//selectTopic, _ := strconv.Atoi(os.Args[1])

	//mqtt.DEBUG = log.New(os.Stderr, "DEBUG ", log.Ltime)
	clientID := "manuNoob0000"+strconv.Itoa(i)

	//create a ClientOptions
	opts := mqtt.NewClientOptions().
		AddBroker("tcp://broker.hivemq.com:1883").
		SetClientID(clientID).
		SetDefaultPublishHandler(messageHandler).
		SetConnectionLostHandler(connLostHandler)

	opts.OnConnect = func(c mqtt.Client) {
		fmt.Printf("############### SUBSCRIPT["+strconv.Itoa(i+1)+"]  ##############\n")
		fmt.Printf("Client connected, subscribing to: " + topic[i] + "\n")

		//Subscribe here, otherwise after connection lost,
		//you may not receive any message
		if token := c.Subscribe(topic[i], 0, nil); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}
		
	}

	//create and start a client using the above ClientOptions
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	time.Sleep(40000 * time.Millisecond)
	c.Disconnect(250)
	if count > 1 {
		fmt.Println("Find topics : "+strconv.Itoa(count))
	}else if count == 1 {
		fmt.Println("Find topics : "+strconv.Itoa(count))
		fmt.Println("Lastopic topics : "+lastTopic)
		//os.Exit(1)
	}else{
		fmt.Println("Find topics : "+strconv.Itoa(count))
	}
	count = 0
}
	// for {
	// 	//Lazy...
	// 	time.Sleep(500 * time.Millisecond)
	// }Î

}
