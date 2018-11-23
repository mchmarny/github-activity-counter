package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/mchmarny/github-activity-counter/handler"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, fmt.Sprintf("{'url': '%s'}", r.URL.Path[1:]))
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/events", handler.GitHubEventHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
