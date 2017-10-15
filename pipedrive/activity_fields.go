package pipedrive

import "net/http"

type ActivityFieldsService service

type ActivityField struct {
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
	MandatoryFlag      bool        `json:"mandatory_flag"`
	Options            []struct {
		ID    string `json:"id"`
		Label string `json:"label"`
	} `json:"options,omitempty"`
}

type ActivityFields struct {
	Success        bool            `json:"success"`
	Data           []ActivityField `json:"data"`
	AdditionalData AdditionalData  `json:"additional_data"`
}

// Return list of all fields for activity.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/ActivityFields/get_activityFields
func (s *ActivityFieldsService) List() (*ActivityFields, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/activityFields", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *ActivityFields

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}
