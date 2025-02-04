package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Open a new terminal
// Access the project package
//      cd league-coding-challenge/
// Access the main package
//      cd main
// Run with
//		go run .
// Send request with:
//		curl -F 'file=@files/matrix.csv' "localhost:8080/echo"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error to read file: %s", err.Error())))
			return
		}
		defer file.Close()
		records, err := csv.NewReader(file).ReadAll()
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error to parse file: %s", err.Error())))
			return
		}

		// Validate matrix
		if err := validateMatrix(records); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Fprintln(w, "Input")
		matrixOutput(w, records)

		fmt.Fprintln(w, "Output")

		fmt.Fprintln(w)
		fmt.Fprintln(w, "Invert Matrix")
		transpose(w, records)

		fmt.Fprintln(w)
		fmt.Fprintln(w, "Flatten")
		numbers := flatten(w, records)

		fmt.Fprintln(w)
		fmt.Fprintln(w, "Sum")
		sum(w, numbers)

		fmt.Fprintln(w)
		fmt.Fprintln(w, "Multiply")
		multiply(w, numbers)
	})
	http.ListenAndServe(":8080", nil)
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

// Transpose: Return the matrix as a string in matrix format where the columns and
func transpose(w http.ResponseWriter, records [][]string) {
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

	matrixOutput(w, transposed)
}

// Flatten: Return the matrix as a 1 line string, with values separated by commas.
func flatten(w http.ResponseWriter, records [][]string) []string {
	var list []string
	for _, row := range records {
		list = append(list, row...)
	}
	fmt.Fprintln(w, strings.Join(list, ","))
	return list
}

// Sum: Return the sum of the integers in the matrix
func sum(w http.ResponseWriter, numbers []string) {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		convertNumbers, _ := strconv.Atoi(numbers[i])
		sum += convertNumbers
	}
	fmt.Fprintln(w, sum)
}

// Multiply: Return the product of the integers in the matrix
func multiply(w http.ResponseWriter, numbers []string) {
	multiply := 1
	for i := 0; i < len(numbers); i++ {
		convertedNumbers, _ := strconv.Atoi(numbers[i])
		multiply *= convertedNumbers
	}
	fmt.Fprintln(w, multiply)
}

func matrixOutput(w http.ResponseWriter, records [][]string) {
	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	fmt.Fprintln(w, response)
}
