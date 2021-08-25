package gohttp

import (
	"net/http"
	"time"
)

type axgosBuilder struct {
	headers                   http.Header
	maxConnectionsPerHost     int
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
	SetMaxConnectionsPerHost(connections int) AxgosBuilder
	SetMaxIdleConnectionsPerHost(connections int) AxgosBuilder
	DisableTimeouts() AxgosBuilder
	SetHttpClient(c *http.Client) AxgosBuilder
	SetBaseURL(url string) AxgosBuilder
	Build() AxgosClient
}

// NewBuilder creates a builder that allows to create and configure the axgos client.
func NewBuilder() AxgosBuilder {
	return &axgosBuilder{}
}

// SetHeaders sets request headers on the client level (will apply to all requests).
func (b *axgosBuilder) SetHeaders(headers http.Header) AxgosBuilder {
	b.headers = headers
	return b
}

// SetConnectionTimeout sets a timeout for the Dialer.
func (b *axgosBuilder) SetConnectionTimeout(timeout time.Duration) AxgosBuilder {
	b.connectionTimeout = timeout
	return b
}

// SetResponseTimeout sets the ResponseHeaderTimeout on the Transport.
func (b *axgosBuilder) SetResponseTimeout(timeout time.Duration) AxgosBuilder {
	b.responseTimeout = timeout
	return b
}

// SetMaxConnectionsPerHost sets a limit of connections per host.
func (b *axgosBuilder) SetMaxConnectionsPerHost(connections int) AxgosBuilder {
	b.maxConnectionsPerHost = connections
	return b
}

// SetMaxIdleConnectionsPerHost sets a limit of idle connections per host,
// should be <= SetMaxConnectionsPerHost.
func (b *axgosBuilder) SetMaxIdleConnectionsPerHost(connections int) AxgosBuilder {
	b.maxIdleConnectionsPerHost = connections
	return b
}

// DisableTimeouts allows the client to make requests without any timeouts.
func (b *axgosBuilder) DisableTimeouts() AxgosBuilder {
	b.disabledTimeouts = true
	return b
}

// SetHttpClient allows a custom client to be provided, otherwise one is created.
func (b *axgosBuilder) SetHttpClient(client *http.Client) AxgosBuilder {
	b.client = client
	return b
}

// SetBaseURL sets a URL that will be used for every request performed by the same client.
func (b *axgosBuilder) SetBaseURL(url string) AxgosBuilder {
	b.baseURL = url
	return b
}

// Build creates and returns the client.
func (b *axgosBuilder) Build() AxgosClient {
	return &axgosClient{builder: b}
}
