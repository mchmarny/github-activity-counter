package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mchmarny/github-activity-counter/handler"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/events", handler.GitHubEventHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
