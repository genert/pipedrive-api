package pipedrive

import (
	"context"
	"net/http"
)

// RecentsService handles activities related
// methods of the Pipedrive API.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Recents
type RecentsService service

// RecentRecordDetails represents a Pipedrive recent record details.
type RecentRecordDetails struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	DefaultCurrency     string `json:"default_currency"`
	Locale              string `json:"locale"`
	Lang                int    `json:"lang"`
	Email               string `json:"email"`
	Phone               string `json:"phone"`
	Activated           bool   `json:"activated"`
	LastLogin           string `json:"last_login"`
	Created             string `json:"created"`
	Modified            string `json:"modified"`
	SignupFlowVariation string `json:"signup_flow_variation"`
	HasCreatedCompany   bool   `json:"has_created_company"`
	IsAdmin             int    `json:"is_admin"`
	TimezoneName        string `json:"timezone_name"`
	TimezoneOffset      string `json:"timezone_offset"`
	ActiveFlag          bool   `json:"active_flag"`
	RoleID              int    `json:"role_id"`
	IconURL             string `json:"icon_url"`
	IsYou               bool   `json:"is_you"`
}

func (rrd RecentRecordDetails) String() string {
	return Stringify(rrd)
}

// RecentRecord represents a Pipedrive recent record.
type RecentRecord struct {
	Item string                `json:"item"`
	ID   int                   `json:"id"`
	Data []RecentRecordDetails `json:"data"`
}

// RecentsResponse represents multiple recents response.
type RecentsResponse struct {
	Success        bool           `json:"success"`
	Data           []RecentRecord `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

// RecentsListOptions specifices the optional parameters to the
// RecentsService.List method.
type RecentsListOptions struct {
	SinceTimestamp string `url:"since_timestamp,omitempty"`
	Items          string `url:"items,omitempty"`
	Start          uint   `url:"start,omitempty"`
	Limit          uint   `url:"limit,omitempty"`
}

// List returns data about all recent changes occured after given timestamp.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Recents/get_recents
func (s *RecentsService) List(ctx context.Context, opt *RecentsListOptions) (*RecentsResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/recents", opt, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *RecentsResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}
