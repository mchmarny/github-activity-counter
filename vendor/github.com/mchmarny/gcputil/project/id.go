package project

import (
	"log"
	"os"
	"strings"

	"github.com/mchmarny/gcputil/meta"
)

var (
	logger      = log.New(os.Stdout, "", 0)
	agentName   = "gcputil"
	projectKeys = []string{
		"GCP_PROJECT",
		"PROJECT",
		"PROJECT_ID",
		"GOOGLE_CLOUD_PROJECT",
		"GCLOUD_PROJECT",
	}
)

// GetIDOrFail checks common projectID env vars first,
// then tries to retreave the project ID from meta data server
// Fails if neither are not successful
func GetIDOrFail() string {
	p, err := deriveProjectID(agentName)
	if err != nil {
		logger.Fatalf("Error while getting project ID: %v", err)
	}
	return p
}

// GetID checks common projectID env vars first,
// then tries to retreave the project ID from meta data server
func GetID() (project string, err error) {
	return deriveProjectID(agentName)
}

func deriveProjectID(agent string) (p string, err error) {
	for _, key := range projectKeys {
		if val, ok := os.LookupEnv(key); ok {
			logger.Printf("Found %s: %s", key, val)
			return strings.TrimSpace(val), nil
		}
	}
	return meta.GetClient(agent).ProjectID()
}

// NumericProjectID returns the current instance's numeric project ID
func NumericProjectID(agent string) (p string, err error) {
	return meta.GetClient(agent).NumericProjectID()
}
