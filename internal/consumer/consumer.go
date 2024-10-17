package consumer

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"log"
)

func ConsumerPulsar(client pulsar.Client, topic, subscriptionName string) {
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic,
		SubscriptionName: subscriptionName,
		Type:             pulsar.Shared,
	})
	if err != nil {
		log.Fatalf("Failed to subscribe to topic %s on topic %s with error: %s", topic, topic, err)
	}

	defer consumer.Close()

	for {
		msg, err := consumer.Receive(context.Background())
		if err != nil {
			log.Fatalf("Failed to receive message: %v", err)
		}

		fmt.Printf("msgID: %d. Msg payload: %s", msg.ID(), string(msg.Payload()))
		consumer.Ack(msg)
	}
}
