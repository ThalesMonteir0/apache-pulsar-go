package main

import (
	"apache-pulsar-product-consumer/internal/consumer"
	"apache-pulsar-product-consumer/internal/producer"
)

func main() {
	forever := make(chan bool)

	client := producer.OpenConnectPulsar()
	defer client.Close()

	go consumer.ConsumerPulsar(client, "teste", "teste-1", 1)
	go consumer.ConsumerPulsar(client, "teste", "teste-1", 2)

	<-forever

}
