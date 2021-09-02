package axgos

import (
	"net/http"
	"testing"
)

func TestGetHeaders(t *testing.T) {
	h1 := make(http.Header)
	h1.Set("Custom-Header", "X-Testing-1")

	h2 := make(http.Header)
	h2.Set("Custom-Header", "X-Testing-2")

	headers := getHeaders(h1, h2)
	if len(headers) != 1 {
		t.Errorf("Want only one header, got %d\n", len(headers))
	}
	if headers.Get("Custom-Header") != "X-Testing-1" {
		t.Errorf("Want Custom-Header value X-Testing-1, got %q\n", headers.Get("Custom-Header"))
	}
}

func TestGetHeaderFromClient(t *testing.T) {
	commonHeaders := make(http.Header)
	commonHeaders.Set("Common", "X-Common-Value")

	builder := axgosBuilder{headers: commonHeaders}
	client := axgosClient{builder: &builder}

	reqHeader := make(http.Header)
	reqHeader.Set("Request", "X-Request-Value")

	gotHeaders := client.getHeaders(reqHeader)
	if len(gotHeaders) != 2 {
		t.Errorf("Expected 2 headers, got %d\n", len(gotHeaders))
	}
	if gotHeaders.Get("Common") != "X-Common-Value" {
		t.Errorf("Expected Common header to be X-Common-Value, got %q\n", gotHeaders.Get("Common"))
	}
	if gotHeaders.Get("Request") != "X-Request-Value" {
		t.Errorf("Expected Common header to be X-Request-Value, got %q\n", gotHeaders.Get("Request"))
	}
}

func TestSpecificHeaderOverridesCommon(t *testing.T) {
	commonHeaders := make(http.Header)
	commonHeaders.Set("User-Agent", "Common")

	builder := axgosBuilder{headers: commonHeaders}
	client := axgosClient{builder: &builder}

	reqHeader := make(http.Header)
	reqHeader.Set("User-Agent", "Specific")

	gotHeaders := client.getHeaders(reqHeader)
	if len(gotHeaders) != 1 {
		t.Errorf("Expected 1 header, got %d\n", len(gotHeaders))
	}
	if gotHeaders.Get("User-Agent") != "Specific" {
		t.Errorf("Expected header to be Specific, got %q\n", gotHeaders.Get("User-Agent"))
	}
}
