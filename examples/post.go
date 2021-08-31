package examples

import "errors"

func Post(p BlogPost) (*BlogPost, error) {
	// `client` is defined in client.go
	res, err := client.Post("/posts", p)
	if err != nil {
		return nil, err
	}

	// Check that the response status code is < 300.
	if !res.OK() {
		return nil, errors.New(res.String())
	}

	// Unmarshal JSON response into Go struct.
	var insertedPost BlogPost
	err = res.Unmarshal(&insertedPost)
	if err != nil {
		return nil, err
	}

	return &insertedPost, nil
}
