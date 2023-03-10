package axgos

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"gitlab.com/germanDV/axgos/core"
	"gitlab.com/germanDV/axgos/mock"
)

const (
	defaultMaxConnectionsPerHost     = 5
	defaultMaxIdleConnectionsPerHost = 5
	defaultResponseTimeout           = 5 * time.Second
	defaultConnectionTimeout         = 2 * time.Second
)

func (c *axgosClient) do(method string, url string, headers http.Header, body interface{}) (*core.Response, error) {
	// Combine common and request-specific headers
	fullHeaders := c.getHeaders(headers)

	// Convert body based on Content-Type header
	reqBody, err := c.getReqBody(fullHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, err
	}

	// Add base URL, if provided when building the client
	if baseURL := c.getBaseURL(); baseURL != "" {
		url = baseURL + url
	}

	// Create request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	// Add headers
	req.Header = fullHeaders

	// Get a client (if nil, create one -> it's created on the first request)
	client := c.getAxgosClient()

	// Make request
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := core.Response{
		StatusCode: res.StatusCode,
		Status:     res.Status,
		ReqHeaders: fullHeaders,
		Headers:    res.Header,
		Body:       respBody,
	}

	return &response, nil
}

// getAxgosClient returns the existing client or creates it if it doesn't exist.
// If mocks are enabled, it returns a mock client.
func (c *axgosClient) getAxgosClient() core.AxgosHttpClient {
	if mock.MockServer.IsEnabled() {
		return mock.MockServer.GetClient()
	}

	c.once.Do(func() {
		if c.builder.client != nil {
			// Return the client set up by the user.
			c.client = c.builder.client
			return
		}

		// User has not set up a client, create one.
		c.client = &http.Client{
			Timeout: c.getConnectionTimeout() + c.getResponseTimeout(),
			Transport: &http.Transport{
				MaxConnsPerHost:       c.getMaxConnectionsPerHost(),
				MaxIdleConnsPerHost:   c.getMaxIdleConnectionsPerHost(),
				ResponseHeaderTimeout: c.getResponseTimeout(),
				DialContext:           (&net.Dialer{Timeout: c.getConnectionTimeout()}).DialContext,
			},
		}
	})

	return c.client
}

func (c *axgosClient) getMaxConnectionsPerHost() int {
	if c.builder.maxConnectionsPerHost > 0 {
		return c.builder.maxConnectionsPerHost
	}
	return defaultMaxConnectionsPerHost
}

func (c *axgosClient) getMaxIdleConnectionsPerHost() int {
	if c.builder.maxIdleConnectionsPerHost > 0 {
		return c.builder.maxIdleConnectionsPerHost
	}
	return defaultMaxIdleConnectionsPerHost
}

func (c *axgosClient) getResponseTimeout() time.Duration {
	if c.builder.disabledTimeouts {
		return 0
	}
	if c.builder.responseTimeout > 0 {
		return c.builder.responseTimeout
	}
	return defaultResponseTimeout
}

func (c *axgosClient) getConnectionTimeout() time.Duration {
	if c.builder.disabledTimeouts {
		return 0
	}
	if c.builder.connectionTimeout > 0 {
		return c.builder.connectionTimeout
	}
	return defaultConnectionTimeout
}

func (c *axgosClient) getReqBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch strings.ToLower(contentType) {
	case "application/xml":
		return xml.Marshal(body)
	default:
		return json.Marshal(body)
	}
}

func (c *axgosClient) getBaseURL() string {
	return c.builder.baseURL
}
