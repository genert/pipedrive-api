package integration

import (
	"context"
	"testing"
)

func TestUserConnectionsService_List(t *testing.T) {
	result, _, err := client.UserConnections.List(context.Background())

	if err != nil {
		t.Errorf("Could not get user connections list: %v", err)
	}

	if result.Success != true {
		t.Error("Got invalid result")
	}
}
