package strung

func Contains(str, subStr string) bool {
	var status bool

	// Look for instances of subStr
	for i := 0; i < len(str); i++ {
		if str[i:i+len(subStr)] == subStr {
			status = true
			break // Once found, break for efficiency
		}
	}
	return status
}

func Split(str, sep string) []string {
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

// Check if file contains non-dot and non-hash characters
func IsHashDot(s []string) bool {
	status := true
	for i := range s {
		for _, ch := range s[i] {
			if ch != '#' && ch != '.' {
				status = false
				break
			}
		}
	}

	return status
}
