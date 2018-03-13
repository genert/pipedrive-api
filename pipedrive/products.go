package pipedrive

import (
	"context"
	"fmt"
	"net/http"
)

// ProductsService handles pipelines related
// methods of the Pipedrive API.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Products
type ProductsService service

// Product represents a Pipedrive product.
type Product struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Code       interface{} `json:"code"`
	Unit       string      `json:"unit"`
	Tax        int         `json:"tax"`
	ActiveFlag bool        `json:"active_flag"`
	Selectable bool        `json:"selectable"`
	FirstChar  string      `json:"first_char"`
	VisibleTo  string      `json:"visible_to"`
	OwnerID    struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Email      string `json:"email"`
		HasPic     bool   `json:"has_pic"`
		PicHash    string `json:"pic_hash"`
		ActiveFlag bool   `json:"active_flag"`
		Value      int    `json:"value"`
	} `json:"owner_id"`
	FilesCount     interface{} `json:"files_count"`
	FollowersCount int         `json:"followers_count"`
	AddTime        string      `json:"add_time"`
	UpdateTime     string      `json:"update_time"`
	Prices         []struct {
		ID           int    `json:"id"`
		ProductID    int    `json:"product_id"`
		Price        int    `json:"price"`
		Currency     string `json:"currency"`
		Cost         int    `json:"cost"`
		OverheadCost int    `json:"overhead_cost"`
	} `json:"prices"`
}

func (p Product) String() string {
	return Stringify(p)
}

// ProductsResponse represents multiple products response.
type ProductsResponse struct {
	Success        bool           `json:"success"`
	Data           []Product      `json:"data"`
	AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

// ProductResponse represents single product response.
type ProductResponse struct {
	Success        bool           `json:"success"`
	Data           []Product      `json:"data"`
	AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

// ProductAttachedDealsResponse represents attached deals response.
type ProductAttachedDealsResponse struct {
	Success        bool           `json:"success"`
	Data           []Deal         `json:"data"`
	AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

// GetAttachedDeals returns product attached deals.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Products/get_products_id_deals
func (s *ProductsService) GetAttachedDeals(ctx context.Context, id int) (*ProductAttachedDealsResponse, *Response, error) {
	uri := fmt.Sprintf("/products/%v/deals", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *ProductAttachedDealsResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// List returns all data about products.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Products/get_products
func (s *ProductsService) List(ctx context.Context) (*ProductsResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/products", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *ProductsResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// ProductFindOptions specifices the optional parameters to the
// ProductFindOptions.Find method.
type ProductFindOptions struct {
	Term string `url:"term"`
}

// Find products.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Products/get_products_find
func (s *ProductsService) Find(ctx context.Context, term string) (*ProductsResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/products/find", &ProductFindOptions{
		Term: term,
	}, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *ProductsResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// GetByID returns a specific product.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Products/get_products_id
func (s *ProductsService) GetByID(ctx context.Context, id int) (*ProductResponse, *Response, error) {
	uri := fmt.Sprintf("/products/%v", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *ProductResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// ProductCreateOptions specifices the optional parameters to the
// ProductsService.Create method.
type ProductCreateOptions struct {
	Name       string     `url:"name"`
	Code       string     `url:"code"`
	Unit       string     `url:"unit"`
	Tax        int        `url:"tax"`
	ActiveFlag ActiveFlag `url:"active_flag"`
	VisibleTo  VisibleTo  `url:"visible_to"`
	OwnerID    int        `url:"owner_id"`
	Prices     string     `url:"prices"`
}

// Create a new product.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Products/post_products
func (s *ProductsService) Create(ctx context.Context, opt *ProductCreateOptions) (*ProductResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/products", nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *ProductResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// ProductUpdateOptions specifices the optional parameters to the
// ProductsService.Update method.
type ProductUpdateOptions struct {
	Name       string     `url:"name"`
	Code       string     `url:"code"`
	Unit       string     `url:"unit"`
	Tax        int        `url:"tax"`
	ActiveFlag ActiveFlag `url:"active_flag"`
	VisibleTo  VisibleTo  `url:"visible_to"`
	OwnerID    int        `url:"owner_id"`
	Prices     string     `url:"prices"`
}

// Update a specific product.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Products/put_products_id
func (s *ProductsService) Update(ctx context.Context, id int, opt *ProductUpdateOptions) (*ProductResponse, *Response, error) {
	uri := fmt.Sprintf("/products/%v", id)
	req, err := s.client.NewRequest(http.MethodPut, uri, nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *ProductResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Delete a specific product.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Products/delete_products_id
func (s *ProductsService) Delete(ctx context.Context, id int) (*Response, error) {
	uri := fmt.Sprintf("/products/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// DeleteFollower removes follower from a specific product.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Products/delete_products_id_followers_follower_id
func (s *ProductsService) DeleteFollower(ctx context.Context, id int, followerID int) (*Response, error) {
	uri := fmt.Sprintf("/products/%v/followers/%v", id, followerID)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
