package multidimencionalarray

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Moodly this function returns the added name with a mood
// Done with multidimentional array
func Moodly() {

	moods := [][]string{
		{
			"happy like a dog 😃",
			"blown away 🤯",
			"awesome 😎",
			"celebrating 🥳",
		},
		{
			"swearing up and down 🤬",
			"mad 😡",
			"sad 😔",
			"exhausted 😫",
		},
	}

	if len(os.Args) < 2 {
		fmt.Println("[your name] [positive|negative]")
		return
	}

	var name = os.Args[1:2]
	var feeling = os.Args[2:3]
	var feelingInd int

	if feeling[0] != "positive" && feeling[0] != "negative" {
		fmt.Println("Second value can only be positive or negative")
		return
	}

	rand.Seed(time.Now().UnixNano())
	randMood := rand.Intn(len(moods[1]))

	if feeling[0] == "positive" {
		feelingInd = 0
	} else if feeling[0] == "negative" {
		feelingInd = 1
	}

	fmt.Println(name[0], "is", moods[feelingInd][randMood])

}
