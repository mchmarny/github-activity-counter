package fn

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mchmarny/github-activity-counter/types"
)

func TestGitHubEventHandler(t *testing.T) {

	const testID = "1234"
	const testEventType = "issue_comment"
	const testFilePath = "samples/issue_comment.json"
	const testSecret = "some-super-long-secret-string"

	data, err := getFileContent(testFilePath)
	if err != nil {
		t.Errorf("Error while opening %s: %v", testFilePath, err)
	}

	key := []byte(testSecret)
	sig := makeNewSignature(key, data)

	r, _ := http.NewRequest("POST", "/", bytes.NewReader(data))
	r.Header.Add(signatureHeader, sig)
	r.Header.Add(eventTypeHeader, testEventType)
	r.Header.Add(deliveryIDHeader, testID)
	if err != nil {
		t.Errorf("Error while creating request %s: %v", testFilePath, err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GitHubEventHandler)
	handler.ServeHTTP(rr, r)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v expected %v",
			status, http.StatusOK)
	}

	body, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Errorf("Error while reading response body %v", err)
	}

	ev := &types.SimpleEvent{}
	err = json.Unmarshal(body, &ev)
	if err != nil {
		t.Errorf("Error while unmarshaling SimpleEvent from body %v", err)
	}

	if ev.ID != testID {
		t.Errorf("Invalid SimpleEvent ID: got: %s expected %s", ev.ID, testID)
	}

}
