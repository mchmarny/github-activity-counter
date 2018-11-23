package counter

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/mchmarny/github-activity-counter/stores"
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

	hookSecretEnvVarName = "HOOK_SECRET"

	projectIDEnvVarName = "GCP_PROJECT"

	pubsubTopicNameEnvVarName = "PUBSUB_EVENTS_TOPIC"
)

var (
	ctx               context.Context
	once              sync.Once
	webHookSecret     string
	configInitializer = defaultConfigInitializer
	store             Storable
)

func defaultConfigInitializer() error {

	log.Print("Initializing configuration using: defaultConfigInitializer")
	ctx = context.Background()

	webHookSecret = os.Getenv(hookSecretEnvVarName)
	if webHookSecret == "" {
		return fmt.Errorf("%s environment variable not set", hookSecretEnvVarName)
	}

	projectID := os.Getenv(projectIDEnvVarName)
	topicName := os.Getenv(pubsubTopicNameEnvVarName)

	// if project and topics are defined then pubsub else in-memory for testing
	if projectID == "" || topicName == "" {
		ims := &stores.InMemoryStore{}
		store = ims
	} else {
		pss := &stores.PubSubStore{
			ProjectID: projectID,
			TopicName: topicName,
			Ctx: ctx,
		}
		store = pss
	}

	err := store.Initialize()
	if err != nil {
		return err
	}

	return nil

}
