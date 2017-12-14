package main

// See: https://github.com/GoogleCloudPlatform/golang-samples/blob/master/pubsub/topics/main.go#L132-L148

import (
	"fmt"
	"log"

	"golang.org/x/net/context"

	"cloud.google.com/go/pubsub"
)

func main() {
	fmt.Println("Starting")

	ctx := context.Background()

	projectID := "testpubsubstuff"
	topicName := "my-test-topic"

	// Setup client
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		panic(err)
	}

	// Publish a message
	err = publish(client, topicName, "This is a foo")
	if err != nil {
		log.Fatalf("Failed to publish: %v", err)
	}

	fmt.Println("Ending")
}

func publish(client *pubsub.Client, topic string, msg string) error {
	ctx := context.Background()

	t := client.Topic(topic)

	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})

	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("Published a message; msg ID: %v\n", id)

	return nil
}
