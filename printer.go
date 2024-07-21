package main

import "os"

func PrintLine(str string) {
	// Write str to terminal
	os.Stdout.WriteString(str+ "\n")
}
