package pipedrive

import (
	"fmt"
	"net/http"
)

type DealFields service

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

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/DealFields/get_dealFields
func (s *DealFields) List() (*DealFieldsResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/dealFields", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *DealFieldsResponse

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

type DealFieldsUpdateOptions struct {
	Id      uint
	Name    string `url:"name"`
	Options string `url:"options,omitempty"`
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/DealFields/delete_dealFields
func (s *DealFields) DeleteMultiple(ids []int) (*Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, "/dealFields", &DeleteMultipleOptions{
		Ids: arrayToString(ids, ","),
	}, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/DealFields/delete_dealFields_id
func (s *DealFields) Delete(id int) (*Response, error) {
	uri := fmt.Sprintf("/dealFields/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
