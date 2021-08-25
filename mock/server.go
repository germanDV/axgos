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

func (m *mockServer) GetClient() core.AxgosHttpClient {
	return m.client
}

func (m *mockServer) Start() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.enabled = true
}

func (m *mockServer) Stop() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.enabled = false
}

func (m *mockServer) IsEnabled() bool {
	return m.enabled
}

func (m *mockServer) Add(mock Mock) {
	m.mu.Lock()
	defer m.mu.Unlock()
	key := m.createKey(mock.Method, mock.Url, mock.ReqBody)
	m.mocks[key] = &mock
}

func (m *mockServer) Flush() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.mocks = make(map[string]*Mock)
}
