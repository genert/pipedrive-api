package pipedrive

import (
	"fmt"
	"net/http"
)

type StagesService service

type Stage struct {
	ID              int         `json:"id"`
	OrderNr         int         `json:"order_nr"`
	Name            string      `json:"name"`
	ActiveFlag      bool        `json:"active_flag"`
	DealProbability int         `json:"deal_probability"`
	PipelineID      int         `json:"pipeline_id"`
	RottenFlag      bool        `json:"rotten_flag"`
	RottenDays      interface{} `json:"rotten_days"`
	AddTime         string      `json:"add_time"`
	UpdateTime      string      `json:"update_time"`
	PipelineName    string      `json:"pipeline_name"`
}

type Stages struct {
	Success bool    `json:"success"`
	Data    []Stage `json:"data"`
}

type SingleStage struct {
	Success bool  `json:"success"`
	Data    Stage `json:"data"`
}

type StagesListOptions struct {
	PipelineId uint `url:"pipeline_id"`
}

type StagesCreateOptions struct {
	Name            string `url:"name"`
	PipelineId      uint   `url:"pipeline_id"`
	DealProbability uint   `url:"deal_probability"`
	RottenFlag      uint8  `url:"rotten_flag"`
	RottenDays      uint   `url:"rotten_days"`
}

type StagesUpdateOptions struct {
	Name            string `url:"name"`
	PipelineId      uint   `url:"pipeline_id"`
	OrderNr         uint   `url:"order_nr"`
	DealProbability uint   `url:"deal_probability"`
	RottenFlag      uint8  `url:"rotten_flag"`
	RottenDays      uint   `url:"rotten_days"`
}

type StagesGetDealsInStageOptions struct {
	FilterId uint  `url:"filter_id"`
	UserId   uint  `url:"user_id"`
	Everyone uint8 `url:"everyone"`
	Start    uint  `url:"start"`
	Limit    uint  `url:"limit"`
}

type StagesDeals struct {
	Success        bool           `json:"success"`
	Data           []Deal         `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

// Returns data about all stages.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Stages
func (s *StagesService) List(opt *StagesListOptions) (*Stages, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/stages", nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *Stages

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Returns data about a specific stage.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Stages/get_stages_id
func (s *StagesService) GetById(id uint) (*SingleStage, *Response, error) {
	uri := fmt.Sprintf("/stages/%v", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *SingleStage

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Lists deals in a specific stage.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Stages/get_stages_id_deals
func (s *StagesService) GetDealsInStage(id uint, opt *StagesGetDealsInStageOptions) (*StagesDeals, *Response, error) {
	uri := fmt.Sprintf("/stages/%v/deals", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *StagesDeals

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Adds a new stage, returns the ID upon success.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Stages/post_stages
func (s *StagesService) Create(opt *StagesCreateOptions) (*SingleStage, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/stages", nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *SingleStage

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Updates the properties of a stage.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Stages/put_stages_id
func (s *StagesService) Update(id uint, opt *StagesGetDealsInStageOptions) (*SingleStage, *Response, error) {
	uri := fmt.Sprintf("/stages/%v", id)
	req, err := s.client.NewRequest(http.MethodPut, uri, nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *SingleStage

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Marks multiple stages as deleted.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Stages/put_stages_id
func (s *StagesService) DeleteMultiple(ids []int) (*Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, "/stages", &DeleteMultipleOptions{
		Ids: arrayToString(ids, ","),
	}, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// Marks a stage as deleted.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Stages/delete_stages_id
func (s *StagesService) Delete(id int) (*Response, error) {
	uri := fmt.Sprintf("/stages/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
