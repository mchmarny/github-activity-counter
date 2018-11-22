package fn

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestParsingIssueComment(t *testing.T) {
	validateSimpleEvent("samples/issue_comment.json", "issue_comment", t)
}

func TestParsingIssues(t *testing.T) {
	validateSimpleEvent("samples/issues.json", "issues", t)
}

func TestCommitComment(t *testing.T) {
	validateSimpleEvent("samples/commit_comment.json", "commit_comment", t)
}

func TestPullRequestReviewComment(t *testing.T) {
	validateSimpleEvent("samples/pull_request_review_comment.json",
		"pull_request_review_comment", t)
}

func TestPullRequest(t *testing.T) {
	validateSimpleEvent("samples/pull_request.json", "pull_request", t)
}

func TestPullRequestReview(t *testing.T) {
	validateSimpleEvent("samples/pull_request_review.json", "pull_request_review", t)
}

func validateSimpleEvent(testDataPath string, issueType string, t *testing.T) {

	jf, err := os.Open(testDataPath)
	if err != nil {
		t.Errorf("Error while opening %s: %v", testDataPath, err)
	}
	defer jf.Close()
	data, _ := ioutil.ReadAll(jf)

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
