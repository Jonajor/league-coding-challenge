package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Transpose: Return the matrix as a string in matrix format where the columns and
func transpose(records [][]string) [][]string {
	rows := len(records)
	cols := len(records[0])
	transposed := make([][]string, cols)

	for i := range transposed {
		transposed[i] = make([]string, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = records[i][j]
		}
	}

	return transposed
}

// validateMatrix checks whether the CSV matrix is a valid NxN square matrix
func validateMatrix(records [][]string) error {
	rows := len(records)
	if rows == 0 {
		return fmt.Errorf("CSV file is empty")
	}

	cols := len(records[0])
	for _, row := range records {
		if len(row) != cols {
			return fmt.Errorf("invalid matrix: inconsistent number of columns")
		}
		for _, cell := range row {
			if _, err := strconv.Atoi(cell); err != nil {
				return fmt.Errorf("invalid matrix: all values must be integers")
			}
		}
	}

	// Ensure it's a square matrix
	if rows != cols {
		return fmt.Errorf("invalid matrix: must be a square matrix (NxN)")
	}

	return nil
}

// Flatten: Return the matrix as a 1 line string, with values separated by commas.
func flatten(records [][]string) []string {
	var list []string
	for _, row := range records {
		list = append(list, row...)
	}
	return list
}

// Sum: Return the sum of the integers in the matrix
func sum(numbers []string) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		convertNumbers, _ := strconv.Atoi(numbers[i])
		sum += convertNumbers
	}
	return sum
}

// Multiply: Return the product of the integers in the matrix
func multiply(numbers []string) int {
	multiply := 1
	for i := 0; i < len(numbers); i++ {
		convertedNumbers, _ := strconv.Atoi(numbers[i])
		multiply *= convertedNumbers
	}
	return multiply
}

func matrixOutput(records [][]string) string {
	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}

	return response
}
