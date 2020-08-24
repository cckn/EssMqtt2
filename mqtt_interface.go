package main

//
//func Connect(clientId string, uri string) mqtt.Client {
//	opts := createClientOptions(clientId, uri)
//	client := mqtt.NewClient(opts)
//	token := client.Connect()
//	for !token.WaitTimeout(3 * time.Second) {
//	}
//	if err := token.Error(); err != nil {
//		log.Fatal(err)
//	}
//	return client
//}
//
//func createClientOptions(clientId string, uri string) *mqtt.ClientOptions {
//	opts := mqtt.NewClientOptions()
//	opts.AddBroker(uri)
//	opts.SetClientID(clientId)
//	return opts
//}
//
//
//func listen(wg *sync.WaitGroup,uri , topic string,client mqtt.Client,file os.File) {
//	//client := connect("sub", uri)
//	for {
//		client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
//			writeToCsv(file, string(msg.Payload()))
//		})
//	}
//}
//
//func Publish(wg *sync.WaitGroup,client mqtt.Client,topic string){
//	for {
//		client.Publish(topic, 0, false, "HanGang")
//		time.Sleep(time.Second*2)
//	}
//}
