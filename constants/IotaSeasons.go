package constants

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Iota Seasons
//
//  Use iota to initialize the season constants.
//
// HINT
//  You can change the order of the constants.
//
// EXPECTED OUTPUT
//  12 3 6 9
// ---------------------------------------------------------

// IotaSeasons is a function
func IotaSeasons() {
	// NOTE : You should remove all the initializers below
	//        first. Then use iota to fix it.
	const (
		Spring = (iota + 1) * 3
		Summer
		Fall
		Winter
	)

	fmt.Println(Winter, Spring, Summer, Fall)
}
