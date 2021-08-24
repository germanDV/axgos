package examples

import (
	"errors"
	"gitlab.com/germanDV/axgos/mock"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	mock.MockServer.Start()
	code := m.Run()
	mock.MockServer.Stop()
	os.Exit(code)
}

func TestGet(t *testing.T) {
	t.Run("ErrorTimeout", func(t *testing.T) {
		mock.MockServer.Flush()
		wantErr := "connection timeout"

		mock.MockServer.Add(mock.Mock{
			Method: http.MethodGet,
			Url:    "https://jsonplaceholder.typicode.com/posts/1",
			Error:  errors.New(wantErr),
		})

		res, err := Get(1)
		if res != nil {
			t.Error("Expected no response")
		}
		if err == nil {
			t.Error("Expected an error")
		}
		if err.Error() != wantErr {
			t.Errorf("Expected error %q, got %q\n", wantErr, err.Error())
		}
	})

	t.Run("ErrorCodeReturned", func(t *testing.T) {
		mock.MockServer.Flush()
		wantBody := `{"message":"forbidden resource"}`

		mock.MockServer.Add(mock.Mock{
			Method:     http.MethodGet,
			Url:        "https://jsonplaceholder.typicode.com/posts/9",
			StatusCode: http.StatusForbidden,
			ResBody:    wantBody,
		})

		res, err := Get(9)
		if res != nil {
			t.Error("Expected no response")
		}
		if err == nil {
			t.Error("Expected an error")
		}
		if err.Error() != wantBody {
			t.Errorf("Expected error %q, got %q", wantBody, err.Error())
		}
	})

	t.Run("ErrorUnmarshalJSON", func(t *testing.T) {
		mock.MockServer.Flush()
		wantErr := "cannot unmarshal string into"

		mock.MockServer.Add(mock.Mock{
			Method:     http.MethodGet,
			Url:        "https://jsonplaceholder.typicode.com/posts/9",
			StatusCode: http.StatusOK,
			ResBody:    `{"id": "will cause error because ID expects an int"}`,
		})

		res, err := Get(9)
		if res != nil {
			t.Error("Expected no response")
		}
		if err == nil {
			t.Error("Expected an error")
		}
		if !strings.Contains(err.Error(), wantErr) {
			t.Errorf("Expected error to contain %q, got %q\n", wantErr, err.Error())
		}
	})

	t.Run("SuccessfulRequest", func(t *testing.T) {
		mock.MockServer.Flush()
		mock.MockServer.Add(mock.Mock{
			Method:     http.MethodGet,
			Url:        "https://jsonplaceholder.typicode.com/posts/4",
			StatusCode: http.StatusOK,
			ResBody:    `{"id": 4, "userId": 1, "title": "eum et est occaecati", "body": "etc"}`,
		})

		res, err := Get(4)
		if err != nil {
			t.Error("Expected no error")
		}
		if res == nil {
			t.Error("Expected response not to be empty")
		}
		if res.ID != 4 {
			t.Errorf("Expect post ID 4, got %d\n", res.ID)
		}
		if res.Body != "etc" {
			t.Errorf("Expect post body etc, got %q\n", res.Body)
		}
	})
}
