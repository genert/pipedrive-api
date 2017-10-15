package pipedrive

import (
	"fmt"
	"net/http"
)

type DealService service

type Deal struct {
	ID         int    `json:"id,omitempty"`
	StageId    int    `json:"stage_id,omitempty"`
	Title      string `json:"title,omitempty"`
	Value      int    `json:"value,omitempty"`
	Currency   string `json:"currency,omitempty"`
	AddTime    string `json:"add_time,omitempty"`
	UpdateTime string `json:"update:time,omitempty"`
}

type Deals struct {
	Success bool   `json:"success,omitempty"`
	Data    []Deal `json:"data,omitempty"`
}

type DealUpdate struct {
	Success bool `json:"success,omitempty"`
	Data    Deal `json:"data,omitempty"`
}

// IssueRequest represents a request to create/edit an issue.
// It is separate from Issue above because otherwise Labels
// and Assignee fail to serialize to the correct JSON.
type DealRequest struct {
	Title     *string   `json:"title,omitempty"`
	Body      *string   `json:"body,omitempty"`
	Labels    *[]string `json:"labels,omitempty"`
	Assignee  *string   `json:"assignee,omitempty"`
	State     *string   `json:"state,omitempty"`
	Milestone *int      `json:"milestone,omitempty"`
	Assignees *[]string `json:"assignees,omitempty"`
}

// List updates about a deal
func (s *DealService) ListDealUpdates(id int) (*Deals, *Response, error) {
	uri := fmt.Sprintf("/deals/%v/flow", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *Deals

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

func (s *DealService) List() (*Deals, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/deals", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *Deals

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

func (s *DealService) Duplicate(id int) (*DealUpdate, *Response, error) {
	uri := fmt.Sprintf("/deals/%v/duplicate", id)
	req, err := s.client.NewRequest(http.MethodPost, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *DealUpdate

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}
