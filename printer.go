package main

import "os"

func PrntLine(str string) {
	// Write str to terminal
	os.StdOut.WriteString(str)
}
