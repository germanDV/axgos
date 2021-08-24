package examples

import (
	"errors"
	"fmt"
)

func Get(postID int) (*BlogPost, error) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", postID)

	// `axgos` is defined in client.go.
	// headers can be provided to Get as a second argument (type http.Header).
	res, err := axgos.Get(url)
	if err != nil {
		return nil, err
	}

	// Check that the response status code is < 300.
	if !res.OK() {
		return nil, errors.New(res.String())
	}

	// Unmarshal JSON response into Go struct.
	var p BlogPost
	err = res.UnmarshalJson(&p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
