package core

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"strings"
)

// Response represents the response returned by http requests.
type Response struct {
	StatusCode int
	Status     string
	ReqHeaders http.Header
	Headers    http.Header
	Body       []byte
}

// OK checks that the response status code is successful (< 300).
func (r *Response) OK() bool {
	return r.StatusCode < 300
}

// Bytes returns the response body as a []byte.
func (r *Response) Bytes() []byte {
	return r.Body
}

// Bytes returns the response body as a string.
func (r *Response) String() string {
	return string(r.Body)
}

// Unmarshal unmarshalls the response body into a Go struct.
// The format will depend on the `Accept` request header, defaulting to JSON.
func (r *Response) Unmarshal(target interface{}) error {
	accepts := r.ReqHeaders.Get("Accept")
	accepts = strings.ToLower(accepts)

	if strings.Contains(accepts, "application/json") {
		return json.Unmarshal(r.Body, target)
	}

	if strings.Contains(accepts, "application/xml") {
		return xml.Unmarshal(r.Body, target)
	}

	// Use JSON as default
	return json.Unmarshal(r.Body, target)
}
