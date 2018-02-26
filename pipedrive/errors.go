package pipedrive

import "net/http"

type RateLimitError struct {
	Rate     Rate
	Response *http.Response
	Message string
}

func (e *RateLimitError) Error() string {
	return "Something went wrong with rate"
}

type ErrorResponse struct {
	Response *http.Response
}

func (e *ErrorResponse) Error() string {
	return "Something went wrong with response"
}
