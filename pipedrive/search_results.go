package pipedrive

import (
	"context"
	"net/http"
)

type SearchResultsService service

type SearchResult struct {
	Type        string  `json:"type"`
	ID          int     `json:"id"`
	Source      string  `json:"source"`
	ResultScore float64 `json:"result_score"`
	Notes       struct {
		Count   int           `json:"count"`
		Content []interface{} `json:"content"`
	} `json:"notes"`
	Fields struct {
		Count int           `json:"count"`
		Names []interface{} `json:"names"`
	} `json:"fields"`
	Title   string `json:"title"`
	Details struct {
		Phone      interface{} `json:"phone"`
		Email      interface{} `json:"email"`
		OrgID      interface{} `json:"org_id"`
		OrgName    interface{} `json:"org_name"`
		OrgAddress string      `json:"org_address"`
		Picture    interface{} `json:"picture"`
	} `json:"details"`
}

type SearchResults struct {
	Success        bool           `json:"success"`
	Data           []SearchResult `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

type SearchResultsListOptions struct {
	Term       string `url:"term,omitempty"`
	ItemType   string `url:"item_type,omitempty"`
	Start      uint   `url:"start,omitempty"`
	Limit      uint   `url:"limit,omitempty"`
	ExactMatch uint8  `url:"exact_match,omitempty"`
}

// Performs a search across the account and returns SearchResults.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/SearchResults/get_searchResults
func (s *SearchResultsService) List(ctx context.Context, opt *SearchResultsListOptions) (*Recents, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/searchResults", opt, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *Recents

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}
