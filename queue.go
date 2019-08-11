package main

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/pubsub"
)

var (
	client *pubsub.Client
	topic  *pubsub.Topic
)

func initQueue(ctx context.Context) {

	c, e := pubsub.NewClient(ctx, projectID)
	if e != nil {
		logger.Fatalf("Error creating PubSub client: %v", e)
	}
	client = c

	t := c.Topic(topicName)
	topicExists, _ := t.Exists(ctx)

	if !topicExists {
		logger.Printf("Topic %s not found, creating...", topic)
		newTop, err := c.CreateTopic(ctx, topicName)
		if err != nil {
			logger.Fatalf("Unable to create topic: %s - %v", topic, err)
		}
		topic = newTop
	}
	topic = t
}

func store(ctx context.Context, event *SimpleEvent) error {

	b, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("Error while marshaling event: %v", err)
	}

	msg := &pubsub.Message{Data: b}
	result := topic.Publish(ctx, msg)
	_, err = result.Get(ctx)
	return err
}
