package mock

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type clientMock struct{}

func (c *clientMock) Do(req *http.Request) (*http.Response, error) {
	body, err := req.GetBody()
	if err != nil {
		return nil, err
	}

	defer body.Close()
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(body)
	if err != nil {
		return nil, err
	}

	var res http.Response
	key := MockServer.createKey(req.Method, req.URL.String(), buf.String())
	mock, _ := MockServer.mocks[key]

	if mock != nil {
		if mock.Error != nil {
			return nil, mock.Error
		}
		res.StatusCode = mock.StatusCode
		res.Body = io.NopCloser(strings.NewReader(mock.ResBody))
		res.ContentLength = int64(len(mock.ResBody))
		res.Request = req
		return &res, nil
	}

	errMsg := fmt.Sprintf("No mock matching %s %q", req.Method, req.URL.String())
	return nil, errors.New(errMsg)
}
