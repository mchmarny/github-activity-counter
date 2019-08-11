package env

import (
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	logger = log.New(os.Stdout, "", 0)
)

// MustGetEnvVar gets set environment variable or fails if fallbackValue i snot set
func MustGetEnvVar(key, fallbackValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		logger.Printf("%s: %s", key, val)
		return strings.TrimSpace(val)
	}

	if fallbackValue == "" {
		logger.Fatalf("Required envvar not set: %s", key)
	}

	logger.Printf("%s: %s (not set, using default)", key, fallbackValue)
	return fallbackValue
}

// MustGetIntEnvVar gets set environment variable or fails if fallbackValue i snot set
func MustGetIntEnvVar(key string, fallbackValue int) int {
	if val, ok := os.LookupEnv(key); ok {
		logger.Printf("%s: %s", key, val)

		port, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalf("failed to parse %s value (%s): %v", key, val, err)
		}
		return port
	}
	logger.Printf("%s: %d (not set, using default)", key, fallbackValue)
	return fallbackValue
}
