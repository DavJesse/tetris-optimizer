package misc

import (
	"log"
	"os"
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
