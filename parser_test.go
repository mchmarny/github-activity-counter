package main

import (
	"bytes"
	"net/http"
	"testing"
)

func TestParseWebHookWithInvalidSig(t *testing.T) {

	var testData = []byte("{}")
	r, _ := http.NewRequest("POST", "/", bytes.NewReader(testData))
	r.Header.Add(signatureHeader, "bad-key")
	r.Header.Add(eventTypeHeader, "not used in test")
	r.Header.Add(deliveryIDHeader, "not used in test")

	_, err := parseGitHubWebHook(r)
	if err == nil {
		t.Error("Expected invalid signature error")
	}

}

func TestParseWebHookHeaders(t *testing.T) {

	const testID = "1234"
	const testEventType = "issue_comment"
	const testFilePath = "sample/issue_comment.json"

	data, err := getFileContent(testFilePath)
	if err != nil {
		t.Errorf("Error while opening %s: %v", testFilePath, err)
	}

	key := []byte(webHookSecret)
	sig := makeNewSignature(key, data)

	r, _ := http.NewRequest("POST", "/", bytes.NewReader(data))
	r.Header.Add(signatureHeader, sig)
	r.Header.Add(eventTypeHeader, testEventType)
	r.Header.Add(deliveryIDHeader, testID)

	se, err := parseGitHubWebHook(r)
	if err != nil {
		t.Errorf("Error while parsing WebHook %v", err)
		return
	}

	if se.ID != testID {
		t.Errorf("Invalid data: Expected %s got %s", testID, se.ID)
	}

	if se.Type != testEventType {
		t.Errorf("Invalid data: Expected %s got %s", testEventType, se.Type)
	}

}

func TestParsingPush(t *testing.T) {
	validateSimpleEvent("sample/push.json", "push", t)
}

func TestParsingIssueComment(t *testing.T) {
	validateSimpleEvent("sample/issue_comment.json", "issue_comment", t)
}

func TestParsingIssues(t *testing.T) {
	validateSimpleEvent("sample/issues.json", "issues", t)
}

func TestCommitComment(t *testing.T) {
	validateSimpleEvent("sample/commit_comment.json", "commit_comment", t)
}

func TestPullRequestReviewComment(t *testing.T) {
	validateSimpleEvent("sample/pull_request_review_comment.json",
		"pull_request_review_comment", t)
}

func TestPullRequest(t *testing.T) {
	validateSimpleEvent("sample/pull_request.json", "pull_request", t)
}

func TestPullRequestReview(t *testing.T) {
	validateSimpleEvent("sample/pull_request_review.json", "pull_request_review", t)
}

func TestNonCountableEvent(t *testing.T) {
	const testDataPath = "sample/delete.json"
	const issueType = "delete"
	data, err := getFileContent(testDataPath)
	if err != nil {
		t.Errorf("Error while getting test file %s: %v", testDataPath, err)
	}

	var testID = "123"
	se, err := parseSimpleEvent(data, testID, issueType)
	if err != nil {
		t.Errorf("Error while parsing %s: %v", issueType, err)
	}

	if se == nil {
		t.Errorf("Unable to parse SimpleEvent from %s: %v", issueType, err)
	}

	if se.Countable {
		t.Errorf("Expected non-countable event from %s: %v", issueType, err)
	}

}

func validateSimpleEvent(testDataPath string, issueType string, t *testing.T) {

	data, err := getFileContent(testDataPath)
	if err != nil {
		t.Errorf("Error while getting test file %s: %v", testDataPath, err)
	}

	var testID = "123"
	se, err := parseSimpleEvent(data, testID, issueType)
	if err != nil {
		t.Errorf("Error while parsing %s: %v", issueType, err)
	}

	if se.ID != testID {
		t.Errorf("Invalid data: Expected %s got %s", testID, se.ID)
	}

	if se.Type != issueType {
		t.Errorf("Invalid data: Expected %s got %s", issueType, se.Type)
	}

	if se.Repo == "" {
		t.Error("Invalid data: nil Repo")
	}

	if se.Actor == "" {
		t.Error("Invalid data: nil Actor")
	}

	if se.EventAt.IsZero() {
		t.Error("Invalid data: nil EventAt")
	}

}
