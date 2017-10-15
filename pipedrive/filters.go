package pipedrive

import "fmt"

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

type SingleFilter struct {
	Success bool `json:"success"`
	Data    struct {
		Filter
		Conditions FilterConditions `json:"conditions"`
	} `json:"data"`
}

type Filters struct {
	Success bool     `json:"success"`
	Data    []Filter `json:"data"`
}

type FiltersListOptions struct {
	Type string `url:"type,omitempty"`
}

type FiltersDeleteMultipleOptions struct {
	Ids string `url:"ids,omitempty"`
}

// Returns data about all filters
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Filters/get_filters
func (s *FiltersService) List(opt *FiltersListOptions) (*Filters, *Response, error) {
	req, err := s.client.NewRequest("GET", "/filters", opt, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *Filters

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Returns data about a specific filter. Note that this also returns the condition lines of the filter.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Filters/get_filters_id
func (s *FiltersService) GetById(id int) (*SingleFilter, *Response, error) {
	uri := fmt.Sprintf("/filters/%v", id)
	req, err := s.client.NewRequest("GET", uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *SingleFilter

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Marks multiple filters as deleted.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Filters/delete_filters
func (s *FiltersService) DeleteMultiple(ids []int) (*Response, error) {
	req, err := s.client.NewRequest("DELETE", "/filter", &FiltersDeleteMultipleOptions{
		Ids: arrayToString(ids, ","),
	}, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// Marks a filter as deleted.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Filters/delete_filters_id
func (s *FiltersService) Delete(id int) (*Response, error) {
	uri := fmt.Sprintf("/filters/%v", id)
	req, err := s.client.NewRequest("DELETE", uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
