package slices1

import (
	"fmt"
	"strings"
)

// SliceComparison tetstin
func SliceComparison() {

	fmt.Print("\n\n\n************  ARRAYS  ************\n\n")

	var arr1 [3]string
	arr2 := [3]string{}
	arr3 := [3]string{"hola"}

	arr3[2] = "hi"

	fmt.Printf("len: %d | arr1: %#v \n", len(arr1), arr1)
	fmt.Printf("len: %d | arr3: %#v \n", len(arr2), arr2)
	fmt.Printf("len: %d | arr2: %#v \n", len(arr3), arr3)

	fmt.Print("\n\n\n************  SLICES  ************\n\n")

	var slc1 []string
	slc2 := []string{}
	slc3 := []string{"hello"}

	// slc3[2] = "hi" // can not assing to a slice like an array, you have to append
	slc3 = append(slc3, "hi")

	fmt.Printf("len: %d |    slc1: %#v \n", len(slc1), slc1)
	fmt.Printf("len: %d |    slc3: %#v \n", len(slc2), slc2)
	fmt.Printf("len: %d |    slc2: %#v \n", len(slc3), slc3)

	print(strings.Repeat("\n", 7))
}
