package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Test validateMatrix
func Test_validateMatrix(t *testing.T) {
	// Scenario 1: empty csv file
	emptyMatrix := [][]string{}
	err_file := validateMatrix(emptyMatrix)

	if err_file == nil {
		t.Errorf("Empty CSV file %v", err_file)
	}

	// Scenario 2: valid matrix
	validMatrix := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}

	err := validateMatrix(validMatrix)
	if err != nil {
		t.Errorf("validateMatrix() error = %v", err)
	}

	// Scenario 3: non NxN matrix
	nonSquareMatrix := [][]string{
		{"1", "2", "3"},
		{"4", "5"},
	}
	err = validateMatrix(nonSquareMatrix)
	if err == nil {
		t.Errorf("Expected error for non-square matrix, but got nil")
	}

	// Scenario 4: Non integer matrix
	nonIntegerMatrix := [][]string{
		{"1", "2", "x"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	err = validateMatrix(nonIntegerMatrix)
	if err == nil {
		t.Errorf("Expected error for non-integer values, but got nil")
	}
}

// Test transpose
func Test_transpose(t *testing.T) {
	matrix := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}

	expected := "1,4,7\n2,5,8\n3,6,9\n"

	w := httptest.NewRecorder()

	transpose(w, matrix)

	result := w.Body.String()

	// Compare with expected output
	if strings.TrimSpace(result) != strings.TrimSpace(expected) {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}

// Test flatten
func TestFlatten(t *testing.T) {
	matrix := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	expected := "1,2,3,4,5,6,7,8,9\n"

	w := httptest.NewRecorder()
	flatten(w, matrix)

	if w.Body.String() != expected {
		t.Errorf("Expected %s but got %s", expected, w.Body.String())
	}
}

// Test sum
func Test_sum(t *testing.T) {
	numbers := []string{"1", "2", "3", "4", "5"}
	expected := "15\n"

	w := httptest.NewRecorder()
	sum(w, numbers)

	if w.Body.String() != expected {
		t.Errorf("Expected sum %s but got %s", expected, w.Body.String())
	}
}

// Test multiply
func Test_multiply(t *testing.T) {
	numbers := []string{"1", "2", "3", "4", "5"}
	expected := "120\n"

	w := httptest.NewRecorder()
	multiply(w, numbers)

	if w.Body.String() != expected {
		t.Errorf("Expected multiply %s but got %s", expected, w.Body.String())
	}
}

// Test Echo Endpoint
func Test_echoEndpoint(t *testing.T) {
	requestBody := "file=@matrix.csv"

	req := httptest.NewRequest("POST", "/", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "multipart/form-data")

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate echo response
		matrixOutput(w, [][]string{
			{"1", "2", "3"},
			{"4", "5", "6"},
			{"7", "8", "9"},
		})
	})

	handler.ServeHTTP(w, req)

	expected := "1,2,3\n4,5,6\n7,8,9\n"
	if strings.TrimSpace(w.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("Expected response:\n%s\nGot:\n%s", expected, w.Body.String())
	}
}
