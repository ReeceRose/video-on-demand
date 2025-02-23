package main

import (
	"context"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("nats://127.0.0.1:4222")
	if err != nil {
		log.Fatalf("NATS: Failed to connect to NATS server: %v", err)
	}
	defer nc.Drain()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatalf("NATS: Failed to create JetStream context: %v", err)
	}

	streamName := "EVENTS"
	subject := "events.>"

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     streamName,
		Subjects: []string{subject},
	})
	if err != nil {
		log.Fatalf("NATS: Failed to create stream: %v", err)
	}

	_, err = js.AddConsumer(streamName, &nats.ConsumerConfig{
		Durable:   "consumer",
		AckPolicy: nats.AckExplicitPolicy,
	})
	if err != nil {
		log.Fatalf("NATS: Failed to create consumer: %v", err)
	}

	sub, err := js.PullSubscribe(subject, "consumer")
	if err != nil {
		log.Fatalf("NATS: Failed to subscribe: %v", err)
	}

	log.Println("Listening for messages...")

	for {
		msgs, err := sub.Fetch(10, nats.Context(nats.Context(context.Background())))
		if err != nil {
			log.Printf("NATS: Error fetching messages: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}

		for _, msg := range msgs {
			log.Printf("Received message: '%s' on subject '%s'", string(msg.Data), msg.Subject)
			msg.Ack()
		}
	}
}
