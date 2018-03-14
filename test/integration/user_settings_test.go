package integration

import (
	"context"
	"testing"
)

func TestUserSettings_List(t *testing.T) {
	result, _, err := client.UserSettings.List(context.Background())

	if err != nil {
		t.Errorf("Could not get webhooks list: %v", err)
	}

	if result.Success != true {
		t.Error("Got invalid result")
	}
}
