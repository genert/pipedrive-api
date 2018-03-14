package pipedrive

import (
	"fmt"
	"net/http"
)

// RateLimitError occurs when Pipedrive returns 403 Forbidden response with a rate limit
// remaining value of 0.
type RateLimitError struct {
	Rate     Rate
	Response *http.Response
	Message  string
}

func (e *RateLimitError) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		e.Response.Request.Method, e.Response.Request.URL,
		e.Response.StatusCode, e.Message)
}

// ErrorResponse reports one or more errors caused by an API request.
type ErrorResponse struct {
	Response *http.Response
	Message  string
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		e.Response.Request.Method, e.Response.Request.URL,
		e.Response.StatusCode, e.Message)
}
