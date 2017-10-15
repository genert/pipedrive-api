package pipedrive

import "testing"

func TestSearchResultsService_List(t *testing.T) {
	opt := &SearchResultsListOptions{
		Term: "test",
	}

	searchResults, _, err := client.SearchResults.List(opt)

	if err != nil {
		t.Errorf("Could not get search results: %v", err)
	}

	if searchResults.Success != true {
		t.Error("Got invalid search results")
	}
}
