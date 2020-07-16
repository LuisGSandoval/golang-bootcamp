package switchstatement

import (
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// STORY
//  You want to write a program that will manipulate a
//  given string to uppercase, lowercase, and title case.
//
// EXERCISE: String Manipulator
//
//  1. Get the operation as the first argument.
//
//  2. Get the string to be manipulated as the 2nd argument.
//
// HINT
//  You can find the manipulation functions in the strings
//  package Go documentation (ToLower, ToUpper, Title).
//
// EXPECTED OUTPUT
//
//  go run main.go
//    [command] [string]
//
//    Available commands: lower, upper and title
//
//  go run main.go lower 'OMG!'
//    omg!
//
//  go run main.go upper 'omg!'
//    OMG!
//
//  go run main.go title "mr. charles darwin"
//    Mr. Charles Darwin
//
//  go run main.go genius "mr. charles darwin"
//    Unknown command: "genius"
// ---------------------------------------------------------

// StringManipulator is a study function
func StringManipulator() {

	args := os.Args

	if len(args) != 3 {
		fmt.Println("usage [command] [string] \n\n\n Available commands: lower, upper and title")
		return
	}

	var resp string

	switch {
	case args[1] == "lower":
		resp = strings.ToLower(args[2])
	case args[1] == "upper":
		resp = strings.ToUpper(args[2])
	case args[1] == "title":
		resp = strings.Title(args[2])
	default:
		fmt.Println("sorry, don't know what you mean")

	}

	fmt.Println(resp)
}
