package slices1

import (
	"fmt"
	"strconv"
	"strings"
)

// ExcSliceNumbers testing
func ExcSliceNumbers() {

	// ---------------------------------------------------------
	// EXERCISE: Slice the numbers
	//
	//   We've a string that contains even and odd numbers.
	//
	//   1. Convert the string to an []int
	//
	//   2. Print the slice
	//
	//   3. Slice it for the even numbers and print it (assign it to a new slice variable)
	//
	//   4. Slice it for the odd numbers and print it (assign it to a new slice variable)
	//
	//   5. Slice it for the two numbers at the middle
	//
	//   6. Slice it for the first two numbers
	//
	//   7. Slice it for the last two numbers (use the len function)
	//
	//   8. Slice the evens slice for the last one number
	//
	//   9. Slice the odds slice for the last two numbers
	//
	//
	// EXPECTED OUTPUT
	//   go run main.go
	//
	//     nums        : [2 4 6 1 3 5]
	//     evens       : [2 4 6]
	//     odds        : [1 3 5]
	//     middle      : [6 1]
	//     first 2     : [2 4]
	//     last 2      : [3 5]
	//     evens last 1: [6]
	//     odds last 2 : [3 5]
	//
	//
	// NOTE
	//
	//  You can also use my prettyslice package for printing the slices.
	//
	//
	// HINT
	//
	//  Find a function in the strings package for splitting the string into []string
	//
	// ---------------------------------------------------------

	data := "2 4 6 1 3 5"

	var (
		numsStr []string
		numsint []int
		evens   []int
		odds    []int
	)

	numsStr = strings.Fields(data)

	for _, num := range numsStr {
		num, _ := strconv.Atoi(num)
		numsint = append(numsint, num)

		// All evens
		if num%2 == 0 {
			evens = append(evens, num)
		}

		// All odds
		if num%2 != 0 {
			odds = append(odds, num)
		}

	}

	fmt.Printf("%-12s %d \n", "nums", numsint)
	fmt.Printf("%-12s %d \n", "evens", evens)
	fmt.Printf("%-12s %d \n", "odds", odds)
	fmt.Printf("%-12s %d \n", "middle", numsint[2:4])
	fmt.Printf("%-12s %d \n", "first 2", numsint[:2])
	fmt.Printf("%-12s %d \n", "last 2", numsint[len(numsint)-1-2:])
	fmt.Printf("%-12s %d \n", "evens last 1", evens[len(evens)-1:])
	fmt.Printf("%-12s %d \n", "odds last 2", odds[len(odds)-1-2:])
}
