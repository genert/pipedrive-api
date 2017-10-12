package pipedrive

import (
	"log"
	"net/http"
)

type DealService service

type Deal struct {
	ID       int 	  `json:"id,omitempty"`
	StageId  int 	  `json:"stage_id,omitempty"`
	Title    string   `json:"title,omitempty"`
}

type Deals struct {
	Success  bool  `json:"success,omitempty"`
	Data []  Deal  `json:"data,omitempty"`
}

func (s *DealService) List() (*Deals, *http.Response, error) {
	uri := s.client.CreateRequestUrl("deals")

	// Build the request
	req, err := s.client.NewRequest("GET", uri, nil)

	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil, nil, err
	}

	var record *Deals

	resp, err := s.client.Do(req, &record)

	if err != nil {
		log.Fatal("Do: ", err)
		return nil, resp, err
	}

	return record, resp, nil
}