package pipedrive

import (
	"fmt"
	"net/http"
)

type DealService service

type Deal struct {
	ID         int    `json:"id,omitempty"`
	StageId    int    `json:"stage_id,omitempty"`
	Title      string `json:"title,omitempty"`
	Value      int    `json:"value,omitempty"`
	Currency   string `json:"currency,omitempty"`
	AddTime    string `json:"add_time,omitempty"`
	UpdateTime string `json:"update:time,omitempty"`
}

type Deals struct {
	Success bool   `json:"success,omitempty"`
	Data    []Deal `json:"data,omitempty"`
}

type DealUpdate struct {
	Success bool `json:"success,omitempty"`
	Data    Deal `json:"data,omitempty"`
}

type DealsMergeOptions struct {
	Id          uint `url:"id"`
	MergeWithId uint `url:"merge_with_id"`
}

type DealsUpdateOptions struct {
	Id             uint   `url:"id"`
	Title          string `url:"title,omitempty"`
	Value          string `url:"value,omitempty"`
	Currency       string `url:"currency,omitempty"`
	UserId         uint   `url:"user_id,omitempty"`
	PersonId       uint   `url:"person_id,omitempty"`
	OrganizationId uint   `url:"org_id,omitempty"`
	StageId        uint   `url:"stage_id,omitempty"`
	Status         string `url:"status,omitempty"`
	LostReason     string `url:"lost_reason,omitempty"`
	VisibleTo      uint   `url:"visible_to,omitempty"`
}

// List updates about a deal
func (s *DealService) ListDealUpdates(id int) (*Deals, *Response, error) {
	uri := fmt.Sprintf("/deals/%v/flow", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *Deals

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

func (s *DealService) List() (*Deals, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/deals", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *Deals

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Duplicate a deal.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/post_deals_id_duplicate
func (s *DealService) Duplicate(id int) (*DealUpdate, *Response, error) {
	uri := fmt.Sprintf("/deals/%v/duplicate", id)
	req, err := s.client.NewRequest(http.MethodPost, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *DealUpdate

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Merges a deal with another deal.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/put_deals_id_merge
func (s *DealService) Merge(opt *DealsMergeOptions) (*Response, error) {
	uri := fmt.Sprintf("/deals/%v/duplicate", opt.Id)
	req, err := s.client.NewRequest(http.MethodPut, uri, opt, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// Updates the properties of a deal.
// https://developers.pipedrive.com/docs/api/v1/#!/Deals/put_deals_id
func (s *DealService) Update(opt *DealsUpdateOptions) (*Response, error) {
	uri := fmt.Sprintf("/deals/%v", opt.Id)
	req, err := s.client.NewRequest(http.MethodPut, uri, opt, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// Deletes a follower from a deal.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/delete_deals_id_followers_follower_id
func (s *DealService) DeleteFollower(id uint, followerId uint) (*Response, error) {
	uri := fmt.Sprintf("/deals/%v/followers/%v", id, followerId)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// Marks multiple deals as deleted.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/delete_deals
func (s *DealService) DeleteMultiple(ids []int) (*Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, "/deals", &DeleteMultipleOptions{
		Ids: arrayToString(ids, ","),
	}, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// Deletes a participant from a deal.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/delete_deals_id_participants_deal_participant_id
func (s *DealService) DeleteParticipant(dealId uint, participantId uint) (*Response, error) {
	uri := fmt.Sprintf("/deals/%v/participants/%v", dealId, participantId)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// Marks a deal as deleted.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/delete_deals_id
func (s *DealService) Delete(id uint) (*Response, error) {
	uri := fmt.Sprintf("/deals/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// Deletes a product attachment from a deal, using the product_attachment_id.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/delete_deals_id_products_product_attachment_id
func (s *DealService) DeleteAttachedProduct(dealId uint, productAttachmentId uint) (*Response, error) {
	uri := fmt.Sprintf("/deals/%v/products/%v", dealId, productAttachmentId)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
