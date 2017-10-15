package pipedrive

import "testing"

func TestActivityFieldsService_List(t *testing.T) {
	searchResults, _, err := client.ActivityFields.List()

	if err != nil {
		t.Error("Could not get search results: %v", err)
	}

	if searchResults.Success != true {
		t.Error("Got invalid activity fields")
	}
}
