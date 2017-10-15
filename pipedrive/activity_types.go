package pipedrive

import (
	"fmt"
	"net/http"
)

type ActivityTypesService service

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

type ActivityTypes struct {
	Success bool           `json:"success"`
	Data    []ActivityType `json:"data"`
}

type SingleActivityType struct {
	Success bool         `json:"success"`
	Data    ActivityType `json:"data"`
}

type ActivityTypesAddOptions struct {
	Name    string `url:"name"`
	IconKey string `url:"icon_key"`
	Color   string `url:"color,omitempty"`
}

type ActivityTypesEditOptions struct {
	Id      string `url:"id"`
	Name    string `url:"name,omitempty"`
	IconKey string `url:"icon_key,omitempty"`
	Color   string `url:"color,omitempty"`
	OrderNr uint   `url:"order_nr,omitempty"`
}

// Returns all activity types.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/ActivityTypes/get_activityTypes
func (s *ActivityTypesService) List() (*ActivityTypes, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/activityTypes", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *ActivityTypes

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Adds a new activity type.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/ActivityTypes/post_activityTypes
func (s *ActivityTypesService) Add(opt *ActivityTypesAddOptions) (*SingleActivityType, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/activityTypes", opt, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *SingleActivityType

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Updates an activity type.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/ActivityTypes/put_activityTypes_id
func (s *ActivityTypesService) Edit(opt *ActivityTypesEditOptions) (*SingleActivityType, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPut, "/activityTypes", opt, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *SingleActivityType

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Marks multiple activity types as deleted.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/ActivityTypes/delete_activityTypes
func (s *ActivityTypesService) DeleteMultiple(ids []int) (*Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, "/activityTypes", &DeleteMultipleOptions{
		Ids: arrayToString(ids, ","),
	}, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// Marks an activity type as deleted.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/ActivityTypes/delete_activityTypes_id
func (s *ActivityTypesService) Delete(id int) (*Response, error) {
	uri := fmt.Sprintf("/activityTypes/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
