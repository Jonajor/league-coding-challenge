# League Backend Challenge - Matrix API

## Overview
his Go web service is built to solve the League backend challenge and processes a CSV matrix while providing five key operations:
- **Echo** (returns the matrix)
- **Invert** (transposes the matrix)
- **Flatten** (converts to a single row)
- **Sum** (adds all numbers)
- **Multiply** (multiplies all numbers)

### API Operations and Expected Responses:

The API expected a **CSV file** to be uploaded. The example receive by email (`matrix.csv`) is is included in the `files` directory.

```csv
1,2,3
4,5,6
7,8,9
```

| **Endpoint**   | **Functionality** | **Expected Response** |
|--------------|----------------|-------------------|
| `/echo` | Returns the original matrix. | `1,2,3\n4,5,6\n7,8,9` |
| `/invert` | Returns the inverted matrix. | `1,4,7\n2,5,8\n3,6,9` |
| `/flatten` | Returns matrix as a single row. | `1,2,3,4,5,6,7,8,9` |
| `/sum` | Returns sum of all elements. | `45` |
| `/multiply` | Returns product of all elements. | `362880` |

---

## How to Run

### 1️ - Install Go
Ensure you have Go installed:
```sh
go version
```

### 2 - Clone the Project from github
On terminal and clone the project below is the command for that:
```sh
git clone https://github.com/Jonajor/league-coding-challenge.git
```

### 3 - Navigate to the Project Directory
```sh
cd league-coding-challenge/main
```

### 4 - Running Application
```sh
go run .
```

### 5 - Sending Requests to the API
```sh
curl -F 'file=@files/matrix.csv' "http://localhost:8080/echo"
curl -F 'file=@files/matrix.csv' "http://localhost:8080/invert"
curl -F 'file=@files/matrix.csv' "http://localhost:8080/flatten"
curl -F 'file=@files/matrix.csv' "http://localhost:8080/sum"
curl -F 'file=@files/matrix.csv' "http://localhost:8080/multiply"
```

### 6 - Testing & Code Coverage
To check test coverage and code quality, run:
```sh
go test -v -cover
```

Expected Test Output:
```sh
=== RUN   Test_HandleEcho
--- PASS: Test_HandleEcho (0.01s)
=== RUN   Test_HandleInvert
--- PASS: Test_HandleInvert (0.00s)
=== RUN   Test_HandleFlatten
--- PASS: Test_HandleFlatten (0.00s)
=== RUN   Test_HandleSum
--- PASS: Test_HandleSum (0.00s)
=== RUN   Test_HandleMultiply
--- PASS: Test_HandleMultiply (0.00s)
=== RUN   Test_HandleMissingFile
--- PASS: Test_HandleMissingFile (0.00s)
=== RUN   Test_HandleNonSquareMatrix
--- PASS: Test_HandleNonSquareMatrix (0.00s)
=== RUN   Test_HandleNonIntegerMatrixEcho
--- PASS: Test_HandleNonIntegerMatrixEcho (0.00s)
=== RUN   Test_HandleNonIntegerMatrixInvert
--- PASS: Test_HandleNonIntegerMatrixInvert (0.00s)
=== RUN   Test_HandleNonIntegerMatrixFlatten
--- PASS: Test_HandleNonIntegerMatrixFlatten (0.00s)
=== RUN   Test_HandleNonIntegerMatrixSum
--- PASS: Test_HandleNonIntegerMatrixSum (0.00s)
=== RUN   Test_HandleNonIntegerMatrixMultiply
--- PASS: Test_HandleNonIntegerMatrixMultiply (0.00s)
=== RUN   Test_validateMatrix
--- PASS: Test_validateMatrix (0.00s)
=== RUN   Test_transpose
--- PASS: Test_transpose (0.00s)
=== RUN   Test_Flatten
--- PASS: Test_Flatten (0.00s)
=== RUN   Test_sum
--- PASS: Test_sum (0.00s)
=== RUN   Test_multiply
--- PASS: Test_multiply (0.00s)
=== RUN   Test_matrixOutput
--- PASS: Test_matrixOutput (0.00s)
PASS
coverage: 89.2% of statements
ok      main/main       0.539s
```