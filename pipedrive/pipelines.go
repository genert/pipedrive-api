package pipedrive

import (
	"context"
	"fmt"
	"net/http"
)

// PipelinesService handles pipelines related
// methods of the Pipedrive API.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Pipelines
type PipelinesService service

// Pipeline represents a Pipedrive pipeline.
type Pipeline struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	URLTitle        string `json:"url_title"`
	OrderNr         int    `json:"order_nr"`
	Active          bool   `json:"active"`
	DealProbability bool   `json:"deal_probability"`
	AddTime         string `json:"add_time"`
	UpdateTime      string `json:"update_time"`
	Selected        bool   `json:"selected"`
}

func (p Pipeline) String() string {
	return Stringify(p)
}

// PipelinesResponse represents multiple pipelines response.
type PipelinesResponse struct {
	Success        bool           `json:"success"`
	Data           []Pipeline     `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

// PipelineResponse represents single pipeline response.
type PipelineResponse struct {
	Success        bool           `json:"success"`
	Data           Pipeline       `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

// PipelineDealsConversionRateResponse represents conversion response.
type PipelineDealsConversionRateResponse struct {
	Success bool `json:"success"`
	Data    struct {
		StageConversions []struct {
			FromStageID    int     `json:"from_stage_id"`
			ToStageID      int     `json:"to_stage_id"`
			ConversionRate float64 `json:"conversion_rate"`
		} `json:"stage_conversions"`
		WonConversion  float64 `json:"won_conversion"`
		LostConversion float64 `json:"lost_conversion"`
	} `json:"data"`
}

// PipelineDealsMovementResponse represents movement response.
type PipelineDealsMovementResponse struct {
	Success bool `json:"success"`
	Data    struct {
		MovementsBetweenStages struct {
			Count int `json:"count"`
		} `json:"movements_between_stages"`
		NewDeals struct {
			Count   int   `json:"count"`
			DealIds []int `json:"deal_ids"`
			Values  struct {
				EUR float64 `json:"EUR"`
			} `json:"values"`
			FormattedValues struct {
				EUR string `json:"EUR"`
			} `json:"formatted_values"`
		} `json:"new_deals"`
		DealsLeftOpen struct {
			Count   int   `json:"count"`
			DealIds []int `json:"deal_ids"`
			Values  struct {
				EUR float64 `json:"EUR"`
			} `json:"values"`
			FormattedValues struct {
				EUR string `json:"EUR"`
			} `json:"formatted_values"`
		} `json:"deals_left_open"`
		WonDeals struct {
			Count   int   `json:"count"`
			DealIds []int `json:"deal_ids"`
			Values  struct {
				EUR int `json:"EUR"`
			} `json:"values"`
			FormattedValues struct {
				EUR string `json:"EUR"`
			} `json:"formatted_values"`
		} `json:"won_deals"`
		LostDeals struct {
			Count   int   `json:"count"`
			DealIds []int `json:"deal_ids"`
			Values  struct {
				EUR int `json:"EUR"`
			} `json:"values"`
			FormattedValues struct {
				EUR string `json:"EUR"`
			} `json:"formatted_values"`
		} `json:"lost_deals"`
		AverageAgeInDays struct {
			AcrossAllStages float64 `json:"across_all_stages"`
			ByStages        []struct {
				StageID int `json:"stage_id"`
				Value   int `json:"value"`
			} `json:"by_stages"`
		} `json:"average_age_in_days"`
	} `json:"data"`
}

// List returns data about all pipelines.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Pipelines/get_pipelines
func (s *PipelinesService) List(ctx context.Context) (*PipelinesResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/pipelines", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *PipelinesResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// GetByID returns data about a specific pipeline.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Pipelines/get_pipelines_id
func (s *PipelinesService) GetByID(ctx context.Context, id int) (*PipelineResponse, *Response, error) {
	uri := fmt.Sprintf("/pipelines/%v", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *PipelineResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// GetDeals returns deal in a specific pipeline.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Pipelines/get_pipelines_id_deals
func (s *PipelinesService) GetDeals(ctx context.Context, id int) (*PipelinesResponse, *Response, error) {
	uri := fmt.Sprintf("/pipelines/%v/deals", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *PipelinesResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// GetDealsConversionRate returns deals conversion rate in a specific pipeline.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Pipelines/get_pipelines_id_conversion_statistics
func (s *PipelinesService) GetDealsConversionRate(ctx context.Context, id int, startDate Timestamp, endDate Timestamp) (*PipelineDealsConversionRateResponse, *Response, error) {
	uri := fmt.Sprintf("/pipelines/%v/conversion_statistics", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, struct {
		StartDate string `url:"start_date"`
		EndDate   string `url:"end_date"`
	}{
		startDate.Format(),
		endDate.Format(),
	}, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *PipelineDealsConversionRateResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// GetDealsMovement returns deals movement in a specific pipeline.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Pipelines/get_pipelines_id_movement_statistics
func (s *PipelinesService) GetDealsMovement(ctx context.Context, id int, startDate Timestamp, endDate Timestamp) (*PipelineDealsMovementResponse, *Response, error) {
	uri := fmt.Sprintf("/pipelines/%v/movement_statistics", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, struct {
		StartDate string `url:"start_date"`
		EndDate   string `url:"end_date"`
	}{
		startDate.Format(),
		endDate.Format(),
	}, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *PipelineDealsMovementResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// PipelineCreateOptions specifices the optional parameters to the
// PipelineCreateOptions.Create method.
type PipelineCreateOptions struct {
	Name            string          `url:"name"`
	DealProbability DealProbability `url:"deal_probability"`
	OrderNr         int             `url:"order_nr"`
	Active          ActiveFlag      `url:"active"`
}

// Create a new pipeline.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Pipelines/post_pipelines
func (s *PipelinesService) Create(ctx context.Context, opt *PipelineCreateOptions) (*PipelineResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/pipelines", nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *PipelineResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// PipelineUpdateOptions specifices the optional parameters to the
// PipelinesService.Update method.
type PipelineUpdateOptions struct {
	Name            string          `url:"name"`
	DealProbability DealProbability `url:"deal_probability"`
	OrderNr         int             `url:"order_nr"`
	Active          ActiveFlag      `url:"active"`
}

// Update a specific pipeline.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Pipelines/put_pipelines_id
func (s *PipelinesService) Update(ctx context.Context, id int, opt *PipelineUpdateOptions) (*PipelineResponse, *Response, error) {
	uri := fmt.Sprintf("/pipelines/%v", id)
	req, err := s.client.NewRequest(http.MethodPost, uri, nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *PipelineResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Delete marks a specific pipeline as deleted.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Pipelines/delete_pipelines_id
func (s *PipelinesService) Delete(ctx context.Context, id int) (*Response, error) {
	uri := fmt.Sprintf("/pipelines/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
