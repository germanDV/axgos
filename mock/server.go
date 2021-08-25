package mock

import (
	"crypto/md5"
	"encoding/hex"
	"gitlab.com/germanDV/axgos/core"
	"strings"
	"sync"
)

type mockServer struct {
	enabled bool
	mu      sync.Mutex
	mocks   map[string]*Mock
	client  core.AxgosHttpClient
}

var MockServer = mockServer{
	mocks:  make(map[string]*Mock),
	client: &clientMock{},
}

func (m *mockServer) createKey(method, url, body string) string {
	hasher := md5.New()
	hasher.Write([]byte(method + url + m.cleanBody(body)))
	key := hex.EncodeToString(hasher.Sum(nil))
	return key
}

func (m *mockServer) cleanBody(body string) string {
	body = strings.TrimSpace(body)
	if body == "" {
		return body
	}

	body = strings.ReplaceAll(body, "\t", "")
	body = strings.ReplaceAll(body, "\n", "")
	return body
}

// GetClient returns the http client stored in the mock struct.
func (m *mockServer) GetClient() core.AxgosHttpClient {
	return m.client
}

// Enable enables mocking responses.
func (m *mockServer) Enable() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.enabled = true
}

// Disable disables mocking responses.
func (m *mockServer) Disable() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.enabled = false
}

// IsEnabled checks if mocked responses are enabled or not.
func (m *mockServer) IsEnabled() bool {
	return m.enabled
}

// Add adds a mock to the mocks map.
// When mocked responses are enabled, the client will return the
// mock response stored in the mocks map instead of making an actual
// http request.
func (m *mockServer) Add(mock Mock) {
	m.mu.Lock()
	defer m.mu.Unlock()
	key := m.createKey(mock.Method, mock.Url, mock.ReqBody)
	m.mocks[key] = &mock
}

// Flush removes all mocks from the mocks map.
func (m *mockServer) Flush() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.mocks = make(map[string]*Mock)
}
