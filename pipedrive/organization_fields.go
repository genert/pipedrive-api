package pipedrive

import "net/http"

type OrganizationsFieldsService service

type OrganizationFields struct {
	Success        bool                `json:"success"`
	Data           []OrganizationField `json:"data"`
	AdditionalData AdditionalData      `json:"additional_data"`
}

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

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/OrganizationFields/get_organizationFields
func (s *OrganizationsFieldsService) List() (*OrganizationFields, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/organizationFields", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *OrganizationFields

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}
