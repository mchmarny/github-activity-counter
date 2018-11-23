package stores

import (
	"errors"
	"log"

	"github.com/mchmarny/github-activity-counter/types"
)

// InMemoryStore represents simple event store
type InMemoryStore struct {
	data map[string]*types.SimpleEvent
}

// Initialize is invoked once per Storable life cycle to configure the store
func (s *InMemoryStore) Initialize(args map[string]interface{}) error {
	log.Print("Initializing InMemoryStore...")

	s.data = map[string]*types.SimpleEvent{}

	return nil
}

// Store persist the SimpleEvent
func (s *InMemoryStore) Store(event *types.SimpleEvent) error {

	if s.data == nil {
		return errors.New("InMemoryStore not initialized")
	}

	log.Printf("Events before: %d", len(s.data))
	s.data[event.ID] = event

	for k, v := range s.data {
		log.Printf("Added Event [ID:%s Type:%s Repo:%s Actor:%s At:%s]",
			k, v.Type, v.Repo, v.Actor, v.EventAt)
	}

	log.Printf("Events after:  %d", len(s.data))

	return nil

}
