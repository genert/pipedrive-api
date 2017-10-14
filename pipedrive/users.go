package pipedrive

import "fmt"

type UsersService service

type User struct {
	ID                  int         `json:"id"`
	Name                string      `json:"name"`
	DefaultCurrency     string      `json:"default_currency"`
	Locale              string      `json:"locale"`
	Lang                int         `json:"lang"`
	Email               string      `json:"email"`
	Phone               interface{} `json:"phone"`
	Activated           bool        `json:"activated"`
	LastLogin           string      `json:"last_login"`
	Created             string      `json:"created"`
	Modified            string      `json:"modified"`
	SignupFlowVariation string      `json:"signup_flow_variation"`
	HasCreatedCompany   bool        `json:"has_created_company"`
	IsAdmin             int         `json:"is_admin"`
	TimezoneName        string      `json:"timezone_name"`
	TimezoneOffset      string      `json:"timezone_offset"`
	ActiveFlag          bool        `json:"active_flag"`
	RoleID              int         `json:"role_id"`
	IconURL             interface{} `json:"icon_url"`
	IsYou               bool        `json:"is_you"`
}

type Users struct {
	Success        bool           `json:"success"`
	Data           []User         `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

type SingleUser struct {
	Success        bool           `json:"success"`
	Data           User           `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

type UsersFindByNameOptions struct {
	Term          string `url:"term,omitempty"`
	SearchByEmail int    `url:"search_by_email,omitempty"`
}

// Returns data about all users within the company.
// https://developers.pipedrive.com/docs/api/v1/#!/Users/get_users
func (s *UsersService) List() (*Users, *Response, error) {
	uri, err := s.client.CreateRequestUrl("/users", nil)

	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", uri, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *Users

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Finds users by their name.
// https://developers.pipedrive.com/docs/api/v1/#!/Users/get_users_find
func (s *UsersService) FindByName(opt *UsersFindByNameOptions) (*Users, *Response, error) {
	uri, err := s.client.CreateRequestUrl("/users/find", opt)

	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", uri, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *Users

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Returns data about a specific user within the company
// https://developers.pipedrive.com/docs/api/v1/#!/Users/get_users_id
func (s *UsersService) GetById(id int) (*SingleUser, *Response, error) {
	uri, err := s.client.CreateRequestUrl(fmt.Sprintf("/users/%v", id), nil)

	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", uri, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *SingleUser

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}
