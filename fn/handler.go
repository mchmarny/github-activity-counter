package fn

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func parseGitHubWebHook(secret []byte, req *http.Request) (*gitHubWebHook, error) {
	hc := gitHubWebHook{}

	if hc.sig = req.Header.Get("x-hub-signature"); len(hc.sig) == 0 {
		return nil, errors.New("No signature")
	}

	if hc.event = req.Header.Get("x-github-event"); len(hc.event) == 0 {
		return nil, errors.New("No event")
	}

	if hc.id = req.Header.Get("x-github-delivery"); len(hc.id) == 0 {
		return nil, errors.New("No event ID")
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	if !checkContentSignature(secret, hc.sig, body) {
		return nil, errors.New("Invalid signature")
	}

	hc.content = body

	return &hc, nil
}

// GitHubEventHandler handles the GitHub WebHook call
func GitHubEventHandler(w http.ResponseWriter, r *http.Request) {

	once.Do(func() {
		configInitializer("GitHubEventHandler")
	})

	hc, err := parseGitHubWebHook([]byte(secret), r)

	w.Header().Set("Content-type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Failed processing hook ('%s')", err)
		io.WriteString(w, "{}")
		return
	}

	log.Printf("Received: %s", hc.event)

	//TODO: parse `hc.Payload` for counter

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "{}")

	return
}
