package pipedrive

import (
	"context"
	"fmt"
	"net/http"
)

type OrganizationFieldsService service

type OrganizationField struct {
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
	UseField           string      `json:"use_field,omitempty"`
	Link               string      `json:"link,omitempty"`
	MandatoryFlag      bool        `json:"mandatory_flag"`
	DisplayField       string      `json:"display_field,omitempty"`
	Options            []struct {
		ID    int    `json:"id"`
		Label string `json:"label"`
	} `json:"options,omitempty"`
	IsSubfield bool `json:"is_subfield,omitempty"`
}

type OrganizationFieldsResponse struct {
	Success        bool                `json:"success"`
	Data           []OrganizationField `json:"data"`
	AdditionalData AdditionalData      `json:"additional_data"`
}

type OrganizationFieldResponse struct {
	Success        bool              `json:"success"`
	Data           OrganizationField `json:"data"`
	AdditionalData AdditionalData    `json:"additional_data"`
}

type OrganizationFieldCreateOptions struct {
	Name      string    `url:"name"`
	FieldType FieldType `url:"field_type"`
	Options   string    `url:"options"`
}

type OrganizationFieldUpdateOptions struct {
	Name    string `url:"name"`
	Options string `url:"options"`
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/OrganizationFields/get_organizationFields
func (s *OrganizationFieldsService) List(ctx context.Context) (*OrganizationFieldsResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/organizationFields", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *OrganizationFieldsResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/OrganizationFields/get_organizationFields_id
func (s *OrganizationFieldsService) GetById(ctx context.Context, id int) (*OrganizationFieldResponse, *Response, error) {
	uri := fmt.Sprintf("/organizationFields/%v", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *OrganizationFieldResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/OrganizationFields/post_organizationFields
func (s *OrganizationFieldsService) Create(ctx context.Context, opt *OrganizationFieldCreateOptions) (*OrganizationFieldResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/organizationFields", nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *OrganizationFieldResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/OrganizationFields/put_organizationFields_id
func (s *OrganizationFieldsService) Update(ctx context.Context, id int, opt *OrganizationFieldUpdateOptions) (*OrganizationFieldResponse, *Response, error) {
	uri := fmt.Sprintf("/organizationFields/%v", id)
	req, err := s.client.NewRequest(http.MethodPut, uri, nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *OrganizationFieldResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/OrganizationFields/delete_organizationFields
func (s *OrganizationFieldsService) DeleteMultiple(ctx context.Context, ids []int) (*Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, "/organizationFields", &DeleteMultipleOptions{
		Ids: arrayToString(ids, ","),
	}, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/OrganizationFields/delete_organizationFields_id
func (s *OrganizationFieldsService) Delete(ctx context.Context, id int) (*Response, error) {
	uri := fmt.Sprintf("/organizationFields/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
