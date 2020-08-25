package Library

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"sync"
	"time"
)

func Connect(clientId string, uri string) mqtt.Client {
	opts := createClientOptions(clientId, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(time.Nanosecond / 999999) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func createClientOptions(clientId string, uri string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(uri)
	opts.SetClientID(clientId)
	return opts
}

func Listen(wg *sync.WaitGroup, uri, topic string, client mqtt.Client, file os.File) {
	//client := connect("sub", uri)
	for {
		client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
			writeToCsv(file, string(msg.Payload()))
		})
	}
}
