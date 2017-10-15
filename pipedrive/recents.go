package pipedrive

import "net/http"

type RecentsService service

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

type RecentRecord struct {
	Item string                `json:"item"`
	ID   int                   `json:"id"`
	Data []RecentRecordDetails `json:"data"`
}

type Recents struct {
	Success        bool           `json:"success"`
	Data           []RecentRecord `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

type RecentsListOptions struct {
	SinceTimestamp string `url:"since_timestamp,omitempty"`
	Items          string `url:"items,omitempty"`
	Start          uint   `url:"start,omitempty"`
	Limit          uint   `url:"limit,omitempty"`
}

// Returns data about all recent changes occured after given timestamp.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Recents/get_recents
func (s *RecentsService) List(opt *RecentsListOptions) (*Recents, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/recents", opt, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *Recents

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}
