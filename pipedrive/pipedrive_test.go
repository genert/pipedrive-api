package pipedrive

import (
	"flag"
	"os"
)

var (
	// Client is the Pipedrive client being tested.
	client *Client

	// Custom flag for integration tests
	apiIntegration = flag.Bool("api", false, "run API integration tests")
)

func init() {
	config := &Config{
		ApiKey: os.Getenv("PIPEDRIVE_API_KEY"),
	}

	client = New(config)
}
