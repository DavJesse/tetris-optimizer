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
