package pipedrive

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// WebhooksService handles webhooks related
// methods of the Pipedrive API.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Webhooks
type WebhooksService service

// Webhook represents a Pipedrive webhook.
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

func (w Webhook) String() string {
	return Stringify(w)
}

// WebhooksResponse represents multiple webhooks response.
type WebhooksResponse struct {
	Status  string    `json:"status,omitempty"`
	Success bool      `json:"success,omitempty"`
	Data    []Webhook `json:"data,omitempty"`
}

// WebhookResponse represents single webhook response.
type WebhookResponse struct {
	Status  string  `json:"status,omitempty"`
	Success bool    `json:"success,omitempty"`
	Data    Webhook `json:"data,omitempty"`
}

// List all webhooks.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Webhooks/get_webhooks
func (s *WebhooksService) List(ctx context.Context) (*WebhooksResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/webhooks", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *WebhooksResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// WebhooksCreateOptions specifices the optional parameters to the
// WebhooksService.Create method.
type WebhooksCreateOptions struct {
	SubscriptionURL  string      `url:"subscription_url"`
	EventAction      EventAction `url:"event_action"`
	DealProbability  EventObject `url:"event_object"`
	UserID           uint        `url:"user_id"`
	HTTPAuthUser     string      `url:"http_auth_user"`
	HTTPAuthPassword string      `url:"http_auth_password"`
}

// Create a webhook.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Webhooks/post_webhooks
func (s *WebhooksService) Create(ctx context.Context, opt *WebhooksCreateOptions) (*WebhookResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/webhooks", nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *WebhookResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Delete a webhook.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Webhooks/delete_webhooks_id
func (s *WebhooksService) Delete(ctx context.Context, id int) (*Response, error) {
	uri := fmt.Sprintf("/webhooks/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
