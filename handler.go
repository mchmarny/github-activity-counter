package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func gitHubEventHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	se, err := parseGitHubWebHook(r)
	if err != nil {
		logger.Printf("Error while processing WebHook: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "{}")
		return
	}

	if se.Countable {
		err = store(r.Context(), se)
		if err != nil {
			logger.Printf("Error while storing event: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "{}")
			return
		}
	}

	err = metricsClient.Publish(r.Context(), se.Type, "github-event-counter", int64(1))
	if err != nil {
		logger.Printf("Error while publishing metrics: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "{}")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(se)

	return
}
