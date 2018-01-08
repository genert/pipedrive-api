package pipedrive

import (
	"fmt"
	"net/http"
)

type ActivitiesService service

type Participants struct {
	PersonID    int  `json:"person_id"`
	PrimaryFlag bool `json:"primary_flag"`
}

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

type SingleActivity struct {
	Success bool     `json:"success"`
	Data    Activity `json:"data"`
}

type Activities struct {
	Success        bool           `json:"success"`
	Data           []Activity     `json:"data"`
	AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

// Returns all activities assigned to a particular user
// https://developers.pipedrive.com/docs/api/v1/#!/Activities/get_activities
func (s *ActivitiesService) List(id int) (*Activities, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/activities", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *Activities

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Returns details of a specific activity.
// https://developers.pipedrive.com/docs/api/v1/#!/Activities/get_activities
func (s *ActivitiesService) GetById(id int) (*Activities, *Response, error) {
	uri := fmt.Sprintf("/activities/%v", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *Activities

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

type ActivitiesCreateOptions struct {
	Subject      string      `url:"subject"`
	Done         uint8       `url:"done"`
	Type         string      `url:"type"`
	DueDate      string      `url:"due_date"`
	DueTime      string      `url:"due_time"`
	Duration     string      `url:"duration"`
	UserId       uint        `url:"user_id"`
	DealId       uint        `url:"user_id"`
	PersonId     uint        `url:"person_id"`
	Participants interface{} `url:"participants"`
	OrgId        uint        `url:"org_id"`
}

// Create an activity.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Activities/post_activities
func (s *ActivitiesService) Create(opt *ActivitiesCreateOptions) (*SingleActivity, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/activities", nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *SingleActivity

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Edit an activity
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Activities/put_activities_id
func (s *ActivitiesService) Update(id int, opt *ActivitiesCreateOptions) (*SingleActivity, *Response, error) {
	uri := fmt.Sprintf("/activities/%v", id)
	req, err := s.client.NewRequest(http.MethodPut, uri, opt, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *SingleActivity

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Delete multiple activities in bulk.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Activities/delete_activities
func (s *ActivitiesService) DeleteMultiple(ids []int) (*Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, "/activities", &DeleteMultipleOptions{
		Ids: arrayToString(ids, ","),
	}, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// Deletes an activity.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Activities/delete_activities_id
func (s *ActivitiesService) Delete(id int) (*Response, error) {
	uri := fmt.Sprintf("/activities/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
