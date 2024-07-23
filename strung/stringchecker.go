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
