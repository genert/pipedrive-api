package pipedrive

import (
	"net/http"
)

type CurrenciesService service

type Currency struct {
	ID       		int 	`json:"id,omitempty"`
	Code  			string 	`json:"code,omitempty"`
	Name    		string  `json:"name,omitempty"`
	DecimalPoints 	int 	`json:"decimal_points,omitempty"`
	Symbol 			string 	`json:"symbol"`
	ActiveFlag 		bool 	`json:"active_flag,omitempty"`
	IsCustomFlag 	bool 	`json:"is_custom_flag,omitempty"`
}

type Currencies struct {
	Success  bool  		`json:"success,omitempty"`
	Data []  Currency  	`json:"data,omitempty"`
}

func (s *CurrenciesService) List() (*Currencies, *http.Response, error) {
	uri := s.client.CreateRequestUrl("/currencies")
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