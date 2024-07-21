package main

import "os"

func PrntLine(str string) {
	// Write str to terminal
	os.Stdout.WriteString(str)
}
