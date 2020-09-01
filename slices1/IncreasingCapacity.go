package slices1

import "github.com/inancgumus/prettyslice"

// IncreasingCapacity testing
func IncreasingCapacity() {

	// in this test we can see how the backing array capacity increases about 25% only instead of double
	prettyslice.PrintBacking = true

	words := []string{1023: ""} // if we chage the number to 1022 and not 1023 the capacity changes double
	words = append(words, "boom!")

	prettyslice.Show("Prueba 1", words)

}
