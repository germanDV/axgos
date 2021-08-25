package core

import "net/http"

// AxgosHttpClient implements the Do func, which makes it work as a http.Client.
type AxgosHttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}
