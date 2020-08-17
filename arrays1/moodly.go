package arrays1

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Moodly this function returns the added name with a mood
func Moodly() {

	moods := [...]string{
		"happy like a dog 😃",
		"wondering   🤔",
		"mad 😡",
		"swearing up and down 🤬",
		"blown away 🤯",
	}

	if len(os.Args) < 2 {
		fmt.Println("[your name]")
		return
	}

	var name = os.Args[1:2]

	rand.Seed(time.Now().UnixNano())
	randMood := rand.Intn(len(moods))

	fmt.Println(name[0], "is", moods[randMood])

}
