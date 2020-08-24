package Library

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"time"
)

func Connect(clientId string, uri string) mqtt.Client {
	println("MQTT Broker Access")
	opts := createClientOptions(clientId, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func createClientOptions(clientId string, uri string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(uri)
	//opts.SetUsername()
	//opts.SetPassword()
	opts.SetClientID(clientId)
	return opts
}
func cb() {
	println("123")
}
func Listen(client mqtt.Client, uri string, topic string, cb mqtt.MessageHandler) {
	//client := Connect("sub", uri)
	//var ss string = "dd"
	//client.Subscribe(topic, 0, cb)
	if token := client.Subscribe(topic, 0, cb); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

}
func MQTTPublish(client mqtt.Client, topic string, payload interface{}) {

	client.Publish(topic, 0, true, payload)
}
