package examples

import (
	"gitlab.com/germanDV/axgos/axgos"
	"net/http"
	"time"
)

var client = createClient()

func createClient() axgos.Client {
	headers := make(http.Header)
	headers.Set("Content-Type", "application/json")
	headers.Set("Accept", "application/json")

	return axgos.
		NewBuilder().
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		SetHeaders(headers).
		SetBaseURL("https://jsonplaceholder.typicode.com").
		Build()
}

type BlogPost struct {
	ID     int    `json:"id,omitempty"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}
