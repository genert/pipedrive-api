package integration

import (
	"context"
	"testing"
)

func TestActivityFieldsService_List(t *testing.T) {
	result, _, err := client.ActivityFields.List(context.Background())

	if err != nil {
		t.Errorf("Could not get results: %v", err)
	}

	if result.Success != true {
		t.Error("Got invalid result")
	}
}
