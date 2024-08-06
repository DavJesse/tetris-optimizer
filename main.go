package main

import (
	"os"

	"tetris/grid"
	"tetris/misc"
	"tetris/tetro"
)

func main() {
	arg := os.Args[1:]

	// Check if user has included the correct number of arguments (outside program name)
	if len(arg) != 1 {
		// Handle instances of no arguments
		if len(arg) == 0 {
			misc.PrintLine("Please include the name of a text file as an argument")
			return
		} else {
			// Handle more than one argument
			misc.PrintLine("You have way too many arguments!\nOnly include one text file as argument")
			return
		}
	}

	// Store contents of arg in 'file'
	file := arg[0]

	// End program in the instance of non-text files
	if !misc.CheckExtension(file) {
		misc.PrintLine(tetro.Errors())
		return
	}

	content, errFile := misc.ReadFile(file) // Read the file parsed as argument

	// React to errors in reading file and when file is empty
	// Exit program in both circumstances
	if errFile != "" || len(content) == 0 {
		if len(content) == 0 {
			return
		}
		misc.PrintLine(errFile)
		return
	}

	tetroSlc := misc.TwoD(content) // Extract tetrominoes in file and store them in 2D slice

	tetroSlc, errValidity := tetro.CheckValidity(tetroSlc) // Check tetrominos for validity

	// In case of error, print error message, end program
	if errValidity != "" {
		misc.PrintLine(errValidity)
		return
	}

	size := grid.Solve(tetroSlc)

	grid.PrintGrid(size, tetroSlc)
}
