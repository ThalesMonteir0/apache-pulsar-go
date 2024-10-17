package producer

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"log"
)

func OpenConnectPulsar() pulsar.Client {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://localhost:6650",
	})
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	return client
}

func PulsarProducer(client pulsar.Client, msgString, topic string) error {
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
	})
	if err != nil {
		return err
	}

	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte(msgString),
	})

	defer producer.Close()

	if err != nil {
		fmt.Printf("Error sending message to topic %s: %v\n", topic, err)
		return err
	}

	fmt.Printf("mensagem produzida")

	return nil
}
