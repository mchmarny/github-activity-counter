package stores

import (
	"log"

	"github.com/mchmarny/github-activity-counter/types"
)

// InMemoryStore represents simple event store
type InMemoryStore struct {
}

// Initialize is invoked once per Storable life cycle to configure the store
func (s *InMemoryStore) Initialize() error {
	log.Print("Initializing InMemoryStore...")
	return nil
}

// Store persist the SimpleEvent
func (s *InMemoryStore) Store(event *types.SimpleEvent) error {

	log.Printf("Added Event [ID:%s Type:%s Repo:%s Actor:%s At:%s]",
		event.ID, event.Type, event.Repo, event.Actor, event.EventAt)

	return nil

}
