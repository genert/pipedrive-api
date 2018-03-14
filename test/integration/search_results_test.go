package integration

import (
	"context"
	"testing"

	"github.com/genert/pipedrive-api/pipedrive"
)

func TestSearchResults_Search(t *testing.T) {
	result, _, err := client.SearchResults.Search(context.Background(), &pipedrive.SearchResultsListOptions{
		Term: "test",
	})

	if err != nil {
		t.Errorf("Could not get search results: %v", err)
	}

	if result.Success != true {
		t.Error("Got invalid search results")
	}
}
