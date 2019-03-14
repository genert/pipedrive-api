package pipedrive

import (
	"context"
	"fmt"
	"net/http"
)

// ActivitiesService handles activities related
// methods of the Pipedrive API.
//
// Pipedrive API dcos: https://developers.pipedrive.com/docs/api/v1/#!/Activities
type ActivitiesService service

// Participants represents a Pipedrive participant.
type Participants struct {
	PersonID    int  `json:"person_id"`
	PrimaryFlag bool `json:"primary_flag"`
}

// Activity represents a Pipedrive activity.
type Activity struct {
	ID                 int            `json:"id"`
	CompanyID          int            `json:"company_id"`
	UserID             int            `json:"user_id"`
	Done               bool           `json:"done"`
	Type               string         `json:"type"`
	ReferenceType      string         `json:"reference_type"`
	ReferenceID        int            `json:"reference_id"`
	DueDate            string         `json:"due_date"`
	DueTime            string         `json:"due_time"`
	Duration           string         `json:"duration"`
	AddTime            string         `json:"add_time"`
	MarkedAsDoneTime   string         `json:"marked_as_done_time"`
	Subject            string         `json:"subject"`
	OrgID              int            `json:"org_id"`
	PersonID           int            `json:"person_id"`
	DealID             int            `json:"deal_id"`
	ActiveFlag         bool           `json:"active_flag"`
	UpdateTime         string         `json:"update_time"`
	GcalEventID        interface{}    `json:"gcal_event_id"`
	GoogleCalendarID   interface{}    `json:"google_calendar_id"`
	GoogleCalendarEtag interface{}    `json:"google_calendar_etag"`
	Note               string         `json:"note"`
	CreatedByUserID    int            `json:"created_by_user_id"`
	Participants       []Participants `json:"participants"`
	OrgName            string         `json:"org_name"`
	PersonName         string         `json:"person_name"`
	DealTitle          string         `json:"deal_title"`
	OwnerName          string         `json:"owner_name"`
	PersonDropboxBcc   string         `json:"person_dropbox_bcc"`
	DealDropboxBcc     string         `json:"deal_dropbox_bcc"`
	AssignedToUserID   int            `json:"assigned_to_user_id"`
}

func (a Activity) String() string {
	return Stringify(a)
}

// ActivityResponse represents single activity response.
type ActivityResponse struct {
	Success bool     `json:"success"`
	Data    Activity `json:"data"`
}

// ActivitiesReponse represents multiple activities response.
type ActivitiesReponse struct {
	Success        bool           `json:"success"`
	Data           []Activity     `json:"data"`
	AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

// List returns all activities assigned to a particular user
//
// https://developers.pipedrive.com/docs/api/v1/#!/Activities/get_activities
func (s *ActivitiesService) List(ctx context.Context) (*ActivitiesReponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/activities", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *ActivitiesReponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// GetByID returns details of a specific activity.
//
// https://developers.pipedrive.com/docs/api/v1/#!/Activities/get_activities
func (s *ActivitiesService) GetByID(ctx context.Context, id int) (*ActivitiesReponse, *Response, error) {
	uri := fmt.Sprintf("/activities/%v", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *ActivitiesReponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Create an activity.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Activities/post_activities
func (s *ActivitiesService) Create(ctx context.Context, opt *ActivitiesCreateOptions) (*ActivityResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/activities", nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *ActivityResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// ActivitiesCreateOptions specifices the optional parameters to the
// ActivitiesService.Update method.
type ActivitiesCreateOptions struct {
	Subject      string      `json:"subject,omitempty"`
	Done         uint8       `json:"done,omitempty"`
	Type         string      `json:"type,omitempty"`
	DueDate      string      `json:"due_date,omitempty"`
	DueTime      string      `json:"due_time,omitempty"`
	Duration     string      `json:"duration,omitempty"`
	UserID       uint        `json:"user_id,omitempty"`
	DealID       uint        `json:"user_id,omitempty"`
	PersonID     uint        `json:"person_id,omitempty"`
	Participants interface{} `json:"participants,omitempty"`
	OrgID        uint        `json:"org_id,omitempty"`
}

// Update an activity
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Activities/put_activities_id
func (s *ActivitiesService) Update(ctx context.Context, id int, opt *ActivitiesCreateOptions) (*ActivityResponse, *Response, error) {
	uri := fmt.Sprintf("/activities/%v", id)
	req, err := s.client.NewRequest(http.MethodPut, uri, opt, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *ActivityResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// DeleteMultiple activities in bulk.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Activities/delete_activities
func (s *ActivitiesService) DeleteMultiple(ctx context.Context, ids []int) (*Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, "/activities", &DeleteMultipleOptions{
		Ids: arrayToString(ids, ","),
	}, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Delete an activity.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Activities/delete_activities_id
func (s *ActivitiesService) Delete(ctx context.Context, id int) (*Response, error) {
	uri := fmt.Sprintf("/activities/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
