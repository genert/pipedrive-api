package pipedrive

import (
	"context"
	"net/http"
)

// CurrenciesService handles currencies related
// methods of the Pipedrive API.
//
// Pipedrive API dcos: https://developers.pipedrive.com/docs/api/v1/#!/Currencies
type CurrenciesService service

// Currency represents a Pipedrive currency.
type Currency struct {
	ID            int    `json:"id,omitempty"`
	Code          string `json:"code,omitempty"`
	Name          string `json:"name,omitempty"`
	DecimalPoints int    `json:"decimal_points,omitempty"`
	Symbol        string `json:"symbol"`
	ActiveFlag    bool   `json:"active_flag,omitempty"`
	IsCustomFlag  bool   `json:"is_custom_flag,omitempty"`
}

func (c Currency) String() string {
	return Stringify(c)
}

// CurrenciesResponse represents multiple currencies response.
type CurrenciesResponse struct {
	Success   bool       `json:"success,omitempty"`
	Data      []Currency `json:"data,omitempty"`
	Error     string     `json:"error"`
	ErrorInfo string     `json:"error_info"`
}

// CurrenciesListOptions specifices the optional parameters to the
// CurrenciesService.List method.
type CurrenciesListOptions struct {
	Term string `url:"term,omitempty"`
}

// List returns all supported currencies in given account which should be used
// when saving monetary values with other objects.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Currencies/get_currencies
func (s *CurrenciesService) List(ctx context.Context, opt *CurrenciesListOptions) (*CurrenciesResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/currencies", opt, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *CurrenciesResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}
