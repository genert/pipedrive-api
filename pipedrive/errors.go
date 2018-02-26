package pipedrive

import "net/http"

type RateLimitError struct {
	Rate     Rate
	Response *http.Response
}

func (e *RateLimitError) Error() string {
	return "Something went wrong"
}

type ErrorResponse struct {
	Response *http.Response
}

func (e *ErrorResponse) Error() string {
	return "Something went wrong"
}
