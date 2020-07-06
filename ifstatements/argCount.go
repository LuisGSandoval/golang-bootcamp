package ifstatements

import (
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Arg Count
//
//  1. Get arguments from command-line.
//  2. Print the expected outputs below depending on the number
//     of arguments.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me args
//
//  go run main.go hello
//    There is one: "hello"
//
//  go run main.go hi there
//    There are two: "hi there"
//
//  go run main.go i wanna be a gopher
//    There are 5 arguments
// ---------------------------------------------------------

// ArgCount is a test function
func ArgCount() {

	argLength := len(os.Args) - 1

	argsString := os.Args[1:]

	finalstring := strings.Join(argsString, " ")

	if argLength < 1 {
		fmt.Println("Give me args")
	} else if argLength == 1 {
		fmt.Printf("There is one: %q \n", finalstring)
	} else if argLength == 1 {
		fmt.Printf("There are two: %q \n", finalstring)
	} else {
		fmt.Printf("There are %d args: %q \n", argLength, finalstring)
	}

}
