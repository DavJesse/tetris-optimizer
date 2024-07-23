package misc

import (
	"log"
	"os"
	"runtime"

	"tetris/strung"
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

func ReadFile(file string) string {
	// Read file parse
	content, err := os.ReadFile(file)
	// Print error & exit program should reading fail
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
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
			result = append(result, token)
			token = nil // Empty token

			// Add non-empty strings to token
		} else {
			token = append(token, raw[i])
		}
	}
	result = append(result, token)

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
