package pipedrive

import (
	"context"
	"fmt"
	"net/http"
)

// GoalsService handles goals related
// methods of the Pipedrive API.
//
// Pipedrive API dcos: https://developers.pipedrive.com/docs/api/v1/#!/Goals
type GoalsService service

// Goal represents a Pipedrive goal.
type Goal struct {
	ID              int         `json:"id"`
	CompanyID       int         `json:"company_id"`
	UserID          int         `json:"user_id"`
	StageID         interface{} `json:"stage_id"`
	ActiveGoalID    int         `json:"active_goal_id"`
	Period          string      `json:"period"`
	Expected        int         `json:"expected"`
	ActiveFlag      bool        `json:"active_flag"`
	AddTime         string      `json:"add_time"`
	GoalType        string      `json:"goal_type"`
	ExpectedSum     int         `json:"expected_sum"`
	Currency        string      `json:"currency"`
	ExpectedType    string      `json:"expected_type"`
	CreatedByUserID int         `json:"created_by_user_id"`
	PipelineID      interface{} `json:"pipeline_id"`
	MasterExpected  int         `json:"master_expected"`
	Delivered       int         `json:"delivered"`
	DeliveredSum    int         `json:"delivered_sum"`
	PeriodStart     string      `json:"period_start"`
	PeriodEnd       string      `json:"period_end"`
	UserName        string      `json:"user_name"`
	Percentage      int         `json:"percentage,omitempty"`
}

func (g Goal) String() string {
	return Stringify(g)
}

// GoalResponse represents single goal response.
type GoalResponse struct {
	Success bool `json:"success"`
	Data    Goal `json:"data"`
}

// GoalsResponse represents multiple goals response.
type GoalsResponse struct {
	Success        bool           `json:"success"`
	Data           []Goal         `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

// GoalsListOptions specifices the optional parameters to the
// GoalsService.List method.
type GoalsListOptions struct {
	UserID   uint  `url:"user_id,omitempty"`
	Everyone uint8 `url:"everyone,omitempty"`
}

// List all goals.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Goals/get_goals
func (s *GoalsService) List(ctx context.Context, opt *GoalsListOptions) (*GoalsResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/goals", opt, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *GoalsResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// GetByID returns data about a specific goal.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Goals/get_goals_id
func (s *GoalsService) GetByID(ctx context.Context, id int) (*GoalResponse, *Response, error) {
	uri := fmt.Sprintf("/goals/%v", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *GoalResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// GoalCreateOptions specifices the optional parameters to the
// GoalsService.Create method.
type GoalCreateOptions struct {
	GoalType     string `url:"goal_type"`
	ExpectedType string `url:"expected_type"`
	UserID       uint   `url:"user_io"`
	StageID      uint   `url:"expected_type"`
	Period       string `url:"period"`
	Expected     uint   `url:"expected"`
	Currency     string `url:"currency"`
	PipelineID   uint   `url:"pipeline_id"`
}

// Create a new goal, returns the ID upon success.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Goals/post_goals
func (s *GoalsService) Create(ctx context.Context, opt *GoalCreateOptions) (*GoalResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/goals", opt, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *GoalResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Update the properties of a goal.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Goals/put_goals_id
func (s *GoalsService) Update(ctx context.Context, id int, opt *GoalCreateOptions) (*GoalResponse, *Response, error) {
	uri := fmt.Sprintf("/goals/%v", id)
	req, err := s.client.NewRequest(http.MethodPost, uri, opt, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *GoalResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// GoalGetResultsByIDOptions specifices the optional parameters to the
// GoalGetResultsByIdOptions.GetResultsByID method.
type GoalGetResultsByIDOptions struct {
	PeriodStart string `url:"period_start"`
	PeriodEnd   uint8  `url:"period_end"`
}

// GetResultsByID lists results of a specific goal.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Goals/get_goals_id_results
func (s *GoalsService) GetResultsByID(ctx context.Context, id int, opt *GoalGetResultsByIDOptions) (*GoalsResponse, *Response, error) {
	uri := fmt.Sprintf("/goals/%v/results", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, opt, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *GoalsResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Delete marks goal as deleted.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Goals/delete_goals_id
func (s *GoalsService) Delete(ctx context.Context, id int) (*Response, error) {
	uri := fmt.Sprintf("/goals/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
