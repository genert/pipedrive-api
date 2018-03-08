package pipedrive

import (
	"fmt"
	"net/http"
)

type NotesService service

type Note struct {
	Id                       int       `json:"id,omitempty"`
	UserId                   int       `json:"user_id,omitempty"`
	DealId                   int       `json:"deal_id,omitempty"`
	PersonId                 int       `json:"person_id,omitempty"`
	OrgId                    int       `json:"org_id,omitempty"`
	Content                  string    `json:"content,omitempty"`
	AddTime                  Timestamp `json:"add_time,omitempty"`
	UpdateTime               Timestamp `json:"update_time,omitempty"`
	ActiveFlag               bool      `json:"active_flag,omitempty"`
	PinnedToDealFlag         bool      `json:"pinned_to_deal_flag,omitempty"`
	PinnedToPersonFlag       bool      `json:"pinned_to_person_flag,omitempty"`
	PinnedToOrganizationFlag bool      `json:"pinned_to_organization_flag,omitempty"`
	LastUpdateUserId         int       `json:"last_update_user_id,omitempty"`
}

type NotesResponse struct {
	Success        bool           `json:"success,omitempty"`
	Data           []Note         `json:"data,omitempty"`
	AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

type NoteResponse struct {
	Success bool `json:"success,omitempty"`
	Data    Note `json:"data,omitempty"`
}

type NoteCreateOptions struct {
	DealId                   uint   `url:"deal_id"`
	Content                  string `url:"content"`
	PersonId                 uint   `url:"person_id"`
	OrgId                    uint   `url:"org_id"`
	PinnedToDealFlag         uint8  `url:"pinned_to_deal_flag"`
	PinnedToOrganizationFlag uint8  `url:"pinned_to_organization_flag"`
	PinnedToPersonFlag       uint8  `url:"pinned_to_person_flag"`
}

type NoteUpdateOptions struct {
	Content                  string `url:"content"`
	DealId                   uint   `url:"deal_id"`
	PersonId                 uint   `url:"person_id"`
	OrgId                    uint   `url:"org_id"`
	PinnedToDealFlag         uint8  `url:"pinned_to_deal_flag"`
	PinnedToOrganizationFlag uint8  `url:"pinned_to_organization_flag"`
	PinnedToPersonFlag       uint8  `url:"pinned_to_person_flag"`
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Notes/get_notes
func (s *NotesService) List() (*NotesResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/notes", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *NotesResponse

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Notes/get_notes_id
func (s *NotesService) GetById(id int) (*NoteResponse, *Response, error) {
	uri := fmt.Sprintf("/notes/%v", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *NoteResponse

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Notes/get_notes_id
func (s *NotesService) Create(opt *NoteCreateOptions) (*NoteResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/filters", nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *NoteResponse

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Notes/put_notes_id
func (s *NotesService) Update(id int, opt *NoteUpdateOptions) (*NoteResponse, *Response, error) {
	uri := fmt.Sprintf("/filters/%v", id)
	req, err := s.client.NewRequest(http.MethodPut, uri, nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *NoteResponse

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Notes/delete_notes_id
func (s *NotesService) Delete(id int) (*Response, error) {
	uri := fmt.Sprintf("/notes/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
