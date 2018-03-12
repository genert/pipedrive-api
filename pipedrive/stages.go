package pipedrive

import (
	"context"
	"fmt"
	"net/http"
)

// StagesService handles stages related
// methods of the Pipedrive API.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Stages
type StagesService service

// Stage represents a Pipedrive stage.
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

func (s Stage) String() string {
	return Stringify(s)
}

// StagesResponse represents multiple stages response.
type StagesResponse struct {
	Success bool    `json:"success"`
	Data    []Stage `json:"data"`
}

// StageResponse represents single stage response.
type StageResponse struct {
	Success bool  `json:"success"`
	Data    Stage `json:"data"`
}

// StageDealsResponse represents stage deals response.
type StageDealsResponse struct {
	Success        bool           `json:"success"`
	Data           []Deal         `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

// StagesListOptions specifices the optional parameters to the
// StagesService.List method.
type StagesListOptions struct {
	PipelineID uint `url:"pipeline_id"`
}

// List returns data about all stages.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Stages
func (s *StagesService) List(ctx context.Context, opt *StagesListOptions) (*StagesResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/stages", nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *StagesResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// GetByID returns data about a specific stage.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Stages/get_stages_id
func (s *StagesService) GetByID(ctx context.Context, id int) (*StageResponse, *Response, error) {
	uri := fmt.Sprintf("/stages/%v", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *StageResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// StagesGetDealsInStageOptions specifices the optional parameters to the
// StagesService.GetDealsInStage method.
type StagesGetDealsInStageOptions struct {
	FilterID uint  `url:"filter_id"`
	UserID   uint  `url:"user_id"`
	Everyone uint8 `url:"everyone"`
	Start    uint  `url:"start"`
	Limit    uint  `url:"limit"`
}

// GetDealsInStage lists deals in a specific stage.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Stages/get_stages_id_deals
func (s *StagesService) GetDealsInStage(ctx context.Context, id int, opt *StagesGetDealsInStageOptions) (*StageDealsResponse, *Response, error) {
	uri := fmt.Sprintf("/stages/%v/deals", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *StageDealsResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// StagesCreateOptions specifices the optional parameters to the
// StagesService.Create method.
type StagesCreateOptions struct {
	Name            string `url:"name"`
	PipelineID      uint   `url:"pipeline_id"`
	DealProbability uint   `url:"deal_probability"`
	RottenFlag      uint8  `url:"rotten_flag"`
	RottenDays      uint   `url:"rotten_days"`
}

// Create a new stage, returns the ID upon success.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Stages/post_stages
func (s *StagesService) Create(ctx context.Context, opt *StagesCreateOptions) (*StageResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/stages", nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *StageResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// StagesUpdateOptions specifices the optional parameters to the
// StagesService.Update method.
type StagesUpdateOptions struct {
	Name            string `url:"name"`
	PipelineID      uint   `url:"pipeline_id"`
	OrderNr         uint   `url:"order_nr"`
	DealProbability uint   `url:"deal_probability"`
	RottenFlag      uint8  `url:"rotten_flag"`
	RottenDays      uint   `url:"rotten_days"`
}

// Update the properties of a stage.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Stages/put_stages_id
func (s *StagesService) Update(ctx context.Context, id int, opt *StagesUpdateOptions) (*StageResponse, *Response, error) {
	uri := fmt.Sprintf("/stages/%v", id)
	req, err := s.client.NewRequest(http.MethodPut, uri, nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *StageResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// DeleteMultiple marks multiple stages as deleted.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Stages/put_stages_id
func (s *StagesService) DeleteMultiple(ctx context.Context, ids []int) (*Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, "/stages", &DeleteMultipleOptions{
		Ids: arrayToString(ids, ","),
	}, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Delete marks a stage as deleted.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Stages/delete_stages_id
func (s *StagesService) Delete(ctx context.Context, id int) (*Response, error) {
	uri := fmt.Sprintf("/stages/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
