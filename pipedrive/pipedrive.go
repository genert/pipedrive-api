package pipedrive

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
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
	client *http.Client // HTTP client used to communicate with the API.

	// Base URL for API requests. Defaults to the public Pipedrive API, but can be
	// set to a domain endpoint to use. BaseURL should
	// always be specified with a trailing slash.
	BaseURL *url.URL
	apiKey  string

	// Reuse a single struct instead of allocating one for each service.
	common service

	Deals          *DealService
	Currencies     *CurrenciesService
	NoteFields     *NoteFieldsService
	Notes          *NotesService
	Recents        *RecentsService
	SearchResults  *SearchResultsService
	Users          *UsersService
	Filters        *FiltersService
	ActivityFields *ActivityFieldsService
	ActivityTypes  *ActivityTypesService
	Authorizations *AuthorizationsService
}

type service struct {
	client *Client
}

type Config struct {
	ApiKey        string
	CompanyDomain string
}

type Rate struct {
	Limit     int       `json:"limit"`
	Remaining int       `json:"remaining"`
	Reset     Timestamp `json:"reset"`
}

type Response struct {
	*http.Response
}

type Timestamp time.Time

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
		err := json.NewEncoder(buf).Encode(body)

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

func (c *Client) Do(request *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(request)

	if err != nil {
		return &Response{
			Response: resp,
		}, err
	}

	defer resp.Body.Close()

	response := newResponse(resp)

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

func New(options *Config) *Client {
	baseURL, _ := url.Parse(defaultBaseUrl)

	c := &Client{
		client:  http.DefaultClient,
		BaseURL: baseURL,
		apiKey:  options.ApiKey,
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
	c.ActivityFields = (*ActivityFieldsService)(&c.common)
	c.ActivityTypes = (*ActivityTypesService)(&c.common)
	c.Authorizations = (*AuthorizationsService)(&c.common)

	return c
}

func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
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
