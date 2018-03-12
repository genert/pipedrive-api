package pipedrive

import (
	"context"
	"fmt"
	"net/http"
)

// UsersService handles users related
// methods of the Pipedrive API.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Users
type UsersService service

// User represents a Pipedrive user.
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

func (u User) String() string {
	return Stringify(u)
}

// UsersResponse represents multiple users response.
type UsersResponse struct {
	Success        bool           `json:"success"`
	Error          string         `json:"error,omitempty"`
	ErrorInfo      string         `json:"error_info,omitempty"`
	Data           []User         `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

// UserSingleResponse represents single user response.
type UserSingleResponse struct {
	Success        bool           `json:"success"`
	Data           User           `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

// UserFollowersResponse represents user followers response.
type UserFollowersResponse struct {
	Success        bool           `json:"success"`
	Data           []int          `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

// UserPermissionsResponse represents user permissions response.
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

// UserRoleSettingsResponse represents user role settings response.
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

// ListFollowers lists followers of a specific user.
//
// https://developers.pipedrive.com/docs/api/v1/#!/Users/get_users_id_followers
func (s *UsersService) ListFollowers(ctx context.Context, id int) (*UserFollowersResponse, *Response, error) {
	uri := fmt.Sprintf("/users/%v/followers", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *UserFollowersResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// List returns data about all users within the company.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Users/get_users
func (s *UsersService) List(ctx context.Context) (*UsersResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/users", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *UsersResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// UserCreateOptions specifices the optional parameters to the
// UsersService.Create method.
type UserCreateOptions struct {
	Name       string `url:"name"`
	Email      string `url:"email"`
	ActiveFlag uint8  `url:"active_flag"`
}

// Create a user.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Users/post_users
func (s *UsersService) Create(ctx context.Context, opt *UserCreateOptions) (*UserSingleResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/users", nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *UserSingleResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// UsersFindByNameOptions specifices the optional parameters to the
// UsersService.FindByName method.
type UsersFindByNameOptions struct {
	Term          string `url:"term,omitempty"`
	SearchByEmail int    `url:"search_by_email,omitempty"`
}

// FindByName finds users by their name.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Users/get_users_find
func (s *UsersService) FindByName(ctx context.Context, opt *UsersFindByNameOptions) (*UsersResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/users/find", opt, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *UsersResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// GetCurrentUserData returns data about an authorized user within the company.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Users/get_users_me
func (s *UsersService) GetCurrentUserData(ctx context.Context) (*UserSingleResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/users/me", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *UserSingleResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// GetByID returns specific user.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Users/get_users_id
func (s *UsersService) GetByID(ctx context.Context, id int) (*UserFollowersResponse, *Response, error) {
	uri := fmt.Sprintf("/users/%v", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *UserFollowersResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// ListUserPermissions lists aggregated permissions over all assigned permission sets for a user.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Users/get_users_id_permissions
func (s *UsersService) ListUserPermissions(ctx context.Context, id int) (*UserPermissionsResponse, *Response, error) {
	uri := fmt.Sprintf("/users/%v/permissions", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *UserPermissionsResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// ListUserRoleSettings lists settings of user's assigned role.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Users/get_users_id_roleSettings
func (s *UsersService) ListUserRoleSettings(ctx context.Context, id int) (*UserRoleSettingsResponse, *Response, error) {
	uri := fmt.Sprintf("/users/%v/roleSettings", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *UserRoleSettingsResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// UsersUpdateUserDetailsOptions specifices the optional parameters to the
// UsersService.UpdateUserDetails method.
type UsersUpdateUserDetailsOptions struct {
	ActiveFlag uint8 `url:"active_flag,omitempty"`
}

// UpdateUserDetails updates the properties of a user. Currently, only active_flag can be updated.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Users/put_users_id
func (s *UsersService) UpdateUserDetails(ctx context.Context, id int, opt *UsersUpdateUserDetailsOptions) (*Response, error) {
	uri := fmt.Sprintf("/users/%v", id)
	req, err := s.client.NewRequest(http.MethodPut, uri, opt, nil)

	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeletePermissionSetAssignmentOptions specifices the optional parameters to the
// UsersService.DeletePermissionSetAssignment method.
type DeletePermissionSetAssignmentOptions struct {
	PermissionSetID uint `url:"permission_set_id,omitempty"`
}

// DeletePermissionSetAssignment deletes a permission set assignment for a user.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Users/delete_users_id_permissionSetAssignments
func (s *UsersService) DeletePermissionSetAssignment(ctx context.Context, id int, opt *DeletePermissionSetAssignmentOptions) (*Response, error) {
	uri := fmt.Sprintf("/users/%v/permissionSetAssignments", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, opt, nil)

	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteRoleAssignmentOptions specifices the optional parameters to the
// UsersService.DeleteRoleAssignment method.
type DeleteRoleAssignmentOptions struct {
	RoleID uint `url:"role_id,omitempty"`
}

// DeleteRoleAssignment deletes a role assignment for a user.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Users/delete_users_id_roleAssignments
func (s *UsersService) DeleteRoleAssignment(ctx context.Context, id int, opt *DeleteRoleAssignmentOptions) (*Response, error) {
	uri := fmt.Sprintf("/users/%v/roleAssignments", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, opt, nil)

	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	if err != nil {
		return resp, err
	}

	return resp, nil
}
