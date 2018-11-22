package fn

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
	EventAt   time.Time       `json:"avent_time,omitempty"`
	Repo      string          `json:"repo,omitempty"`
	Actor     string          `json:"actor,omitempty"`
}

// ReviewEvent is simple version of the GitHub structure
type ReviewEvent struct {
	Review ReviewActivity `json:"review"`
	Repo   SimpleRepo     `json:"repository"`
}

// PullRequestEvent is simple version of the GitHub structure
type PullRequestEvent struct {
	PR   SimpleUserActivity `json:"pull_request"`
	Repo SimpleRepo         `json:"repository"`
}

// IssuesEvent is simple version of the GitHub structure
type IssuesEvent struct {
	Issue SimpleUserActivity `json:"issue"`
	Repo  SimpleRepo         `json:"repository"`
}

// CommentEvent is simple version of the GitHub structure
type CommentEvent struct {
	Comment SimpleUserActivity `json:"comment"`
	Repo    SimpleRepo         `json:"repository"`
}

// SimpleUserActivity is simple version of the GitHub structure
type SimpleUserActivity struct {
	User      SimpleUser `json:"user"`
	CreatedAt time.Time  `json:"updated_at"`
}

// ReviewActivity is simple version of the GitHub structure
type ReviewActivity struct {
	User      SimpleUser `json:"user"`
	CreatedAt time.Time  `json:"submitted_at"`
}

// SimpleRepo is simple version of the GitHub structure
type SimpleRepo struct {
	Name string `json:"full_name"`
}

// SimpleUser is simple version of the GitHub structure
type SimpleUser struct {
	Name string `json:"login"`
}
