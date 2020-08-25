package main

import (
	"EssMqtt/Library"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	const (
		uri      = "tcp://broker.hivemq.com:1883"
		topic    = "1234"
		fileName = "test07.csv"
	)

	//	file, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)

	file, error := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE, 0666)
	if error != nil {
		panic(error)
	}

	defer file.Close()

	println("@@MQTTPublish Start@@")
	client := Library.Connect("aassdd", uri)

	wg.Add(1)
	go Library.Listen(&wg, uri, topic, client, *file)
	//	go Publish(&wg, client, topic)
	wg.Wait()
}
