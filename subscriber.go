package main

// See: https://github.com/GoogleCloudPlatform/golang-samples/blob/master/pubsub/subscriptions/main.go#L87-L130

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
	subName := "foo-subscription"

	// Setup client
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Setup topic
	topic := client.Topic(topicName)
	// ok, err := topic.Exists(ctx)
	_, err = topic.Exists(ctx)
	if err != nil {
		log.Fatalf("Failed to get topic: %v", err)
	}
	// if ok {
	// 	return t
	// }

	// TODO: Probably a good idea to create a subscription here?
	// "Subscriptions with no activity (push successes or pull requests) for 31 days may be deleted automatically. You can also delete a subscription manually. While you can create new subscriptions with the same name as a deleted one, note that the new subscription has no relation to the old one, even though they have the same name."

	err = pull(client, subName, topic)
	if err != nil {
		log.Fatalf("Failed to pull: %v", err)
	}

	fmt.Println("Ending")
}

func pull(client *pubsub.Client, name string, topic *pubsub.Topic) error {
	ctx := context.Background()

	sub := client.Subscription(name)
	sub.ReceiveSettings.MaxOutstandingMessages = 0

	fmt.Println("Listening..")
	err := sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Printf("Got message: %q\n", string(msg.Data))
		msg.Ack()
	})
	if err != nil {
		return err
	}

	return nil
}
