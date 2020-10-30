package structwarmup

import "fmt"

// Warmup testing
func Warmup() {

	// ---------------------------------------------------------
	// EXERCISE: Warm Up
	//
	//  Starting with this exercise, you'll build a command-line
	//  game store.
	//
	//  1. Declare the following structs:
	//
	//     + item: id (int), name (string), price (int)
	//
	//     + game: embed the item, genre (string)
	//
	//
	//  2. Create a game slice using the following data:
	//
	//     id  name          price    genre
	//
	//     1   god of war    50       action adventure
	//     2   x-com 2       30       strategy
	//     3   minecraft     20       sandbox
	//
	//
	//  3. Print all the games.
	//
	// EXPECTED OUTPUT
	//  Please run the solution to see the output.
	// ---------------------------------------------------------

	type Item struct {
		id    int
		name  string
		price int
	}

	type Game struct {
		genre string
		Item
	}

	data := []Game{
		{
			Item: Item{
				id:    1,
				name:  "God of war",
				price: 50,
			},
			genre: "action adventure",
		},
		{
			Item: Item{
				id:    2,
				name:  "x-com 2",
				price: 30,
			},
			genre: "action adventure",
		},
		{
			Item: Item{
				id:    3,
				name:  "minecraft",
				price: 20,
			},
			genre: "action adventure",
		},
	}

	fmt.Print(data)

}
