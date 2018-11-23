package types

import (
	"testing"
	"time"
)

func TestSimpleEventTypes(t *testing.T) {

	//TODO: Why? What is there to tests here?
	ev1 := &SimpleEvent{
		ID:        "1",
		Type:      "push",
		Countable: true,
		EventAt:   time.Now(),
		Repo:      "test",
		Actor:     "me",
	}

	ev2 := ev1

	if ev1.ID != ev2.ID {
		t.Error("Unable to compare simple event types")
	}

}
