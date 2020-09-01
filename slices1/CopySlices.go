package slices1

import "github.com/inancgumus/prettyslice"

// CopySlices testing
func CopySlices() {
	data := []float64{10, 25, 30, 50}
	data = append(data, []float64{3}...)
	prettyslice.Show("data", data)

	copy(data[:1], []float64{55, 65, 7, 8, 9}[:1]) // the copy function in reallity replaces the content of a slices with the content we tell it to add to it

	prettyslice.Show("data", data)
}
