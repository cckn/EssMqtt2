package main

import (
	"encoding/csv"
	"fmt"
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
	opts.SetClientID(clientId)
	return opts
}

func listen(wg *sync.WaitGroup, uri, topic string, client mqtt.Client, file os.File) {
	//client := connect("sub", uri)
	for {
		client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
			writeToCsv(file, string(msg.Payload()))
		})
	}
}

func Publish(wg *sync.WaitGroup, client mqtt.Client, topic string) {

	for {
		client.Publish(topic, 2, false, "HanGang")
		time.Sleep(time.Second * 2)

	}
}

func writeToCsv(file os.File, msg string) {
	writer := csv.NewWriter(&file)
	defer writer.Flush()

	record := []string{time.Now().String(), msg}
	fmt.Println(record)
	writer.Write(record)
}

func main() {
	var wg sync.WaitGroup

	const (
		uri      = "tcp://broker.hivemq.com:1883"
		topic    = "test/topic12/17"
		fileName = "test07.csv"
	)

	file, _ := os.Create(fileName)
	defer file.Close()

	println("@@MQTTPublish Start@@")
	client := Connect("qwertyuiop", uri)

	wg.Add(2)

	go listen(&wg, uri, topic, client, *file)
	go Publish(&wg, client, topic)

	wg.Wait()
}
