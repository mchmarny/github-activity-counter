package meta

import (
	"log"
	"net/http"
	"os"

	m "cloud.google.com/go/compute/metadata"
)

var (
	logger = log.New(os.Stdout, "", 0)
)

// GetClient returns the raw metadata client
func GetClient(agent string) *m.Client {
	return m.NewClient(&http.Client{
		Transport: userAgentTransport{
			userAgent: agent,
			base:      http.DefaultTransport,
		},
	})
}

type userAgentTransport struct {
	userAgent string
	base      http.RoundTripper
}

// RoundTrip implements the transport interface
// // https://godoc.org/cloud.google.com/go/compute/metadata#example-NewClient
func (t userAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", t.userAgent)
	return t.base.RoundTrip(req)
}
