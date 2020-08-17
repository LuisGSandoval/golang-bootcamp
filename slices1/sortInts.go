package slices1

import (
	"fmt"
	"sort"
)

// SortingAnArray testing
func SortingAnArray() {
	intsArr := [6]int{1, 56, 3, 87, 2, 234}
	fmt.Println(intsArr)

	// we can use a built in function to aort the array,
	// the thing is that it only works with slices
	// so we're going to turn an array to a slice
	//   + how?
	//   int[:] => adding courly braces with a colon inside

	sort.Ints(intsArr[:])

	fmt.Println(intsArr)

}
