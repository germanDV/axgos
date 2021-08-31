package mock

import (
	"fmt"
	"gitlab.com/germanDV/axgos/core"
	"net/http"
)

type Mock struct {
	Method     string
	Url        string
	ReqBody    string
	ResBody    string
	StatusCode int
	Error      error
}

func (m *Mock) GetResponse() (*core.Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	return &core.Response{
		StatusCode: m.StatusCode,
		Status:     fmt.Sprintf("%d %s", m.StatusCode, http.StatusText(m.StatusCode)),
		Body:       []byte(m.ResBody),
	}, nil
}
