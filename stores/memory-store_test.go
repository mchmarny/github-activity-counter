package stores

import (
	"testing"
	"time"

	"github.com/mchmarny/github-activity-counter/types"
)

func TestDefaultConfigInitializer(t *testing.T) {

	store := &InMemoryStore{}

	err := store.Initialize()
	if err != nil {
		t.Errorf("Error while initializing InMemoryStore: %v", err)
	}

	ev1 := &types.SimpleEvent{
		ID:        "1",
		Type:      "push",
		Countable: true,
		EventAt:   time.Now(),
		Repo:      "test",
		Actor:     "me",
	}

	err = store.Store(ev1)
	if err != nil {
		t.Errorf("Error while storing event: %v", err)
	}

}
