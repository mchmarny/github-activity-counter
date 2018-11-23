package handler

import (
	"github.com/mchmarny/github-activity-counter/handler/types"
)

// Storable represents simple event store
type Storable interface {
	// Initialize is invoked once per Storable life cycle to configure the store
	Initialize(args map[string]interface{}) error
	// Store persist the SimpleEvent
	Store(event *types.SimpleEvent) error
}
