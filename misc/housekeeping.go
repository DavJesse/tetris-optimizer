package misc

import (
	"log"
	"os"
	"runtime"
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

func splitString(str, sep string) []string {
	var result []string
	var token string

	for i := 0; i < len(str); i++ {
		// Find instances of separator
		if i <= len(str)-len(sep) && str[i:i+len(sep)] == sep {
			result = append(result, token) // Append contents of token to result
			token = ""                     // Empty token
			i += len(sep) - 1              // Skip characters of separator

			// Add to token characters that are not part of separator
		} else {
			token += string(str[i])
		}
	}
	// Append to result any character that may be in token at the end of loop
	result = append(result, token)
	return result
}

func TwoD(str string) [][]string {
	var result [][]string
	var token []string
	sep := osChecker()

	raw := splitString(str, sep)

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
