package pipedrive

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"
	"net/url"
	"strings"
)

const (
	defaultBaseUrl = "api.pipedrive.com/"

	libraryVersion = "1"

	hostProtocol = "https"
)

type Client struct {
	client   *http.Client // HTTP client used to communicate with the API.

	// Base URL for API requests. Defaults to the public Pipedrive API, but can be
	// set to a domain endpoint to use. BaseURL should
	// always be specified with a trailing slash.
	BaseURL  *url.URL
	apiKey   string

	// Reuse a single struct instead of allocating one for each service.
	common 			service

	Deals 			*DealService

	Currencies 		*CurrenciesService
	NoteFields		*NoteFieldsService
}

type service struct {
	client *Client
}

type Config struct {
	ApiKey 			string
	CompanyDomain 	string
}

func (c *Client) NewRequest(method, url string, body interface {}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	u, err := c.BaseURL.Parse(url)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, u.String(), nil)

	if err != nil {
		return nil, err
	}

	if method == http.MethodGet {
		request.Header.Set("Accept", "application/json")
	}

	return request, nil
}

func (c *Client) Do(request *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(request)

	if err != nil {
		return resp, nil
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)

	if err == io.EOF {
		return resp, nil
	}

	return resp, nil
}

func (c *Client) CreateRequestPayload() string {
	payload := url.Values{};

	payload.Add("api_token", c.apiKey)

	return payload.Encode()
}

func (c *Client) CreateRequestUrl(path string) string {
	uri, err := c.BaseURL.Parse(hostProtocol + "://" + defaultBaseUrl + "v" + libraryVersion)

	if err != nil {
		panic(err)
		return ""
	}

	uri.Path += path

	parameters := url.Values{}

	parameters.Add("api_token", c.apiKey)

	uri.RawQuery = parameters.Encode()

	return uri.String()
}

func New(options *Config) *Client {
	baseURL, _ := url.Parse(defaultBaseUrl)

	c := &Client{
		client: http.DefaultClient,
		BaseURL: baseURL,
		apiKey: options.ApiKey,
	}

	c.common.client = c

	c.Deals = (*DealService)(&c.common)

	c.Currencies = (*CurrenciesService)(&c.common)
	c.NoteFields = (*NoteFieldsService)(&c.common)

	return c
}