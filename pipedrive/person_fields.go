package pipedrive

import (
	"context"
	"fmt"
	"net/http"
)

// PersonFieldsService handles person fields related
// methods of the Pipedrive API.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/PersonFields
type PersonFieldsService service

// PersonField represents a Pipedrive person field.
type PersonField struct {
	ID                 int         `json:"id"`
	Key                string      `json:"key"`
	Name               string      `json:"name"`
	OrderNr            int         `json:"order_nr"`
	PicklistData       interface{} `json:"picklist_data,omitempty"`
	FieldType          string      `json:"field_type"`
	AddTime            string      `json:"add_time"`
	UpdateTime         string      `json:"update_time"`
	ActiveFlag         bool        `json:"active_flag"`
	EditFlag           bool        `json:"edit_flag"`
	IndexVisibleFlag   bool        `json:"index_visible_flag"`
	DetailsVisibleFlag bool        `json:"details_visible_flag"`
	AddVisibleFlag     bool        `json:"add_visible_flag"`
	ImportantFlag      bool        `json:"important_flag"`
	BulkEditAllowed    bool        `json:"bulk_edit_allowed"`
	SearchableFlag     bool        `json:"searchable_flag"`
	FilteringAllowed   bool        `json:"filtering_allowed"`
	SortableFlag       bool        `json:"sortable_flag"`
	UseField           string      `json:"use_field,omitempty"`
	Link               string      `json:"link,omitempty"`
	MandatoryFlag      bool        `json:"mandatory_flag"`
	DisplayField       string      `json:"display_field,omitempty"`
	Autocomplete       string      `json:"autocomplete,omitempty"`
	Options            []struct {
		ID    int    `json:"id"`
		Label string `json:"label"`
	} `json:"options,omitempty"`
}

// PersonFieldsResponse represents multiple person fields response.
type PersonFieldsResponse struct {
	Success        bool           `json:"success"`
	Data           []PersonField  `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

// PersonFieldResponse represents single person field response.
type PersonFieldResponse struct {
	Success        bool           `json:"success"`
	Data           PersonField    `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

// List all person fields.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/PersonFields/get_personFields
func (s *PersonFieldsService) List(ctx context.Context) (*PersonFieldsResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/personFields", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *PersonFieldsResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// GetByID returns a specific person field.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/ProductFields/get_productFields_id
func (s *PersonFieldsService) GetByID(ctx context.Context, id int) (*PersonFieldResponse, *Response, error) {
	uri := fmt.Sprintf("/personFields/%v", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *PersonFieldResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// PersonFieldCreateOptions specifices the optional parameters to the
// PersonFieldsService.Create method.
type PersonFieldCreateOptions struct {
	Name      string    `url:"name"`
	FieldType FieldType `url:"field_type"`
	Options   string    `url:"options"`
}

// Create a person field.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/ProductFields/post_productFields
func (s *PersonFieldsService) Create(ctx context.Context, opt *PersonFieldCreateOptions) (*ProductFieldResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/personFields", nil, opt)

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

// PersonFieldUpdateOptions specifices the optional parameters to the
// PersonFieldsService.Update method.
type PersonFieldUpdateOptions struct {
	Name    string `url:"name"`
	Options string `url:"options"`
}

// Update a person field.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/PersonFields/put_personFields_id
func (s *PersonFieldsService) Update(ctx context.Context, id int, opt *PersonFieldUpdateOptions) (*PersonFieldResponse, *Response, error) {
	uri := fmt.Sprintf("/personFields/%v", id)
	req, err := s.client.NewRequest(http.MethodPut, uri, nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *PersonFieldResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// DeleteMultiple marks multiple person fields as deleted.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/PersonFields/delete_personFields
func (s *PersonFieldsService) DeleteMultiple(ctx context.Context, ids []int) (*Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, "/personFields", &DeleteMultipleOptions{
		Ids: arrayToString(ids, ","),
	}, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Delete marks person field as deleted.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/PersonFields/delete_personFields_id
func (s *PersonFieldsService) Delete(ctx context.Context, id int) (*Response, error) {
	uri := fmt.Sprintf("/personFields/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
