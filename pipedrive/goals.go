package pipedrive

import (
	"fmt"
	"net/http"
)

type GoalsService service

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

type SingleGoal struct {
	Success bool `json:"success"`
	Data    Goal `json:"data"`
}

type Goals struct {
	Success        bool           `json:"success"`
	Data           []Goal         `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

type GoalsListOptions struct {
	UserId   uint  `url:"user_id"`
	Everyone uint8 `url:"everyone"`
}

type GoalGetResultsByIdOptions struct {
	PeriodStart string `url:"period_start"`
	PeriodEnd   uint8  `url:"period_end"`
}

type GoalCreateOptions struct {
	GoalType     string `url:"goal_type"`
	ExpectedType string `url:"expected_type"`
	UserId       uint   `url:"user_io"`
	StageId      uint   `url:"expected_type"`
	Period       string `url:"period"`
	Expected     uint   `url:"expected"`
	Currency     string `url:"currency"`
	PipelineId   uint   `url:"pipeline_id"`
}

// Returns data about all goals.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Goals/get_goals
func (s *GoalsService) List(opt *GoalsListOptions) (*Goals, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/goals", opt, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *Goals

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Returns data about a specific goal.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Goals/get_goals_id
func (s *GoalsService) GetById(id int) (*SingleGoal, *Response, error) {
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

// Adds a new goal, returns the ID upon success.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Goals/post_goals
func (s *GoalsService) Create(opt *GoalCreateOptions) (*SingleGoal, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/goals", opt, nil)

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

// Updates the properties of a goal.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Goals/put_goals_id
func (s *GoalsService) Update(id int, opt *GoalCreateOptions) (*SingleGoal, *Response, error) {
	uri := fmt.Sprintf("/goals/%v", id)
	req, err := s.client.NewRequest(http.MethodPost, uri, opt, nil)

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

// Lists results of a specific goal.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Goals/get_goals_id_results
func (s *GoalsService) GetResultsById(id int, opt *GoalGetResultsByIdOptions) (*Goals, *Response, error) {
	uri := fmt.Sprintf("/goals/%v/results", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, opt, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *Goals

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Marks goal as deleted.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Goals/delete_goals_id
func (s *GoalsService) Delete(id int) (*Response, error) {
	uri := fmt.Sprintf("/goals/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
