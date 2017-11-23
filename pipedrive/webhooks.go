package pipedrive

import (
	"fmt"
	"net/http"
	"time"
)

type WebhooksService service

type Webhook struct {
	ID               int         `json:"id"`
	CompanyID        int         `json:"company_id"`
	OwnerID          int         `json:"owner_id"`
	UserID           int         `json:"user_id"`
	EventAction      string      `json:"event_action"`
	EventObject      string      `json:"event_object"`
	SubscriptionURL  string      `json:"subscription_url"`
	IsActive         int         `json:"is_active"`
	AddTime          time.Time   `json:"add_time"`
	RemoveTime       interface{} `json:"remove_time"`
	Type             string      `json:"type"`
	HTTPAuthUser     interface{} `json:"http_auth_user"`
	HTTPAuthPassword interface{} `json:"http_auth_password"`
	AdditionalData   struct{}    `json:"additional_data"`
	LastDeliveryTime time.Time   `json:"last_delivery_time"`
	LastHTTPStatus   int         `json:"last_http_status"`
	AdminID          int         `json:"admin_id"`
}

type Webhooks struct {
	Status  string    `json:"status,omitempty"`
	Success bool      `json:"success,omitempty"`
	Data    []Webhook `json:"data,omitempty"`
}

type WebhooksCreateOptions struct {
	SubscriptionUrl  string      `url:"subscription_url"`
	EventAction      EventAction `url:"event_action"`
	DealProbability  EventObject `url:"event_object"`
	UserId           uint        `url:"user_id"`
	HttpAuthUser     string      `url:"http_auth_user"`
	HttpAuthPassword string      `url:"http_auth_password"`
}

// Returns data about all webhooks of a company.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Webhooks/get_webhooks
func (s *WebhooksService) List() (*Webhooks, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/webhooks", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *Webhooks

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Creates a new webhook and returns its details.
// Note that specifying an event which triggers the webhook combines 2 parameters -
// 'event_action' and 'event_object'.
// E.g., use '*.*' for getting notifications about all events,
// 'added.deal' for any newly added deals, 'deleted.persons' for any deleted persons, etc.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Webhooks/post_webhooks
func (s *WebhooksService) Create(opt *WebhooksCreateOptions) (*SingleStage, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/webhooks", nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *SingleStage

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Deletes the specified webhook.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Webhooks/delete_webhooks_id
func (s *WebhooksService) Delete(id int) (*Response, error) {
	uri := fmt.Sprintf("/webhooks/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
