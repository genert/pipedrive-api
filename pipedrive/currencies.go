package pipedrive

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
	Success bool       `json:"success,omitempty"`
	Data    []Currency `json:"data,omitempty"`
}

// Returns all supported currencies in given account which should be used
// when saving monetary values with other objects.
// The 'code' parameter of the returning objects is
// the currency code according to ISO 4217 for all non-custom currencies.
// https://developers.pipedrive.com/docs/api/v1/#!/Currencies/get_currencies
func (s *CurrenciesService) List() (*Currencies, *Response, error) {
	uri, err := s.client.CreateRequestUrl("/currencies", nil)

	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", uri, nil)

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
