package fn

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func parseGitHubWebHook(secret []byte, req *http.Request) (*SimpleEvent, error) {

	var sig string
	if sig = req.Header.Get(signatureHeader); len(sig) == 0 {
		return nil, errors.New("No signature")
	}
	log.Printf("Sig: %s", sig)

	eventType := req.Header.Get(eventTypeHeader)
	if eventType == "" {
		return nil, errors.New("No event")
	}
	log.Printf("Type: %s", eventType)

	eventID := req.Header.Get(deliveryIDHeader)
	if eventID == "" {
		return nil, errors.New("No event ID")
	}
	log.Printf("ID: %s", eventID)

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	if !checkContentSignature(secret, sig, body) {
		return nil, errors.New("Invalid signature")
	}

	return parseSimpleEvent(body, eventID, eventType)

}

func parseSimpleEvent(body []byte, eventID, eventType string) (*SimpleEvent, error) {

	// placeholder for returned struct
	se := &SimpleEvent{
		Type:      eventType,
		ID:        eventID,
		Raw:       (json.RawMessage)(body),
		Countable: true,
	}

	// TODO: Could this be cast into some kind of countable type?
	switch se.Type {

	case "issue_comment":
		fallthrough
	case "commit_comment":
		ev := &CommentEvent{}
		err := json.Unmarshal(body, &ev)
		if err != nil {
			log.Printf("Error parsing %s: %v", se.Type, err)
			return nil, err
		}
		se.Actor = ev.Comment.User.Name
		se.EventAt = ev.Comment.CreatedAt
		se.Repo = ev.Repo.Name

	case "issues":
		ev := &IssuesEvent{}
		err := json.Unmarshal(body, &ev)
		if err != nil {
			log.Printf("Error parsing issues: %v", err)
			return nil, err
		}
		se.Actor = ev.Issue.User.Name
		se.EventAt = ev.Issue.CreatedAt
		se.Repo = ev.Repo.Name

	case "pull_request":
		fallthrough
	case "pull_request_review_comment":
		ev := &PullRequestEvent{}
		err := json.Unmarshal(body, &ev)
		if err != nil {
			log.Printf("Error parsing %s: %v", se.Type, err)
			return nil, err
		}
		se.Actor = ev.PR.User.Name
		se.EventAt = ev.PR.CreatedAt
		se.Repo = ev.Repo.Name

	case "pull_request_review":
		ev := &ReviewEvent{}
		err := json.Unmarshal(body, &ev)
		if err != nil {
			log.Printf("Error parsing %s: %v", se.Type, err)
			return nil, err
		}
		se.Actor = ev.Review.User.Name
		se.EventAt = ev.Review.CreatedAt
		se.Repo = ev.Repo.Name

	case "push":
		ev := &SimplePushEvent{}
		err := json.Unmarshal(body, &ev)
		if err != nil {
			log.Printf("Error parsing %s: %v", se.Type, err)
			return nil, err
		}
		se.Actor = ev.User.Name
		// There is no push time, using WebHook execution time
		se.EventAt = time.Now()
		se.Repo = ev.Repo.Name

	default:
		se.Countable = false
		log.Printf("Uncountable type: %s", se.Type)
	}

	return se, nil

}
