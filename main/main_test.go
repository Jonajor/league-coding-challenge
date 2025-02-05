package main

import (
	"bytes"
	"mime/multipart"

	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test Echo Endpoint
func Test_HandleEcho(t *testing.T) {
	matrix := "1,2,3\n4,5,6\n7,8,9"
	expected := "1,2,3\n4,5,6\n7,8,9"

	w, req := createUploadRequest(t, "/echo", matrix)
	handleEcho(w, req)

	if strings.TrimSpace(w.Body.String()) != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, w.Body.String())
	}
}

// Test invert endpoint
func Test_HandleInvert(t *testing.T) {
	matrix := "1,2,3\n4,5,6\n7,8,9"
	expected := "1,4,7\n2,5,8\n3,6,9\n"

	w, req := createUploadRequest(t, "/invert", matrix)
	handleInvert(w, req)

	assert.Equal(t, expected, w.Body.String(), "Should be the same.")
	assert.Equal(t, 200, w.Code, "Should return 200 status code")
}

// Test flatten endpoint
func Test_HandleFlatten(t *testing.T) {
	matrix := "1,2,3\n4,5,6\n7,8,9"
	expected := "1,2,3,4,5,6,7,8,9\n"

	w, req := createUploadRequest(t, "/flatten", matrix)
	handleFlatten(w, req)

	assert.Equal(t, expected, w.Body.String(), "Should be the same.")
	assert.Equal(t, 200, w.Code, "Should return 200 status code")
}

// Test sum endpoint
func Test_HandleSum(t *testing.T) {
	matrix := "1,2,3\n4,5,6\n7,8,9"
	expected := "45\n"

	w, req := createUploadRequest(t, "/sum", matrix)
	handleSum(w, req)

	assert.Equal(t, expected, w.Body.String(), "Should be the same.")
	assert.Equal(t, 200, w.Code, "Should return 200 status code")
}

// Test multiply endpoint
func Test_HandleMultiply(t *testing.T) {
	matrix := "1,2,3\n4,5,6\n7,8,9"
	expected := "362880\n"

	w, req := createUploadRequest(t, "/multiply", matrix)
	handleMultiply(w, req)

	assert.Equal(t, expected, w.Body.String(), "Should be the same.")
	assert.Equal(t, 200, w.Code, "Should return 200 status code")
}

// Test error handling for missing file
func Test_HandleMissingFile(t *testing.T) {
	req := httptest.NewRequest("POST", "/sum", nil)
	w := httptest.NewRecorder()

	handleSum(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected HTTP 400 for missing file, got %d", w.Code)
	}
}

// Test error handling for non-square matrix
func Test_HandleNonSquareMatrix(t *testing.T) {
	matrix := "1,2,3\n4,5,6"
	w, req := createUploadRequest(t, "/sum", matrix)

	handleSum(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Should return 400 status code")
}

// Test error handling for non-integer values
func Test_HandleNonIntegerMatrixEcho(t *testing.T) {
	matrix := "1,2,a\n4,5,6\n7,8,9"
	w, req := createUploadRequest(t, "/echo", matrix)

	handleEcho(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Should return 400 status code")
}

func Test_HandleNonIntegerMatrixInvert(t *testing.T) {
	matrix := "1,2,a\n4,5,6\n7,8,9"
	w, req := createUploadRequest(t, "/invert", matrix)

	handleEcho(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Should return 400 status code")
}

func Test_HandleNonIntegerMatrixFlatten(t *testing.T) {
	matrix := "1,2,a\n4,5,6\n7,8,9"
	w, req := createUploadRequest(t, "/flatten", matrix)

	handleFlatten(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Should return 400 status code")
}

func Test_HandleNonIntegerMatrixSum(t *testing.T) {
	matrix := "1,2,a\n4,5,6\n7,8,9"
	w, req := createUploadRequest(t, "/sum", matrix)

	handleSum(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Should return 400 status code")
}

func Test_HandleNonIntegerMatrixMultiply(t *testing.T) {
	matrix := "1,2,a\n4,5,6\n7,8,9"
	w, req := createUploadRequest(t, "/multiply", matrix)

	handleMultiply(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Should return 400 status code")
}

func createUploadRequest(t *testing.T, endpoint string, fileContent string) (*httptest.ResponseRecorder, *http.Request) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", "matrix.csv")
	if err != nil {
		t.Fatalf("Failed to create form file: %v", err)
	}

	_, err = part.Write([]byte(fileContent))
	if err != nil {
		t.Fatalf("Failed to write CSV content: %v", err)
	}
	writer.Close()

	req := httptest.NewRequest("POST", endpoint, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	return w, req
}
