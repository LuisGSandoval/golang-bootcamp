package structwarmup

import "fmt"

// ListGames testing
func ListGames() {

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
