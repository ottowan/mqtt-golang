package main

import (
	"fmt"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	logger "github.com/sirupsen/logrus"
)

func messageHandler(c mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func main() {

	//root/device/node/prop/attribute
	topic := []string{
		"home/alarm/kitchen/heat/celcius",
		"home/alarm/frontdoor/heat/celcius",
		"home/alarm/bedroom/smoke/carbon",
		"home/tv/livingroom/on_off/time",
		"home/tv/bedroom/on_off/time",
		"home/tv/livingroom/record/time",
		"home/tv/bedroom/record/time",
		"home/aircondition/bedroom/on_off/time",
		"home/aircondition/bedroom/temperature/celcius",
		"home/aircondition/livingroom/on_off/time",
		"home/aircondition/livingroom/humandity/percent",
		"home/aircondition/livingroom/temperature/celcius",
		"home/speaker/livingroom/on_off/time",
		"home/speaker/livingroom/up_down/decibel",
		"home/camera/frontdoor/facedetect/bit",
		"home/camera/frontdoor/record/time",
		"home/lighting/frontdoor/on_off/time",
		"home/GARAGE/lighting/on_off/time",
		"home/garage/bedroom/on_off/time",
		"home/garage/kitchen/on_off/time",
		"home/garage/lighting/dim/lumen",
		"home/garage/lighting/on_off/time",
		"home/fan/livingroom/on_off/time",
		"home/projector/livingroom/on_off/time",
		"home/blub/livingroom/on_off/time",
		"myhouse/kitchen/oven",
		"house/living-room-light/set",
		"home/bathroom/humidity",
		"myhouse/kitchen/lights",
		"myhouse/garage/#",
		"home/bathroom/humidity",
		"home/bedroom/temperature 21",
		"home/bedroom/temperature/units C",
		"myhouse/garage/lights/doorlight",
		"light-switch/status",
		"house/sensor1",
		"myhouse/garage/lights/ceilinglight",
		"light-switch",
		"house/sensor2",
		"/myhouse/garage/garagedoor",
		"light-switch/set",
		"house/sensor10",
		"sensor1/stat",
		"knxgateway1/status/Kitchen/Lights/Front Left",
		"hm/status/Light Kitchen Front Left/LEVEL",
		"toplevelname/status/itemname",
		"toplevelname/get/itemname",
		"/zigbee/0013a20040410034/adc7",
		"hm/set/light/state",
		"/home/door/sensor/battery",
	}

	for i := 0; i < len(topic); i++ {
		clientID := "thecobNoob0000" + strconv.Itoa(i)
		opts := mqtt.NewClientOptions()
		opts.AddBroker("tcp://broker.hivemq.com:1883")
		opts.SetClientID(clientID)
		message := "message from " + clientID
		opts.SetKeepAlive(time.Second * time.Duration(60))
		// If lost connection, reconnect again
		opts.SetConnectionLostHandler(func(client mqtt.Client, e error) {
			logger.Warn(fmt.Sprintf("Connection lost : %v", e))
		})

		// connect to broker
		client := mqtt.NewClient(opts)
		token := client.Connect()
		if token.Wait() && token.Error() != nil {
			logger.Fatalf("Fail to connect broker, %v", token.Error())
		}

		// publish to topic
		token = client.Publish(topic[i], byte(0), false, message)
		if token.Wait() && token.Error() != nil {
			logger.Errorf("Fail to publish, %v", token.Error())
		}

		fmt.Println("Sending round[" + strconv.Itoa(1) + "]topic [" + strconv.Itoa(i+1) + "] : " + topic[i])

	}

}
