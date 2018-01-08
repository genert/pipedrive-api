package pipedrive

import "net/http"

type CurrenciesService service

type Currency struct {
	ID            int    `json:"id,omitempty"`
	Code          string `json:"code,omitempty"`
	Name          string `json:"name,omitempty"`
	DecimalPoints int    `json:"decimal_points,omitempty"`
	Symbol        string `json:"symbol"`
	ActiveFlag    bool   `json:"active_flag,omitempty"`
	IsCustomFlag  bool   `json:"is_custom_flag,omitempty"`
}

type Currencies struct {
	Success   bool       `json:"success,omitempty"`
	Data      []Currency `json:"data,omitempty"`
	Error     string     `json:"error"`
	ErrorInfo string     `json:"error_info"`
}

type CurrenciesListOptions struct {
	Term string `url:"term,omitempty"`
}

// Returns all supported currencies in given account which should be used
// when saving monetary values with other objects.
// The 'code' parameter of the returning objects is
// the currency code according to ISO 4217 for all non-custom currencies.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Currencies/get_currencies
func (s *CurrenciesService) List(opt *CurrenciesListOptions) (*Currencies, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/currencies", opt, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *Currencies

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}
