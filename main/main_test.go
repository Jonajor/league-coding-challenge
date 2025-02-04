package main

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Test Echo Endpoint
func Test_echoEndpoint(t *testing.T) {
	// Simulated CSV content
	matrix := "1,2,3\n4,5,6\n7,8,9"

	// Create a buffer to write our multipart request into
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create a form file field (simulating file upload)
	part, err := writer.CreateFormFile("file", "matrix.csv")
	if err != nil {
		t.Fatalf("Failed to create form file: %v", err)
	}

	// Write the matrix content into the file part
	_, err = part.Write([]byte(matrix))
	if err != nil {
		t.Fatalf("Failed to write CSV content: %v", err)
	}

	// Close the writer to finalize the multipart form
	writer.Close()

	// Create the request using the body
	req := httptest.NewRequest("POST", "/echo", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Record the response
	w := httptest.NewRecorder()

	// Call the actual handler
	handleEcho(w, req)

	// Expected response
	expected := "1,2,3\n4,5,6\n7,8,9"

	// Compare response with expected output (trim spaces and newlines)
	actual := strings.TrimSpace(w.Body.String())
	expected = strings.TrimSpace(expected)

	if w.Code != http.StatusOK {
		t.Errorf("Expected HTTP 200 OK but got %d", w.Code)
	}
	if actual != expected {
		t.Errorf("Expected response:\n%s\nGot:\n%s", expected, actual)
	}
}
