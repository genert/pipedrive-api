package integration

import "testing"

func TestUserConnectionsService_List(t *testing.T) {
	result, _, err := client.UserConnections.List()

	if err != nil {
		t.Errorf("Could not get webhooks list: %v", err)
	}

	if result.Success != true {
		t.Error("Got invalid result")
	}
}
