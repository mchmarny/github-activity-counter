package counter

import (
	"github.com/mchmarny/github-activity-counter/types"
)

// Storable represents simple event store
type Storable interface {
	// Initialize is invoked once per Storable life cycle to configure the store
	Initialize() error
	// Store persist the SimpleEvent
	Store(event *types.SimpleEvent) error
}
