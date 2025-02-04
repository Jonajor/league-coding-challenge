package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
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
//		curl -F 'file=@files/matrix.csv' "localhost:8080/invert"
//		curl -F 'file=@files/matrix.csv' "localhost:8080/flatten"
//		curl -F 'file=@files/matrix.csv' "localhost:8080/sum"
//		curl -F 'file=@files/matrix.csv' "localhost:8080/multiply"

func main() {
	http.HandleFunc("/", handleEcho)
	http.HandleFunc("/invert", handleInvert)
	http.HandleFunc("/flatten", handleFlatten)
	http.HandleFunc("/sum", handleSum)
	http.HandleFunc("/multiply", handleMultiply)

	http.ListenAndServe(":8080", nil)
}

// 1. Echo (given) - Return the matrix as a string in matrix format.
func handleEcho(w http.ResponseWriter, r *http.Request) {
	records, err := parseCSV(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, matrixOutput(records))
}

// 2. Invert - Return the matrix as a string in matrix format where the columns and rows are inverted
func handleInvert(w http.ResponseWriter, r *http.Request) {
	records, err := parseCSV(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, matrixOutput(transpose(records)))
}

// 3. Flatten - Return the matrix as a 1 line string, with values separated by commas.
func handleFlatten(w http.ResponseWriter, r *http.Request) {
	records, err := parseCSV(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, strings.Join(flatten(records), ","))
}

// 4. Sum - Return the sum of the integers in the matrix
func handleSum(w http.ResponseWriter, r *http.Request) {
	records, err := parseCSV(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, sum(flatten(records)))
}

// 5. Multiply - Return the product of the integers in the matrix
func handleMultiply(w http.ResponseWriter, r *http.Request) {
	records, err := parseCSV(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, multiply(flatten(records)))
}

// Common function to parse the CSV file
func parseCSV(r *http.Request) ([][]string, error) {
	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, fmt.Errorf("error to read file: %s", err)
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error to parse file: %s", err)
	}

	// Validate matrix
	if err := validateMatrix(records); err != nil {
		return nil, err
	}

	return records, nil
}
