package main

import "os"

func main() {
	arg := os.Args[1:]

	if len(arg) != 1 {
		if len(arg) == 0 {
			PrintLine("Please include the name of a text file as an argument")
		}
	}
}
