package pipedrive

import (
	"context"
	"fmt"
	"net/http"
)

type FiltersService service

type Filter struct {
	ID            int         `json:"id"`
	Name          string      `json:"name"`
	ActiveFlag    bool        `json:"active_flag"`
	Type          string      `json:"type"`
	TemporaryFlag interface{} `json:"temporary_flag"`
	UserID        int         `json:"user_id"`
	AddTime       string      `json:"add_time"`
	UpdateTime    string      `json:"update_time"`
	VisibleTo     string      `json:"visible_to"`
	CustomViewID  int         `json:"custom_view_id"`
}

type FilterConditions struct {
	Glue       string `json:"glue"`
	Conditions []struct {
		Glue       string `json:"glue"`
		Conditions []struct {
			Object     string      `json:"object"`
			FieldID    string      `json:"field_id"`
			Operator   string      `json:"operator"`
			Value      string      `json:"value"`
			ExtraValue interface{} `json:"extra_value"`
		} `json:"conditions"`
	} `json:"conditions"`
}

type FilterResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Filter
		Conditions FilterConditions `json:"conditions"`
	} `json:"data"`
}

type FiltersResponse struct {
	Success bool     `json:"success"`
	Data    []Filter `json:"data"`
}

type FiltersListOptions struct {
	Type string `url:"type,omitempty"`
}

type FilterUpdateOptions struct {
	Name       string `url:"name,omitempty"`
	Conditions string `url:"conditions"`
}

type FilterCreateOptions struct {
	Name       string `url:"name"`
	Conditions string `url:"conditions"`
	Type       string `url:"type"`
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Filters/get_filters
func (s *FiltersService) List(ctx context.Context, opt *FiltersListOptions) (*FiltersResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/filters", opt, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *FiltersResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Filters/get_filters_id
func (s *FiltersService) GetById(ctx context.Context, id int) (*FilterResponse, *Response, error) {
	uri := fmt.Sprintf("/filters/%v", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *FilterResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Filters/post_filters
func (s *FiltersService) Create(ctx context.Context, opt *FilterCreateOptions) (*FilterResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/filters", nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *FilterResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Filters/put_filters_id
func (s *FiltersService) Update(ctx context.Context, id int, opt *FilterUpdateOptions) (*FilterResponse, *Response, error) {
	uri := fmt.Sprintf("/filters/%v", id)
	req, err := s.client.NewRequest(http.MethodPut, uri, nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *FilterResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Filters/delete_filters
func (s *FiltersService) DeleteMultiple(ctx context.Context, ids []int) (*Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, "/filter", &DeleteMultipleOptions{
		Ids: arrayToString(ids, ","),
	}, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Filters/delete_filters_id
func (s *FiltersService) Delete(ctx context.Context, id int) (*Response, error) {
	uri := fmt.Sprintf("/filters/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
