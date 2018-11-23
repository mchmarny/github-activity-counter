package types

import (
	"encoding/json"
	"time"
)

// SimpleEvent represents a simple version of the gitHub Event
type SimpleEvent struct {
	ID        string          `json:"id,omitempty"`
	Type      string          `json:"type,omitempty"`
	Raw       json.RawMessage `json:"raw,omitempty"`
	Countable bool            `json:"countable,omitempty"`
	EventAt   time.Time       `json:"event_time,omitempty"`
	Repo      string          `json:"repo,omitempty"`
	Actor     string          `json:"actor,omitempty"`
}
