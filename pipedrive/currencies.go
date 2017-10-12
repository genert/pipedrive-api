package pipedrive

import (
	"net/http"
	"log"
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
	uri := s.client.CreateRequestUrl("currencies")

	// Build the request
	req, err := s.client.NewRequest("GET", uri, nil)

	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil, nil, err
	}

	var record *Currencies

	resp, err := s.client.Do(req, &record)

	if err != nil {
		log.Fatal("Do: ", err)
		return nil, resp, err
	}

	return record, resp, nil
}