package pipedrive

import (
	"fmt"
	"net/http"
)

type UsersService service

type User struct {
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

type Users struct {
	Success        bool           `json:"success"`
	Error          string         `json:"error,omitempty"`
	ErrorInfo      string         `json:"error_info,omitempty"`
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

type DeleteRoleAssignmentOptions struct {
	RoleId uint `url:"role_id,omitempty"`
}

// Returns data about all users within the company.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Users/get_users
func (s *UsersService) List() (*Users, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/users", nil, nil)

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
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Users/get_users_find
func (s *UsersService) FindByName(opt *UsersFindByNameOptions) (*Users, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/users/find", opt, nil)

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
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Users/get_users_id
func (s *UsersService) GetById(id int) (*SingleUser, *Response, error) {
	uri := fmt.Sprintf("/users/%v", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

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

// Delete a role assignment for a user.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Users/delete_users_id_roleAssignments
func (s *UsersService) DeleteRoleAssignment(id int, opt *DeleteRoleAssignmentOptions) (*Response, error) {
	uri := fmt.Sprintf("/users/%v/roleAssignments", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)

	if err != nil {
		return resp, err
	}

	return resp, nil
}
