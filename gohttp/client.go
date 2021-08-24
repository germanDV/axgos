package gohttp

import (
	"gitlab.com/germanDV/axgos/core"
	"net/http"
	"sync"
)

type Client interface {
	Get(url string, headers ...http.Header) (*core.AxgosResponse, error)
	Post(url string, body interface{}, headers ...http.Header) (*core.AxgosResponse, error)
	Put(url string, body interface{}, headers ...http.Header) (*core.AxgosResponse, error)
	Patch(url string, body interface{}, headers ...http.Header) (*core.AxgosResponse, error)
	Delete(url string, headers ...http.Header) (*core.AxgosResponse, error)
	Options(url string, headers ...http.Header) (*core.AxgosResponse, error)
}

type httpClient struct {
	builder *clientBuilder
	once    sync.Once
	client  *http.Client
}

func (c *httpClient) Get(url string, headers ...http.Header) (*core.AxgosResponse, error) {
	return c.do(http.MethodGet, url, getHeaders(headers...), nil)
}

func (c *httpClient) Post(url string, body interface{}, headers ...http.Header) (*core.AxgosResponse, error) {
	return c.do(http.MethodPost, url, getHeaders(headers...), body)
}

func (c *httpClient) Put(url string, body interface{}, headers ...http.Header) (*core.AxgosResponse, error) {
	return c.do(http.MethodPut, url, getHeaders(headers...), body)
}

func (c *httpClient) Patch(url string, body interface{}, headers ...http.Header) (*core.AxgosResponse, error) {
	return c.do(http.MethodPatch, url, getHeaders(headers...), body)
}

func (c *httpClient) Delete(url string, headers ...http.Header) (*core.AxgosResponse, error) {
	return c.do(http.MethodDelete, url, getHeaders(headers...), nil)
}

func (c *httpClient) Options(url string, headers ...http.Header) (*core.AxgosResponse, error) {
	return c.do(http.MethodOptions, url, getHeaders(headers...), nil)
}
