package gohttp

import (
	"net/http"
	"time"
)

type axgosBuilder struct {
	headers                   http.Header
	maxIdleConnectionsPerHost int
	connectionTimeout         time.Duration
	responseTimeout           time.Duration
	disabledTimeouts          bool
	client                    *http.Client
	baseURL                   string
}

type AxgosBuilder interface {
	SetHeaders(headers http.Header) AxgosBuilder
	SetConnectionTimeout(timeout time.Duration) AxgosBuilder
	SetResponseTimeout(timeout time.Duration) AxgosBuilder
	SetMaxIdleConnectionsPerHost(connections int) AxgosBuilder
	DisableTimeouts() AxgosBuilder
	SetHttpClient(c *http.Client) AxgosBuilder
	SetBaseURL(url string) AxgosBuilder
	Build() AxgosClient
}

func NewBuilder() AxgosBuilder {
	return &axgosBuilder{}
}

func (b *axgosBuilder) SetHeaders(headers http.Header) AxgosBuilder {
	b.headers = headers
	return b
}

func (b *axgosBuilder) SetConnectionTimeout(timeout time.Duration) AxgosBuilder {
	b.connectionTimeout = timeout
	return b
}

func (b *axgosBuilder) SetResponseTimeout(timeout time.Duration) AxgosBuilder {
	b.responseTimeout = timeout
	return b
}

func (b *axgosBuilder) SetMaxIdleConnectionsPerHost(connections int) AxgosBuilder {
	b.maxIdleConnectionsPerHost = connections
	return b
}

func (b *axgosBuilder) DisableTimeouts() AxgosBuilder {
	b.disabledTimeouts = true
	return b
}

func (b *axgosBuilder) SetHttpClient(client *http.Client) AxgosBuilder {
	b.client = client
	return b
}

func (b *axgosBuilder) SetBaseURL(url string) AxgosBuilder {
	b.baseURL = url
	return b
}

func (b *axgosBuilder) Build() AxgosClient {
	return &axgosClient{builder: b}
}
