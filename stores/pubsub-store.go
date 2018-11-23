package stores

import (
	"errors"
	"log"
	"context"

	"encoding/json"
	"fmt"
	"cloud.google.com/go/pubsub"

	"github.com/mchmarny/github-activity-counter/types"
)

// PubSubStore pushes events to pubsub topic
type PubSubStore struct {
	ProjectID string
	TopicName string
	Ctx context.Context

	client *pubsub.Client
	topic   *pubsub.Topic
}

// Initialize is invoked once per Storable life cycle to configure the store
func (s *PubSubStore) Initialize() error {
	log.Print("Initializing PubSubStore...")

	if s.ProjectID == "" {
		return errors.New("ProjectID not set")
	}

	if s.TopicName == "" {
		return errors.New("TopicName not set")
	}

	if s.Ctx == nil {
		return errors.New("Ctx not set")
	}

	c, err := pubsub.NewClient(s.Ctx, s.ProjectID)
	if err != nil {
		log.Panicf("Failed to create client: %v", err)
	}


	s.client = c
	s.topic = s.client.Topic(s.TopicName)

	return nil
}

// Store persist the SimpleEvent
func (s *PubSubStore) Store(event *types.SimpleEvent) error {

	b, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("Error while marshaling event: %v", err)
	}

	msg := &pubsub.Message{Data: b}
	result := s.topic.Publish(s.Ctx, msg)
	_, err = result.Get(s.Ctx)
	if err != nil {
		return fmt.Errorf("Error while publishing message: %v", err)
	}

	return nil

}
