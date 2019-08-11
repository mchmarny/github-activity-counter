package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	mt "github.com/mchmarny/gcputil/metric"
)

func initTests() {

	ctx := context.Background()

	m, err := mt.NewClient(ctx)
	if err != nil {
		logger.Fatalf("Error creating metrics client : %v", err)
	}
	metricsClient = m

	initQueue(ctx)
}

func TestGitHubEventHandler(t *testing.T) {

	initTests()

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
	if err != nil {
		t.Errorf("Error while creating request %s: %v", testFilePath, err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(gitHubEventHandler)
	handler.ServeHTTP(rr, r)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v expected %v",
			status, http.StatusOK)
	}

	body, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Errorf("Error while reading response body %v", err)
	}

	ev := &SimpleEvent{}
	err = json.Unmarshal(body, &ev)
	if err != nil {
		t.Errorf("Error while unmarshaling SimpleEvent from body %v", err)
	}

	if ev.ID != testID {
		t.Errorf("Invalid SimpleEvent ID: got: %s expected %s", ev.ID, testID)
	}

}
