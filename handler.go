package fn

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// GitHubEventHandler handles the GitHub WebHook call
func GitHubEventHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	once.Do(func() {
		if err := configInitializer(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Error while initializing configuration: %v", err)
			io.WriteString(w, "{}")
			return
		}
	})

	se, err := parseGitHubWebHook([]byte(webHookSecret), r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error while processing WebHook: %v", err)
		io.WriteString(w, "{}")
		return
	}

	//TODO: parse se (SimpleEvent)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(se)

	return
}
