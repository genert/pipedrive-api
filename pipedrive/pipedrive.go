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
	BaseURL  *url.URL
	apiKey   string

	common service

	Deals *DealService
}

type service struct {
	client *Client
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

func (c *Client) CreateRequestUrl(resource string) string {
	return "https://api.pipedrive.com/v1/deals?api_token=bc5b30cb07ac9572597b427c1767ab650eef03ef"
}

func New(apiKey string) *Client {
	baseURL, _ := url.Parse(defaultBaseUrl)

	c := &Client{
		client: http.DefaultClient,
		BaseURL: baseURL,
		apiKey: apiKey,
	}

	c.common.client = c

	c.Deals = (*DealService)(&c.common)

	return c;
}