package pipedrive

import (
	"context"
	"fmt"
	"net/http"
)

// ActivityTypesService handles activities related
// methods of the Pipedrive API.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/ActivityTypes
type ActivityTypesService service

// ActivityType represents a Pipedrive activity type.
type ActivityType struct {
	ID           int         `json:"id"`
	OrderNr      int         `json:"order_nr"`
	Name         string      `json:"name"`
	KeyString    string      `json:"key_string"`
	IconKey      string      `json:"icon_key"`
	ActiveFlag   bool        `json:"active_flag"`
	Color        interface{} `json:"color"`
	IsCustomFlag bool        `json:"is_custom_flag"`
	AddTime      string      `json:"add_time"`
	UpdateTime   interface{} `json:"update_time"`
}

func (at ActivityType) String() string {
	return Stringify(at)
}

// ActivityTypesResponse represents multiple activity types response.
type ActivityTypesResponse struct {
	Success        bool           `json:"success"`
	Data           []ActivityType `json:"data"`
	AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

// ActivityTypeResponse represents single activity type response.
type ActivityTypeResponse struct {
	Success bool         `json:"success"`
	Data    ActivityType `json:"data"`
}

// List all activity types.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/ActivityTypes/get_activityTypes
func (s *ActivityTypesService) List(ctx context.Context) (*ActivityTypesResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/activityTypes", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *ActivityTypesResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// ActivityTypesAddOptions specifices the optional parameters to the
// ActivityTypesService.Create method.
type ActivityTypesAddOptions struct {
	Name    string `url:"name"`
	IconKey string `url:"icon_key"`
	Color   string `url:"color,omitempty"`
}

// Create a new activity type.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/ActivityTypes/post_activityTypes
func (s *ActivityTypesService) Create(ctx context.Context, opt *ActivityTypesAddOptions) (*ActivityTypeResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/activityTypes", nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *ActivityTypeResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// ActivityTypesEditOptions specifices the optional parameters to the
// ActivityTypesService.Update method.
type ActivityTypesEditOptions struct {
	Name    string `url:"name,omitempty"`
	IconKey string `url:"icon_key,omitempty"`
	Color   string `url:"color,omitempty"`
	OrderNr uint   `url:"order_nr,omitempty"`
}

// Update activity type.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/ActivityTypes/put_activityTypes_id
func (s *ActivityTypesService) Update(ctx context.Context, id int, opt *ActivityTypesEditOptions) (*ActivityTypeResponse, *Response, error) {
	uri := fmt.Sprintf("/activityTpes/%v", id)
	req, err := s.client.NewRequest(http.MethodPut, uri, nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *ActivityTypeResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// DeleteMultiple deletes activity types in bulk.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/ActivityTypes/delete_activityTypes
func (s *ActivityTypesService) DeleteMultiple(ctx context.Context, ids []int) (*Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, "/activityTypes", &DeleteMultipleOptions{
		Ids: arrayToString(ids, ","),
	}, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Delete an activity type.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/ActivityTypes/delete_activityTypes_id
func (s *ActivityTypesService) Delete(ctx context.Context, id int) (*Response, error) {
	uri := fmt.Sprintf("/activityTypes/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
