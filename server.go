package main

import (
	"log"
	"net/http"

	"github.com/mchmarny/ghcounter/fn"
)

func main() {
	http.HandleFunc("/", fn.GitHubEventHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
