package axgos

import "net/http"

func getHeaders(headers ...http.Header) http.Header {
	// If any headers are present for the current request, take the first ones.
	if len(headers) > 0 {
		return headers[0]
	}
	return http.Header{}
}

func (c *axgosClient) getHeaders(customHeaders http.Header) http.Header {
	ret := make(http.Header)

	// Add common headers
	for k, v := range c.builder.headers {
		if len(v) > 0 {
			ret.Set(k, v[0])
		}
	}

	// Add custom headers for the current request
	for k, v := range customHeaders {
		if len(v) > 0 {
			ret.Set(k, v[0])
		}
	}

	return ret
}
