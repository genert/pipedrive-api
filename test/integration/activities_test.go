package integration

import (
	"context"
	"testing"
)

func TestActivitiesService_List(t *testing.T) {
	result, _, err := client.Activities.List(context.Background())

	if err != nil {
		t.Errorf("Could not get result: %v", err)
	}

	if result.Success != true {
		t.Error("Got invalid result")
	}
}
