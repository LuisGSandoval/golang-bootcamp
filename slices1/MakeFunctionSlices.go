package slices1

// MakeFunctionSlices testing
func MakeFunctionSlices() {
	// prettyslice.PrintBacking = true
	// prettyslice.MaxPerLine = 10

	tasks := []string{1022: "jump", "run", "read"}

	// var optimalAppend []string // they tell this is much harder task
	optimalAppend := make([]string, 0, len(tasks))

	// prettyslice.Show("optimalAppend", optimalAppend)

	for _, task := range tasks {
		optimalAppend = append(optimalAppend, task)

		// prettyslice.Show("optimalAppend", optimalAppend)

	}
}
