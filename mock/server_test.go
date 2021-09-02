package mock

import (
	"net/http"
	"testing"
)

var ms mockServer

func TestAdd(t *testing.T) {
	m := Mock{
		Method:     http.MethodGet,
		Url:        "https://jsonplaceholder.typicode.com/posts/9",
		StatusCode: http.StatusForbidden,
		ResBody:    `{"message":"forbidden resource"}`,
	}

	MockServer.Add(m)
	wantedKey := MockServer.createKey(m.Method, m.Url, m.ReqBody)

	if len(MockServer.mocks) != 1 {
		t.Errorf("Expected 1 mock to be present, found %d\n", len(MockServer.mocks))
	}

	got, ok := MockServer.mocks[wantedKey]
	if !ok {
		t.Errorf("Expected mock with key %q to exist but it was not found.\n", wantedKey)
	}
	if got.Url != m.Url {
		t.Errorf("Expected mock to have URL %q, got %q\n", m.Url, got.Url)
	}
}

func TestFlush(t *testing.T) {
	MockServer.Add(Mock{
		Method:     http.MethodGet,
		Url:        "https://jsonplaceholder.typicode.com/posts/4",
		StatusCode: http.StatusOK,
		ResBody:    `{"id": 4, "userId": 1, "title": "eum et est occaecati", "body": "etc"}`,
	})

	MockServer.Add(Mock{
		Method:     http.MethodDelete,
		Url:        "https://jsonplaceholder.typicode.com/posts/4",
		StatusCode: http.StatusOK,
	})

	if len(MockServer.mocks) < 2 {
		t.Errorf("Expected at least 2 mocks to be present, found %d\n", len(MockServer.mocks))
	}

	MockServer.Flush()
	if len(MockServer.mocks) != 0 {
		t.Errorf("Expected all mocks to be flushed, found %d\n", len(MockServer.mocks))
	}
}

func TestEnableDisable(t *testing.T) {
	if MockServer.IsEnabled() {
		t.Error("Expected mock server to be disabled by default.")
	}

	MockServer.Enable()
	if !MockServer.IsEnabled() {
		t.Error("Expected mock server to be enabled.")
	}

	MockServer.Disable()
	if MockServer.IsEnabled() {
		t.Error("Expected mock server to be disabled.")
	}
}
