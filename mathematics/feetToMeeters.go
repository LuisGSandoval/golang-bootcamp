package mathematics

import (
	"fmt"
	"os"
	"strconv"
)

// FeetToMeeters is a function that returns the equivalent meters to a given distance in feet
func FeetToMeeters() {

	fmt.Println("Distance conversion")
	fmt.Println(" 👣Feet to meters 𐃏")

	if len(os.Args) < 2 || len(os.Args[1]) < 1 {
		fmt.Println("please enter a value")
		return
	}

	arg := os.Args[1]

	feet, err := strconv.ParseFloat(arg, 64)

	if err != nil {
		fmt.Printf("🚫👎🏼Error: invalid input %v \n", arg)
		return
	}

	meters := feet * 0.3048

	fmt.Println(feet, "ft. = ", meters, "mts.")

}
