package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	ev "github.com/mchmarny/gcputil/env"
	mt "github.com/mchmarny/gcputil/metric"
	pj "github.com/mchmarny/gcputil/project"
)

var (
	logger        = log.New(os.Stdout, "", 0)
	projectID     = pj.GetIDOrFail()
	port          = ev.MustGetEnvVar("PORT", "8080")
	topicName     = ev.MustGetEnvVar("TOPIC", "preprocessd")
	webHookSecret = ev.MustGetEnvVar("HOOK_SECRET", "")
	metricsClient *mt.Client
)

func main() {

	ctx := context.Background()

	m, err := mt.NewClient(ctx)
	if err != nil {
		logger.Fatalf("Error creating metrics client : %v", err)
	}
	metricsClient = m

	initQueue(ctx)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})

	http.HandleFunc("/v1/github", gitHubEventHandler)

	hostPort := net.JoinHostPort("0.0.0.0", port)

	if err := http.ListenAndServe(hostPort, nil); err != nil {
		logger.Fatal(err)
	}

}
