package main

import (
	"apache-pulsar-product-consumer/internal/producer"
	"fmt"
)

func main() {
	client := producer.OpenConnectPulsar()
	defer client.Close()

	for i := 0; i < 10; i++ {
		err := producer.PulsarProducer(client, fmt.Sprintf("teste-%d", i), "teste")
		if err != nil {
			fmt.Printf("Error opening producer connection: %v\n", err)
		}
	}

}
