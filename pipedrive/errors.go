package pipedrive

import "net/http"

// RateLimitError occurs when Pipedrive returns 403 Forbidden response with a rate limit
// remaining value of 0.
type RateLimitError struct {
	Rate     Rate
	Response *http.Response
	Message  string
}

func (e *RateLimitError) Error() string {
	return "Something went wrong with rate"
}

// ErrorResponse reports one or more errors caused by an API request.
type ErrorResponse struct {
	Response *http.Response
}

func (e *ErrorResponse) Error() string {
	return "Something went wrong with response"
}
