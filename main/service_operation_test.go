package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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

	expected := [][]string{
		{"1", "4", "7"},
		{"2", "5", "8"},
		{"3", "6", "9"},
	}

	result := transpose(matrix)

	assert.Equal(t, expected, result, "Should be the same.")

}

// Test flatten
func Test_Flatten(t *testing.T) {
	matrix := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	expected := "1,2,3,4,5,6,7,8,9"

	numbers := strings.Join(flatten(matrix), ",")

	assert.Equal(t, expected, numbers, "Should be the same.")
}

// Test sum
func Test_sum(t *testing.T) {
	numbers := []string{"1", "2", "3", "4", "5"}
	expected := 15

	result := sum(numbers)

	assert.Equal(t, expected, result, "Should be the same.")
}

// Test multiply
func Test_multiply(t *testing.T) {
	numbers := []string{"1", "2", "3", "4", "5"}
	expected := 120

	result := multiply(numbers)

	assert.Equal(t, expected, result, "Should be the same.")
}
