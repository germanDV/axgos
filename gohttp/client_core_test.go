package gohttp

import (
	"testing"
)

func TestGetReqBody(t *testing.T) {
	client := httpClient{}

	t.Run("No body", func(t *testing.T) {
		body, err := client.getReqBody("", nil)
		if err != nil {
			t.Errorf("Expected no error, got %s\n", err)
		}
		if body != nil {
			t.Errorf("Expected no body, got %s\n", body)
		}
	})

	t.Run("JSON", func(t *testing.T) {
		payload := []string{"one", "two"}
		body, err := client.getReqBody("application/json", payload)
		if err != nil {
			t.Errorf("Expected no error, got %s\n", err)
		}

		want := `["one","two"]`
		got := string(body)
		if want != got {
			t.Errorf("Expected %s, got: %s\n", want, got)
		}
	})

	t.Run("JSON as default", func(t *testing.T) {
		payload := []string{"a", "b"}
		body, err := client.getReqBody("", payload)
		if err != nil {
			t.Errorf("Expected no error, got %s\n", err)
		}

		want := `["a","b"]`
		got := string(body)
		if want != got {
			t.Errorf("Expected %s, got: %s\n", want, got)
		}
	})

	t.Run("XML", func(t *testing.T) {
		payload := []string{"foo", "bar"}
		body, err := client.getReqBody("application/xml", payload)
		if err != nil {
			t.Errorf("Expected no error, got %s\n", err)
		}

		want := `<string>foo</string><string>bar</string>`
		got := string(body)
		if want != got {
			t.Errorf("Expected %s, got: %s\n", want, got)
		}
	})
}
