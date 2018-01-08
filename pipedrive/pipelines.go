package pipedrive

import (
	"fmt"
	"net/http"
)

type PipelinesService service

type Pipelines struct {
	Success        bool           `json:"success"`
	Data           []Pipeline     `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

type Pipeline struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	URLTitle        string `json:"url_title"`
	OrderNr         int    `json:"order_nr"`
	Active          bool   `json:"active"`
	DealProbability bool   `json:"deal_probability"`
	AddTime         string `json:"add_time"`
	UpdateTime      string `json:"update_time"`
	Selected        bool   `json:"selected"`
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Pipelines/get_pipelines
func (s *PipelinesService) List() (*Pipelines, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/pipelines", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *Pipelines

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

func (s *PipelinesService) GetById(id int) (*SingleGoal, *Response, error) {
	uri := fmt.Sprintf("/goals/%v", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *SingleGoal

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Pipelines/delete_pipelines_id
func (s *PipelinesService) Delete(id int) (*Response, error) {
	uri := fmt.Sprintf("/pipelines/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
