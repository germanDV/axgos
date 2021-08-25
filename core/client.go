package core

import "net/http"

type AxgosHttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}
