package constants

import (
	"fmt"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: No Conversions Allowed
//
//  1. Fix the program without doing any conversion.
//  2. Explain why it doesn't work.
//
// EXPECTED OUTPUT
//  10h0m0s later...
// ---------------------------------------------------------

// NoConversionsAllowed this is a test function
func NoConversionsAllowed() {
	const later time.Duration = 10

	hours, _ := time.ParseDuration("1h")

	fmt.Printf("%s later...\n", hours*later)
}

// Explanaition: it didn't work because later was not the same type as hours
