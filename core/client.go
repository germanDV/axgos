package core

import "net/http"

type AxgosClient interface {
	Do(req *http.Request) (*http.Response, error)
}
