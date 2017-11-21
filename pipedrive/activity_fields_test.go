package pipedrive

import "testing"

func TestActivityFieldsService_List(t *testing.T) {
	if *apiIntegration {
		searchResults, _, err := client.ActivityFields.List()

		if err != nil {
			t.Errorf("Could not get search results: %v", err)
		}

		if searchResults.Success != true {
			t.Error("Got invalid activity fields")
		}
	}
}
