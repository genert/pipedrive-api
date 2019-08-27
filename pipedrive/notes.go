package pipedrive

import (
	"context"
	"fmt"
	"net/http"
)

// NotesService handles activities related
// methods of the Pipedrive API.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Notes
type NotesService service

// Note represents a Pipedrive note.
type Note struct {
	ID                       int       `json:"id,omitempty"`
	UserID                   int       `json:"user_id,omitempty"`
	DealID                   int       `json:"deal_id,omitempty"`
	PersonID                 int       `json:"person_id,omitempty"`
	OrgID                    int       `json:"org_id,omitempty"`
	Content                  string    `json:"content,omitempty"`
	AddTime                  Timestamp `json:"add_time,omitempty"`
	UpdateTime               Timestamp `json:"update_time,omitempty"`
	ActiveFlag               bool      `json:"active_flag,omitempty"`
	PinnedToDealFlag         bool      `json:"pinned_to_deal_flag,omitempty"`
	PinnedToPersonFlag       bool      `json:"pinned_to_person_flag,omitempty"`
	PinnedToOrganizationFlag bool      `json:"pinned_to_organization_flag,omitempty"`
	LastUpdateUserID         int       `json:"last_update_user_id,omitempty"`
}

func (n Note) String() string {
	return Stringify(n)
}

// NotesResponse represents multiple notes response.
type NotesResponse struct {
	Success        bool           `json:"success,omitempty"`
	Data           []Note         `json:"data,omitempty"`
	AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

// NoteResponse represents a single note response.
type NoteResponse struct {
	Success bool `json:"success,omitempty"`
	Data    Note `json:"data,omitempty"`
}

// List returns notes.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Notes/get_notes
func (s *NotesService) List(ctx context.Context) (*NotesResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/notes", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *NotesResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// GetByID returns a specific note by id.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Notes/get_notes_id
func (s *NotesService) GetByID(ctx context.Context, id int) (*NoteResponse, *Response, error) {
	uri := fmt.Sprintf("/notes/%v", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *NoteResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// NoteCreateOptions specifices the optional parameters to the
// NotesService.Create method.
type NoteCreateOptions struct {
	DealID                   uint   `json:"deal_id"`
	Content                  string `json:"content"`
	PersonID                 uint   `json:"person_id"`
	OrgID                    uint   `json:"org_id"`
	PinnedToDealFlag         uint8  `json:"pinned_to_deal_flag"`
	PinnedToOrganizationFlag uint8  `json:"pinned_to_organization_flag"`
	PinnedToPersonFlag       uint8  `json:"pinned_to_person_flag"`
}

// Create a note.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Notes/get_notes_id
func (s *NotesService) Create(ctx context.Context, opt *NoteCreateOptions) (*NoteResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/notes", nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *NoteResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// NoteUpdateOptions specifices the optional parameters to the
// NotesService.Update method.
type NoteUpdateOptions struct {
	Content                  string `url:"content"`
	DealID                   uint   `url:"deal_id"`
	PersonID                 uint   `url:"person_id"`
	OrgID                    uint   `url:"org_id"`
	PinnedToDealFlag         uint8  `url:"pinned_to_deal_flag"`
	PinnedToOrganizationFlag uint8  `url:"pinned_to_organization_flag"`
	PinnedToPersonFlag       uint8  `url:"pinned_to_person_flag"`
}

// Update a specific note.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Notes/put_notes_id
func (s *NotesService) Update(ctx context.Context, id int, opt *NoteUpdateOptions) (*NoteResponse, *Response, error) {
	uri := fmt.Sprintf("/notes/%v", id)
	req, err := s.client.NewRequest(http.MethodPut, uri, nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *NoteResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Delete marks note as deleted.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Notes/delete_notes_id
func (s *NotesService) Delete(ctx context.Context, id int) (*Response, error) {
	uri := fmt.Sprintf("/notes/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
