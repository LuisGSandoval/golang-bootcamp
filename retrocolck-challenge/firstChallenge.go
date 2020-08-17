package retroclock

import (
	"fmt"
	"time"
)

// RenderClock is the function that will render the clock
func RenderClock() {

	for shift := 0; ; shift++ {

		// screen.Clear()
		// screen.MoveTopLeft()
		print("\033[H\033[2J") // this line clears the screen

		now := time.Now()
		hour, min, sec, _ := now.Hour(), now.Minute(), now.Second(), now.Second()/10
		ssec := now.Nanosecond() / 1e8

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10], // hours
			colon,
			digits[min/10], digits[min%10], // mins
			colon,
			digits[sec/10], digits[sec%10], // secs
			dot,
			digits[ssec],
		}

		// alarmed := sec%10 == 0
		// if alarmed {
		// 	clock = alarmSign
		// }

		for line := 0; line < 5; line++ { // this is the height of each character

			for index, digit := range clock {
				next := clock[index][line]
				if (digit == colon || digit == dot) && sec%2 == 0 {
					next = "   "
				}
				fmt.Print(next, "  ")
			}

			// l := len(clock)
			// s, e := shift%l, l
			// if shift%(l*2) >= l {
			// 	s, e = 0, s
			// }
			// for j := 0; j < l-e; j++ {
			// 	fmt.Print("     ")
			// }

			// 	// draw the digits starting from 's' to 'e'
			// for i := s; i < e; i++ {
			// 	next := clock[i][line]
			// 	if clock[i] == colon && sec%2 == 0 {
			// 		next = "   "
			// 	}

			// 	fmt.Print(next, "  ")
			// }

			fmt.Println("\r")

		}

		time.Sleep(time.Millisecond * 10)
	}
}
