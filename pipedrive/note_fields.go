package pipedrive

import (
	"context"
	"net/http"
)

// NoteFieldsService handles note field related
// methods of the Pipedrive API.
//
// Pipedrive API dcos: https://developers.pipedrive.com/docs/api/v1/#!/ActivityTypes
type NoteFieldsService service

// Option represents a Pipedrive option for note field.
type Option struct {
	ID    int    `json:"id,omitempty"`
	Label string `json:"label,omitempty"`
}

// NoteField represents a Pipedrive note field.
type NoteField struct {
	ID                   int      `json:"id,omitempty"`
	Key                  string   `json:"key,omitempty"`
	Name                 string   `json:"name,omitempty"`
	ActiveFlag           bool     `json:"active_flag,omitempty"`
	FieldType            string   `json:"field_type,omitempty"`
	EditFlag             int      `json:"edit_flag,omitempty"`
	MandatoryFlag        bool     `json:"mandatory_flag,omitempty"`
	VisibleInExportsFlag bool     `json:"visible_in_exports_flag,omitempty"`
	Options              []Option `json:"options,omitempty"`
}

func (nf NoteField) String() string {
	return Stringify(nf)
}

// NoteFieldsResponse represents multiple note fields esponse.
type NoteFieldsResponse struct {
	Success        bool           `json:"success,omitempty"`
	Data           []NoteField    `json:"data,omitempty"`
	AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

// List returns all fields for note.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/NoteFields/get_noteFields
func (s *NoteFieldsService) List(ctx context.Context) (*NoteFieldsResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/noteFields", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *NoteFieldsResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}
