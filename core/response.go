package core

import (
	"encoding/json"
	"net/http"
)

// AxgosResponse represents the response returned by http requests.
type AxgosResponse struct {
	StatusCode int
	Status     string
	Headers    http.Header
	Body       []byte
}

// OK checks that the response status code is successful (< 300).
func (r *AxgosResponse) OK() bool {
	return r.StatusCode < 300
}

// Bytes returns the response body as a []byte.
func (r *AxgosResponse) Bytes() []byte {
	return r.Body
}

// Bytes returns the response body as a string.
func (r *AxgosResponse) String() string {
	return string(r.Body)
}

// UnmarshalJson unmarshalls the response body into a Go struct.
func (r *AxgosResponse) UnmarshalJson(target interface{}) error {
	return json.Unmarshal(r.Body, target)
}
