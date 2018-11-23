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

	storageTypeEnvVarName = "STORE_TYPE"
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

	storeInitArgs := map[string]interface{}{
		"TBDLater": 1,
	}

	storageType := os.Getenv(storageTypeEnvVarName)
	if storageType == "" {

		ims := &stores.InMemoryStore{}
		store = ims
		err := store.Initialize(storeInitArgs)
		if err != nil {
			return err
		}
	}

	return nil

}
