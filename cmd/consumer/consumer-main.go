package main

import (
	"apache-pulsar-product-consumer/internal/consumer"
	"apache-pulsar-product-consumer/internal/producer"
)

func main() {
	client := producer.OpenConnectPulsar()
	defer client.Close()

	consumer.ConsumerPulsar(client, "teste", "teste-1")

}
