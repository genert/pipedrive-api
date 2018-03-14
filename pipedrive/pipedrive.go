package pipedrive

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	defaultBaseUrl = "api.pipedrive.com/"

	libraryVersion = "1"

	hostProtocol = "https"

	// The amount of requests current API token can perform for the 10 seconds window.
	headerRateLimit = "X-RateLimit-Limit"

	// The amount of requests left for the 10 seconds window.
	headerRateRemaining = "X-RateLimit-Remaining"

	// The amount of seconds before the limit resets.
	headerRateReset = "X-RateLimit-Reset"
)

type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Base URL for API requests. Defaults to the public Pipedrive API, but can be
	// set to a domain endpoint to use. BaseURL should
	// always be specified with a trailing slash.
	BaseURL *url.URL
	apiKey  string

	rateMutex   sync.Mutex
	currentRate Rate

	// Reuse a single struct instead of allocating one for each service.
	common service

	Deals             *DealService
	Currencies        *CurrenciesService
	NoteFields        *NoteFieldsService
	Notes             *NotesService
	Recents           *RecentsService
	SearchResults     *SearchResultsService
	Users             *UsersService
	Filters           *FiltersService
	Activities        *ActivitiesService
	ActivityFields    *ActivityFieldsService
	ActivityTypes     *ActivityTypesService
	Authorizations    *AuthorizationsService
	Stages            *StagesService
	Webhooks          *WebhooksService
	UserConnections   *UserConnectionsService
	GoalsService      *GoalsService
	PipelinesService  *PipelinesService
	UserSettings      *UserSettingsService
	Files             *FilesService
	ProductFields     *ProductFieldsService
	Products          *ProductsService
	PersonFields      *PersonFieldsService
	OrganizationField *OrganizationFieldsService
	DealFields        *DealFieldsService
	Persons           *PersonsService
	Organizations     *OrganizationsService
}

type service struct {
	client *Client
}

type Config struct {
	APIKey        string
	CompanyDomain string
}

type Rate struct {
	Limit     int       `json:"limit"`
	Remaining int       `json:"remaining"`
	Reset     Timestamp `json:"reset"`
}

func (r Rate) String() string {
	return Stringify(r)
}

type Response struct {
	*http.Response
	Rate
}

// Parse the rate from response headers.
func parseRateFromResponse(r *http.Response) Rate {
	var rate Rate

	if limit := r.Header.Get(headerRateLimit); limit != "" {
		rate.Limit, _ = strconv.Atoi(limit)
	}

	if remaining := r.Header.Get(headerRateRemaining); remaining != "" {
		rate.Remaining, _ = strconv.Atoi(remaining)
	}

	if reset := r.Header.Get(headerRateReset); reset != "" {
		if value, _ := strconv.ParseInt(reset, 10, 64); value != 0 {
			rate.Reset = Timestamp{time.Unix(value, 0)}
		}
	}

	return rate
}

func (c *Client) NewRequest(method, url string, opt interface{}, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	u, err := c.createRequestUrl(url, opt)

	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter

	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)

		if err != nil {
			return nil, err
		}
	}

	request, err := http.NewRequest(method, u, buf)

	if err != nil {
		return nil, err
	}

	if body != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	return request, nil
}

func (c *Client) checkRateLimitBeforeDo(req *http.Request) *RateLimitError {
	c.rateMutex.Lock()
	rate := c.currentRate
	c.rateMutex.Unlock()

	if !rate.Reset.Time.IsZero() && rate.Remaining == 0 {
		resp := &http.Response{
			Status:     http.StatusText(http.StatusForbidden),
			StatusCode: http.StatusForbidden,
			Request:    req,
			Header:     make(http.Header, 0),
			Body:       ioutil.NopCloser(bytes.NewBufferString("")),
		}

		return &RateLimitError{
			Rate:     rate,
			Response: resp,
			Message:  fmt.Sprintf("API rate limit of %v exceeded.", rate.Limit),
		}
	}

	return nil
}

func (c *Client) checkResponse(r *http.Response) error {
	if code := r.StatusCode; 200 <= code && code <= 299 {
		return nil
	}

	data, err := ioutil.ReadAll(r.Body)
	errorResponse := &ErrorResponse{Response: r}

	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}

	switch {
	case r.StatusCode == http.StatusForbidden && r.Header.Get(headerRateRemaining) == "0":
		return &RateLimitError{
			Rate:     parseRateFromResponse(r),
			Response: errorResponse.Response,
		}

	default:
		return errorResponse
	}
}

// Do sends an API request and returns the API response. T
//
// The provided ctx must be non-nil. If it is canceled or times out,
// ctx.Err() will be returned.
func (c *Client) Do(ctx context.Context, request *http.Request, v interface{}) (*Response, error) {
	if err := c.checkRateLimitBeforeDo(request); err != nil {
		return &Response{
			Response: err.Response,
		}, err
	}

	resp, err := c.client.Do(request)

	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}

	defer func() {
		io.CopyN(ioutil.Discard, resp.Body, 512)
		resp.Body.Close()
	}()

	response := newResponse(resp)

	c.rateMutex.Lock()
	c.currentRate = response.Rate
	c.rateMutex.Unlock()

	err = c.checkResponse(response.Response)

	if err != nil {
		return response, err
	}

	err = json.NewDecoder(resp.Body).Decode(v)

	if err == io.EOF {
		return response, nil
	}

	return response, nil
}

func (c *Client) createRequestUrl(path string, opt interface{}) (string, error) {
	uri, err := c.BaseURL.Parse(hostProtocol + "://" + defaultBaseUrl + "v" + libraryVersion)

	if err != nil {
		return path, err
	}

	uri.Path += path

	v := reflect.ValueOf(opt)

	if v.Kind() == reflect.Ptr && v.IsNil() {
		parameters := url.Values{}
		parameters.Add("api_token", c.apiKey)

		uri.RawQuery = parameters.Encode()

		return uri.String(), nil
	}

	qs, err := query.Values(opt)

	if err != nil {
		return path, err
	}

	qs.Add("api_token", c.apiKey)

	uri.RawQuery = qs.Encode()

	return uri.String(), nil
}

func (c *Client) SetOptions(options ...func(*Client) error) error {
	for _, opt := range options {
		if err := opt(c); err != nil {
			return err
		}
	}

	return nil
}

func NewClient(options *Config) *Client {
	baseURL, _ := url.Parse(defaultBaseUrl)

	c := &Client{
		client:  http.DefaultClient,
		BaseURL: baseURL,
		apiKey:  options.APIKey,
	}

	c.common.client = c

	c.Deals = (*DealService)(&c.common)
	c.Currencies = (*CurrenciesService)(&c.common)
	c.NoteFields = (*NoteFieldsService)(&c.common)
	c.Notes = (*NotesService)(&c.common)
	c.Recents = (*RecentsService)(&c.common)
	c.SearchResults = (*SearchResultsService)(&c.common)
	c.Users = (*UsersService)(&c.common)
	c.Filters = (*FiltersService)(&c.common)
	c.Activities = (*ActivitiesService)(&c.common)
	c.ActivityFields = (*ActivityFieldsService)(&c.common)
	c.ActivityTypes = (*ActivityTypesService)(&c.common)
	c.Authorizations = (*AuthorizationsService)(&c.common)
	c.Stages = (*StagesService)(&c.common)
	c.Webhooks = (*WebhooksService)(&c.common)
	c.UserConnections = (*UserConnectionsService)(&c.common)
	c.GoalsService = (*GoalsService)(&c.common)
	c.PipelinesService = (*PipelinesService)(&c.common)
	c.UserSettings = (*UserSettingsService)(&c.common)
	c.Files = (*FilesService)(&c.common)
	c.ProductFields = (*ProductFieldsService)(&c.common)
	c.Products = (*ProductsService)(&c.common)
	c.PersonFields = (*PersonFieldsService)(&c.common)
	c.OrganizationField = (*OrganizationFieldsService)(&c.common)
	c.DealFields = (*DealFieldsService)(&c.common)
	c.Persons = (*PersonsService)(&c.common)
	c.Organizations = (*OrganizationsService)(&c.common)

	return c
}

func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	response.Rate = parseRateFromResponse(r)

	return response
}

// Converts []int to string by separator.
func arrayToString(a []int, sep string) string {
	b := make([]string, len(a))

	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}

	return strings.Join(b, sep)
}
