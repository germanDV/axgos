package axgos

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

type Builder interface {
	SetHeaders(headers http.Header) Builder
	SetConnectionTimeout(timeout time.Duration) Builder
	SetResponseTimeout(timeout time.Duration) Builder
	SetMaxConnectionsPerHost(connections int) Builder
	SetMaxIdleConnectionsPerHost(connections int) Builder
	DisableTimeouts() Builder
	SetHttpClient(c *http.Client) Builder
	SetBaseURL(url string) Builder
	Build() Client
}

// NewBuilder creates a builder that allows to create and configure the axgos client.
func NewBuilder() Builder {
	return &axgosBuilder{}
}

// SetHeaders sets request headers on the client level (will apply to all requests).
func (b *axgosBuilder) SetHeaders(headers http.Header) Builder {
	b.headers = headers
	return b
}

// SetConnectionTimeout sets a timeout for the Dialer.
func (b *axgosBuilder) SetConnectionTimeout(timeout time.Duration) Builder {
	b.connectionTimeout = timeout
	return b
}

// SetResponseTimeout sets the ResponseHeaderTimeout on the Transport.
func (b *axgosBuilder) SetResponseTimeout(timeout time.Duration) Builder {
	b.responseTimeout = timeout
	return b
}

// SetMaxConnectionsPerHost sets a limit of connections per host.
func (b *axgosBuilder) SetMaxConnectionsPerHost(connections int) Builder {
	b.maxConnectionsPerHost = connections
	return b
}

// SetMaxIdleConnectionsPerHost sets a limit of idle connections per host,
// should be <= SetMaxConnectionsPerHost.
func (b *axgosBuilder) SetMaxIdleConnectionsPerHost(connections int) Builder {
	b.maxIdleConnectionsPerHost = connections
	return b
}

// DisableTimeouts allows the client to make requests without any timeouts.
func (b *axgosBuilder) DisableTimeouts() Builder {
	b.disabledTimeouts = true
	return b
}

// SetHttpClient allows a custom client to be provided, otherwise one is created.
func (b *axgosBuilder) SetHttpClient(client *http.Client) Builder {
	b.client = client
	return b
}

// SetBaseURL sets a URL that will be used for every request performed by the same client.
func (b *axgosBuilder) SetBaseURL(url string) Builder {
	b.baseURL = url
	return b
}

// Build creates and returns the client.
func (b *axgosBuilder) Build() Client {
	return &axgosClient{builder: b}
}
