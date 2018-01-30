package integration

import (
	"github.com/genert/pipedrive-api/pipedrive"
	"os"
)

var (
	client *pipedrive.Client
)

func init() {
	token := os.Getenv("PIPEDRIVE_API_TOKEN")

	if token == "" {
		print("No API key found. Integration tests won't run!\n\n")
		os.Exit(1)
	} else {
		config := &pipedrive.Config{
			ApiKey: token,
		}

		client = pipedrive.NewClient(config)
	}
}
