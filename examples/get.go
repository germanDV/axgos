package examples

import (
	"errors"
	"fmt"
	"net/http"
)

func Get(postID int) (*BlogPost, error) {
	url := fmt.Sprintf("/posts/%d", postID)

	// create auth header
	headers := make(http.Header)
	headers.Set("Authorization", "Bearer my-token-abc123")

	// `client` is defined in client.go.
	res, err := client.Get(url, headers)
	if err != nil {
		return nil, err
	}

	// Check that the response status code is < 300.
	if !res.OK() {
		return nil, errors.New(res.String())
	}

	fmt.Println(res.StatusCode)
	fmt.Println(res.String())
	fmt.Println(res.Bytes())

	// Unmarshal JSON response into Go struct.
	var p BlogPost
	err = res.Unmarshal(&p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
