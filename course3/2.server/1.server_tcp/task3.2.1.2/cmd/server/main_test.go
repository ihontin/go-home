package main

import (
	"github.com/jarcoal/httpmock"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestServer(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	expectedBody := `<!DOCTYPE html>
<html>
<head>
<title>Webserver</title>
</head>
<body>
hello world
</body>
</html>`
	httpmock.RegisterResponder("GET", "http://localhost:8080/",
		httpmock.NewStringResponder(200, expectedBody))

	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if strings.TrimSpace(string(body)) != strings.TrimSpace(expectedBody) {
		t.Errorf("Expected body %q, got %q", expectedBody, body)
	}
}
