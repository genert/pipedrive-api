package pipedrive

import (
	"context"
	"fmt"
	"net/http"
)

// OrganizationFieldsService handles organization fields related
// methods of the Pipedrive API.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/OrganizationFields
type OrganizationFieldsService service

// OrganizationField represents a Pipedrive organization field.
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

func (of OrganizationField) String() string {
	return Stringify(of)
}

// OrganizationFieldsResponse represents multiple organization fields response.
type OrganizationFieldsResponse struct {
	Success        bool                `json:"success"`
	Data           []OrganizationField `json:"data"`
	AdditionalData AdditionalData      `json:"additional_data"`
}

// OrganizationFieldResponse represents single organization field response.
type OrganizationFieldResponse struct {
	Success        bool              `json:"success"`
	Data           OrganizationField `json:"data"`
	AdditionalData AdditionalData    `json:"additional_data"`
}

// List all organization fields within company.
//
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

// GetByID returns a specific organization field.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/OrganizationFields/get_organizationFields_id
func (s *OrganizationFieldsService) GetByID(ctx context.Context, id int) (*OrganizationFieldResponse, *Response, error) {
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

// OrganizationFieldCreateOptions specifices the optional parameters to the
// OrganizationFieldsService.Create method.
type OrganizationFieldCreateOptions struct {
	Name      string    `url:"name"`
	FieldType FieldType `url:"field_type"`
	Options   string    `url:"options"`
}

// Create a new organization field.
//
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

// OrganizationFieldUpdateOptions specifices the optional parameters to the
// OrganizationFieldsService.Update method.
type OrganizationFieldUpdateOptions struct {
	Name    string `url:"name"`
	Options string `url:"options"`
}

// Update a specific organization field.
//
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

// DeleteMultiple marks organization fields as deleted.
//
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

// Delete marks a specific organization field as deleted.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/OrganizationFields/delete_organizationFields_id
func (s *OrganizationFieldsService) Delete(ctx context.Context, id int) (*Response, error) {
	uri := fmt.Sprintf("/organizationFields/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
