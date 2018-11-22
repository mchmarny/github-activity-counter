package fn

import (
	"context"
	"log"
	"os"
	"sync"
)

const (
	// sha1Prefix is the prefix used by GitHub before the HMAC hexdigest.
	sha1Prefix = "sha1"
	// signatureHeader is the GitHub header key used to pass the HMAC hexdigest.
	signatureHeader = "X-Hub-Signature"
	// eventTypeHeader is the GitHub header key used to pass the event type.
	eventTypeHeader = "X-Github-Event"
	// deliveryIDHeader is the GitHub header key used to pass the unique ID for the webhook event.
	deliveryIDHeader = "X-Github-Delivery"
)

var (
	ctx               context.Context
	once              sync.Once
	region            string
	projectID         string
	secret            string
	configInitializer = defaultConfigInitializer
)

func defaultConfigInitializer(fn string) {

	log.Printf("%s configuration using: defaultConfigInitializer", fn)

	ctx = context.Background()

	secret = os.Getenv("HOOK_SECRET")
	if secret == "" {
		log.Fatalf("HOOK_SECRET environment variable not set")
	}

	projectID = os.Getenv("GCP_PROJECT")
	if projectID == "" {
		log.Printf("GCP_PROJECT environment variable not set")
		projectID = "NOT_SET"
	}

	region = os.Getenv("FUNCTION_REGION")
	if region == "" {
		log.Printf("FUNCTION_REGION not set, using default: us-central1")
		region = "us-central1"
	}

}
