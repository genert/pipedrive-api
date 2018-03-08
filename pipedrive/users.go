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

type UsersResponse struct {
	Success        bool           `json:"success"`
	Error          string         `json:"error,omitempty"`
	ErrorInfo      string         `json:"error_info,omitempty"`
	Data           []User         `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

type UserSingleResponse struct {
	Success        bool           `json:"success"`
	Data           User           `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

type UserFollowersResponse struct {
	Success        bool           `json:"success"`
	Data           []int          `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

type UserPermissionsResponse struct {
	Success bool `json:"success"`
	Data    struct {
		CanAddProducts              bool `json:"can_add_products"`
		CanBulkEditItems            bool `json:"can_bulk_edit_items"`
		CanChangeVisibilityOfItems  bool `json:"can_change_visibility_of_items"`
		CanDeleteActivities         bool `json:"can_delete_activities"`
		CanDeleteDeals              bool `json:"can_delete_deals"`
		CanEditDealsClosedDate      bool `json:"can_edit_deals_closed_date"`
		CanEditProducts             bool `json:"can_edit_products"`
		CanEditSharedFilters        bool `json:"can_edit_shared_filters"`
		CanExportDataFromLists      bool `json:"can_export_data_from_lists"`
		CanFollowOtherUsers         bool `json:"can_follow_other_users"`
		CanMergeDeals               bool `json:"can_merge_deals"`
		CanMergeOrganizations       bool `json:"can_merge_organizations"`
		CanMergePeople              bool `json:"can_merge_people"`
		CanSeeCompanyWideStatistics bool `json:"can_see_company_wide_statistics"`
		CanSeeDealsListSummary      bool `json:"can_see_deals_list_summary"`
		CanSeeHiddenItemsNames      bool `json:"can_see_hidden_items_names"`
		CanSeeOtherUsers            bool `json:"can_see_other_users"`
		CanSeeOtherUsersStatistics  bool `json:"can_see_other_users_statistics"`
		CanShareFilters             bool `json:"can_share_filters"`
		CanUseAPI                   bool `json:"can_use_api"`
	} `json:"data"`
}

type UserRoleSettingsResponse struct {
	Success bool `json:"success"`
	Data    struct {
		DealDefaultVisibility    int `json:"deal_default_visibility"`
		OrgDefaultVisibility     int `json:"org_default_visibility"`
		PersonDefaultVisibility  int `json:"person_default_visibility"`
		ProductDefaultVisibility int `json:"product_default_visibility"`
		DealAccessLevel          int `json:"deal_access_level"`
		OrgAccessLevel           int `json:"org_access_level"`
		PersonAccessLevel        int `json:"person_access_level"`
		ProductAccessLevel       int `json:"product_access_level"`
	} `json:"data"`
}

type UsersFindByNameOptions struct {
	Term          string `url:"term,omitempty"`
	SearchByEmail int    `url:"search_by_email,omitempty"`
}

type UsersUpdateUserDetailsOptions struct {
	ActiveFlag uint8 `url:"active_flag,omitempty"`
}

type DeletePermissionSetAssignmentOptions struct {
	PermissionSetId uint `url:"permission_set_id,omitempty"`
}

type DeleteRoleAssignmentOptions struct {
	RoleId uint `url:"role_id,omitempty"`
}

type UserCreateOptions struct {
	Name       string `url:"name"`
	Email      string `url:"email"`
	ActiveFlag uint8  `url:"active_flag"`
}

// https://developers.pipedrive.com/docs/api/v1/#!/Users/get_users_id_followers
func (s *UsersService) ListFollowers(id int) (*UserFollowersResponse, *Response, error) {
	uri := fmt.Sprintf("/users/%v/followers", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *UserFollowersResponse

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Users/get_users
func (s *UsersService) List() (*UsersResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/users", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *UsersResponse

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Users/post_users
func (s *UsersService) Create(opt *UserCreateOptions) (*UserSingleResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/users", nil, body)

	if err != nil {
		return nil, nil, err
	}

	var record *UserSingleResponse

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Users/get_users_find
func (s *UsersService) FindByName(opt *UsersFindByNameOptions) (*UsersResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/users/find", opt, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *UsersResponse

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Users/get_users_me
func (s *UsersService) GetCurrentUser() (*UserSingleResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/users/me", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *UserSingleResponse

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Users/get_users_id
func (s *UsersService) GetById(id int) (*UserFollowersResponse, *Response, error) {
	uri := fmt.Sprintf("/users/%v", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *UserFollowersResponse

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Users/get_users_id_permissions
func (s *UsersService) ListUserPermissions(id int) (*UserPermissionsResponse, *Response, error) {
	uri := fmt.Sprintf("/users/%v/permissions", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *UserPermissionsResponse

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Users/get_users_id_roleSettings
func (s *UsersService) ListUserRoleSettings(id int) (*UserRoleSettingsResponse, *Response, error) {
	uri := fmt.Sprintf("/users/%v/roleSettings", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *UserRoleSettingsResponse

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Users/put_users_id
func (s *UsersService) UpdateUserDetails(id int, opt *UsersUpdateUserDetailsOptions) (*Response, error) {
	uri := fmt.Sprintf("/users/%v", id)
	req, err := s.client.NewRequest(http.MethodPut, uri, opt, nil)

	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Users/delete_users_id_permissionSetAssignments
func (s *UsersService) DeletePermissionSetAssignment(id int, opt *DeletePermissionSetAssignmentOptions) (*Response, error) {
	uri := fmt.Sprintf("/users/%v/permissionSetAssignments", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, opt, nil)

	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Users/delete_users_id_roleAssignments
func (s *UsersService) DeleteRoleAssignment(id int, opt *DeleteRoleAssignmentOptions) (*Response, error) {
	uri := fmt.Sprintf("/users/%v/roleAssignments", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, opt, nil)

	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)

	if err != nil {
		return resp, err
	}

	return resp, nil
}
