package gohttp

import (
	"gitlab.com/germanDV/axgos/core"
	"net/http"
	"sync"
)

type AxgosClient interface {
	Get(url string, headers ...http.Header) (*core.AxgosResponse, error)
	Post(url string, body interface{}, headers ...http.Header) (*core.AxgosResponse, error)
	Put(url string, body interface{}, headers ...http.Header) (*core.AxgosResponse, error)
	Patch(url string, body interface{}, headers ...http.Header) (*core.AxgosResponse, error)
	Delete(url string, headers ...http.Header) (*core.AxgosResponse, error)
	Options(url string, headers ...http.Header) (*core.AxgosResponse, error)
}

type axgosClient struct {
	builder *axgosBuilder
	once    sync.Once
	client  *http.Client
}

// Get performs a GET request.
func (c *axgosClient) Get(url string, headers ...http.Header) (*core.AxgosResponse, error) {
	return c.do(http.MethodGet, url, getHeaders(headers...), nil)
}

// Post performs a POST request.
func (c *axgosClient) Post(url string, body interface{}, headers ...http.Header) (*core.AxgosResponse, error) {
	return c.do(http.MethodPost, url, getHeaders(headers...), body)
}

// Put performs a PUT request.
func (c *axgosClient) Put(url string, body interface{}, headers ...http.Header) (*core.AxgosResponse, error) {
	return c.do(http.MethodPut, url, getHeaders(headers...), body)
}

// Patch performs a PATCH request.
func (c *axgosClient) Patch(url string, body interface{}, headers ...http.Header) (*core.AxgosResponse, error) {
	return c.do(http.MethodPatch, url, getHeaders(headers...), body)
}

// Delete performs a DELETE request.
func (c *axgosClient) Delete(url string, headers ...http.Header) (*core.AxgosResponse, error) {
	return c.do(http.MethodDelete, url, getHeaders(headers...), nil)
}

// Options performs a OPTIONS request.
func (c *axgosClient) Options(url string, headers ...http.Header) (*core.AxgosResponse, error) {
	return c.do(http.MethodOptions, url, getHeaders(headers...), nil)
}
