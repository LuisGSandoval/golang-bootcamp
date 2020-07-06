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

	arg := os.Args[1]

	feet, _ := strconv.ParseFloat(arg, 64)

	meters := feet * 0.3048

	fmt.Println(feet, "ft. = ", meters, "mts.")

}
