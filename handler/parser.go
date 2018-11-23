package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/mchmarny/github-activity-counter/handler/types"
)

func parseGitHubWebHook(secret []byte, req *http.Request) (*types.SimpleEvent, error) {

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

func parseSimpleEvent(body []byte, eventID, eventType string) (*types.SimpleEvent, error) {

	// placeholder for returned struct
	se := &types.SimpleEvent{
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
		ev := &types.CommentEvent{}
		err := json.Unmarshal(body, &ev)
		if err != nil {
			log.Printf("Error parsing %s: %v", se.Type, err)
			return nil, err
		}
		se.Actor = ev.Comment.User.Name
		se.EventAt = ev.Comment.CreatedAt
		se.Repo = ev.Repo.Name

	case "issues":
		ev := &types.IssuesEvent{}
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
		ev := &types.PullRequestEvent{}
		err := json.Unmarshal(body, &ev)
		if err != nil {
			log.Printf("Error parsing %s: %v", se.Type, err)
			return nil, err
		}
		se.Actor = ev.PR.User.Name
		se.EventAt = ev.PR.CreatedAt
		se.Repo = ev.Repo.Name

	case "pull_request_review":
		ev := &types.ReviewEvent{}
		err := json.Unmarshal(body, &ev)
		if err != nil {
			log.Printf("Error parsing %s: %v", se.Type, err)
			return nil, err
		}
		se.Actor = ev.Review.User.Name
		se.EventAt = ev.Review.CreatedAt
		se.Repo = ev.Repo.Name

	case "push":
		ev := &types.SimplePushEvent{}
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
