package slices1

import "fmt"

// SlicesAsignedLiteralValues testing
func SlicesAsignedLiteralValues() {

	// ---------------------------------------------------------
	// EXERCISE: Assign empty slices
	//
	//   Assign empty slices to all the slices that you've declared in the previous
	//   exercise, and print them here.
	//
	//
	// EXPECTED OUTPUT
	//  names    : []string 0 false
	//  distances: []int 0 false
	//  data     : []uint8 0 false
	//  ratios   : []float64 0 false
	//  alives   : []bool 0 false
	// ---------------------------------------------------------

	var (
		names     []string
		distances []int
		data      []uint8
		ratios    []float64
		alives    []bool
	)

	names = []string{"Luis", "Nataly", "Casper"}
	distances = []int{2, 57, 6}
	data = []uint8{3, 6, 255}
	ratios = []float64{4.3, .6, .004}
	alives = []bool{true, false}

	fmt.Printf("%-10s : %T %d %t\n", "names", names, len(names), names == nil)
	fmt.Printf("%-10s : %T %d %t\n", "distances", distances, len(distances), distances == nil)
	fmt.Printf("%-10s : %T %d %t\n", "data", data, len(data), data == nil)
	fmt.Printf("%-10s : %T %d %t\n", "ratios", ratios, len(ratios), ratios == nil)
	fmt.Printf("%-10s : %T %d %t\n", "alives", alives, len(alives), alives == nil)

}
