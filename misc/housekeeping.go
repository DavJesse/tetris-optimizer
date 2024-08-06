package misc

import (
	"os"
	"runtime"

	"tetris/strung"
	"tetris/tetro"
)

func CheckExtension(str string) bool {
	var status bool
	start := len(str) - 4 // Start index of .txt extension

	// Check if str ends with .txt
	if str[start:] == ".txt" {
		status = true
	}
	return status
}

func ReadFile(file string) (string, string) {
	var errors string

	// Read file parse
	content, err := os.ReadFile(file)
	// Print error & exit program should reading fail
	// Populate error message should there be an error
	if err != nil {
		errors = tetro.Errors()
	}

	return string(content), errors
}

func TwoD(str string) [][]string {
	var result [][]string
	var token []string
	sep := osChecker()

	raw := strung.Split(str, sep)

	for i := 0; i < len(raw); i++ {
		// Find empty stings separating tetrominoes...
		// Append contents of token to result
		if raw[i] == "" {
			if len(token) > 0 { // Only append populated tokens
				result = append(result, token)
				token = nil // Empty token
			}

			// Add non-empty strings to token
		} else {
			token = append(token, raw[i])
		}
	}
	// Append left over slices at end of loop
	// Only append populated slices
	if len(token) > 0 {
		result = append(result, token)
	}

	return result
}

func osChecker() string {
	var sep string
	os := runtime.GOOS // check user's operating system

	// If user is running windows, nominate "\r\n" as seperator
	if os == "windows" {
		sep = "\r\n"
	} else {
		sep = "\n" // Use "\n" as default separator
	}

	return sep
}
