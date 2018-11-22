package fn

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// GitHubEventHandler handles the GitHub WebHook call
func GitHubEventHandler(w http.ResponseWriter, r *http.Request) {

	once.Do(func() {
		configInitializer("GitHubEventHandler")
	})

	w.Header().Set("Content-type", "application/json")

	se, err := parseGitHubWebHook([]byte(secret), r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Failed processing hook ('%s')", err)
		io.WriteString(w, "{}")
		return
	}

	//TODO: parse se (SimpleEvent)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(se)

	return
}
