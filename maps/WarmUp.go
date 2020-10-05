package maps

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Warm-up
//
//  Create and print the following maps.
//
//  1. Phone numbers by last name
//  2. Product availability by Product ID
//  3. Multiple phone numbers by last name
//  4. Shopping basket by Customer ID
//
//     Each item in the shopping basket has a Product ID and
//     quantity. Through the map, you can tell:
//     "Mr. X has bought Y bananas"
//
// ---------------------------------------------------------

// WarmUp testing
func WarmUp() {
	// Hint: Store phone numbers as text

	// #1
	// Key        : Last name
	// Element    : Phone number
	agenda := map[string]string{
		"Camargo":  "3022161006",
		"Sandoval": "3133102680",
		"Martinez": "3133102680",
	}

	// #2
	// Key        : Product ID
	// Element    : Available / Unavailable

	products := map[int]bool{
		5436534: false,
		5436535: true,
		5436536: true,
		5436537: false,
	}

	// #3
	// Key        : Last name
	// Element    : Phone numbers
	complexAgenda := map[string][]string{
		"Sandoval": {"3133102680", "3118216617"},
	}

	// #4
	// Key        : Customer ID
	// Element Key:
	//   Key: Product ID Element: Quantity

	basket := map[int]map[int]int{
		1: {
			123:   4,
			5435:  3,
			54624: 1,
			896:   1,
		},
	}

	fmt.Printf("phones     : %#v\n", agenda)
	fmt.Printf("products   : %#v\n", products)
	fmt.Printf("multiPhones: %#v\n", complexAgenda)
	fmt.Printf("basket     : %#v\n", basket)

}
