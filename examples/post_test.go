package examples

import (
	"errors"
	"gitlab.com/germanDV/axgos/mock"
	"net/http"
	"strings"
	"testing"
)

func TestPost(t *testing.T) {
	t.Run("ErrorTimeout", func(t *testing.T) {
		mock.MockServer.Flush()
		wantErr := "connection timeout"

		mock.MockServer.Add(mock.Mock{
			Method:  http.MethodPost,
			Url:     "https://jsonplaceholder.typicode.com/posts",
			ReqBody: `{"title":"Lisa","body":"Simpson","userId":97}`,
			Error:   errors.New(wantErr),
		})

		res, err := Post(BlogPost{Title: "Lisa", Body: "Simpson", UserID: 97})
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
		wantBody := `{"message":"you don't have permission to perform this action"}`

		mock.MockServer.Add(mock.Mock{
			Method:     http.MethodPost,
			Url:        "https://jsonplaceholder.typicode.com/posts",
			ReqBody:    `{"title":"Lisa","body":"Simpson","userId":95}`,
			StatusCode: http.StatusUnauthorized,
			ResBody:    wantBody,
		})

		res, err := Post(BlogPost{Title: "Lisa", Body: "Simpson", UserID: 95})
		if res != nil {
			t.Error("Expected no response")
		}
		if err == nil {
			t.Error("Expected an error")
		}
		if err.Error() != wantBody {
			t.Errorf("Expected error %q, got %q\n", wantBody, err.Error())
		}
	})

	t.Run("ErrorUnmarshalJsonResponse", func(t *testing.T) {
		mock.MockServer.Flush()
		wantErr := "cannot unmarshal string into"

		mock.MockServer.Add(mock.Mock{
			Method:     http.MethodPost,
			Url:        "https://jsonplaceholder.typicode.com/posts",
			ReqBody:    `{"title":"Lisa","body":"Simpson","userId":98}`,
			StatusCode: http.StatusOK,
			ResBody:    `{"title": "Lisa","body":"Simpson","userId":"bad","id":4}`,
		})

		res, err := Post(BlogPost{Title: "Lisa", Body: "Simpson", UserID: 98})
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
			Method:     http.MethodPost,
			Url:        "https://jsonplaceholder.typicode.com/posts",
			ReqBody:    `{"title":"Lisa","body":"Simpson","userId":99}`,
			StatusCode: http.StatusOK,
			ResBody:    `{"title": "Lisa","body":"Simpson","userId":99,"id":4}`,
		})

		res, err := Post(BlogPost{Title: "Lisa", Body: "Simpson", UserID: 99})
		if err != nil {
			t.Error("Expected no error")
		}
		if res == nil {
			t.Error("Expected response not to be empty")
		}
		if res.Title != "Lisa" {
			t.Errorf("Expected response to be a post with title Lisa, got %q", res.Title)
		}
		if res.ID != 4 {
			t.Errorf("Expected response to be a post with id 4, got %d", res.ID)
		}
	})
}
