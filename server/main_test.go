package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGitHubEventHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("Error while creating request %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(defaultHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v expected %v",
			status, http.StatusOK)
	}

	_, err = ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Errorf("Error while reading response body %v", err)
	}

}
