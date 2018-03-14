package integration

import (
	"context"
	"testing"
)

func TestWebhooksService_List(t *testing.T) {
	result, _, err := client.Webhooks.List(context.Background())

	if err != nil {
		t.Errorf("Could not get webhooks list: %v", err)
	}

	if result.Success != true {
		t.Error("Got invalid result")
	}
}
