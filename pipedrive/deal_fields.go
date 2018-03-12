package pipedrive

import (
	"context"
	"fmt"
	"net/http"
)

type DealFieldsService service

type DealField struct {
	ID                 int         `json:"id"`
	Key                string      `json:"key"`
	Name               string      `json:"name"`
	OrderNr            int         `json:"order_nr,omitempty"`
	PicklistData       interface{} `json:"picklist_data,omitempty"`
	FieldType          string      `json:"field_type"`
	AddTime            string      `json:"add_time,omitempty"`
	UpdateTime         string      `json:"update_time,omitempty"`
	ActiveFlag         bool        `json:"active_flag"`
	EditFlag           bool        `json:"edit_flag"`
	IndexVisibleFlag   bool        `json:"index_visible_flag,omitempty"`
	DetailsVisibleFlag bool        `json:"details_visible_flag,omitempty"`
	AddVisibleFlag     bool        `json:"add_visible_flag,omitempty"`
	ImportantFlag      bool        `json:"important_flag,omitempty"`
	BulkEditAllowed    bool        `json:"bulk_edit_allowed,omitempty"`
	SearchableFlag     bool        `json:"searchable_flag,omitempty"`
	FilteringAllowed   bool        `json:"filtering_allowed,omitempty"`
	SortableFlag       bool        `json:"sortable_flag,omitempty"`
	UseField           string      `json:"use_field,omitempty"`
	Link               string      `json:"link,omitempty"`
	MandatoryFlag      bool        `json:"mandatory_flag"`
	IsSubfield         bool        `json:"is_subfield,omitempty"`
	Options            []struct {
		ID    string `json:"id"`
		Label string `json:"label"`
	} `json:"options,omitempty"`
	BulkEditAllowedConditions struct {
		Status string `json:"status"`
	} `json:"bulk_edit_allowed_conditions,omitempty"`
}

type DealFieldsResponse struct {
	Success        bool           `json:"success"`
	Data           []DealField    `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

type DealFieldResponse struct {
	Success        bool           `json:"success"`
	Data           DealField      `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

type DealFieldCreateOptions struct {
	Name      string    `url:"name"`
	FieldType FieldType `url:"field_type"`
	Options   string    `url:"options"`
}

type DealFieldUpdateOptions struct {
	Name    string `url:"name"`
	Options string `url:"options"`
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/DealFields/get_dealFields
func (s *DealFieldsService) List(ctx context.Context) (*DealFieldsResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/dealFields", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *DealFieldsResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/DealFields/get_dealFields_id
func (s *DealFieldsService) GetById(ctx context.Context, id int) (*DealFieldResponse, *Response, error) {
	uri := fmt.Sprintf("/dealFields/%v", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *DealFieldResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/DealFields/post_dealFields
func (s *DealFieldsService) Create(ctx context.Context, opt *DealFieldCreateOptions) (*DealFieldResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/dealFields", nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *DealFieldResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/DealFields/put_dealFields_id
func (s *DealFieldsService) Update(ctx context.Context, id int, opt *DealFieldUpdateOptions) (*ProductFieldResponse, *Response, error) {
	uri := fmt.Sprintf("/dealFields/%v", id)
	req, err := s.client.NewRequest(http.MethodPut, uri, nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *ProductFieldResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/DealFields/delete_dealFields
func (s *DealFieldsService) DeleteMultiple(ctx context.Context, ids []int) (*Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, "/dealFields", &DeleteMultipleOptions{
		Ids: arrayToString(ids, ","),
	}, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/DealFields/delete_dealFields_id
func (s *DealFieldsService) Delete(ctx context.Context, id uint) (*Response, error) {
	uri := fmt.Sprintf("/dealFields/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
