package gohttp

import (
	"net/http"
	"time"
)

type clientBuilder struct {
	headers                   http.Header
	maxIdleConnectionsPerHost int
	connectionTimeout         time.Duration
	responseTimeout           time.Duration
	disabledTimeouts          bool
	client                    *http.Client
	baseURL                   string
}

type ClientBuilder interface {
	SetHeaders(headers http.Header) ClientBuilder
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetResponseTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConnectionsPerHost(connections int) ClientBuilder
	DisableTimeouts() ClientBuilder
	SetHttpClient(c *http.Client) ClientBuilder
	SetBaseURL(url string) ClientBuilder
	Build() Client
}

func NewBuilder() ClientBuilder {
	return &clientBuilder{}
}

func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}

func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	c.connectionTimeout = timeout
	return c
}

func (c *clientBuilder) SetResponseTimeout(timeout time.Duration) ClientBuilder {
	c.responseTimeout = timeout
	return c
}

func (c *clientBuilder) SetMaxIdleConnectionsPerHost(connections int) ClientBuilder {
	c.maxIdleConnectionsPerHost = connections
	return c
}

func (c *clientBuilder) DisableTimeouts() ClientBuilder {
	c.disabledTimeouts = true
	return c
}

func (c *clientBuilder) SetHttpClient(client *http.Client) ClientBuilder {
	c.client = client
	return c
}

func (c *clientBuilder) SetBaseURL(url string) ClientBuilder {
	c.baseURL = url
	return c
}

func (c *clientBuilder) Build() Client {
	return &httpClient{builder: c}
}
