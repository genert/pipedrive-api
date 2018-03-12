package pipedrive

import (
	"context"
	"net/http"
)

// UserConnectionsService handles activities related
// methods of the Pipedrive API.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/UserConnections
type UserConnectionsService service

// UserConnections represents a Pipedrive user connections.
type UserConnections struct {
	Success bool `json:"success"`
	Data    struct {
		Google string `json:"google"`
	} `json:"data"`
}

// List returns data about all connections for authorized user.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/UserConnections/get_userConnections
func (s *UserConnectionsService) List(ctx context.Context) (*UserConnections, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/userConnections", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *UserConnections

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}
