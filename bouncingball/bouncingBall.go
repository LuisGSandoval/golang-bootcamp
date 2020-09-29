package bouncingball

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/inancgumus/screen"
	"github.com/mattn/go-runewidth"
	"golang.org/x/crypto/ssh/terminal"
)

// BouncingBall testing
func BouncingBall() {

	const (
		printEmpty rune = ' '
		printBall  rune = '⚽'

		maxFrames = 1200
		velocity  = time.Millisecond * 50
	)

	width, height, err := terminal.GetSize(int(os.Stdout.Fd()))

	if err != nil {
		log.Fatal(err)
		return
	}

	var (
		py, px   int
		ppy, ppx int
		vy, vx   int = 1, 1
		char     rune
	)

	ballWidth := runewidth.RuneWidth(printBall)

	// adjust the width and height
	width /= ballWidth / 2
	// height-- // there is a 1 pixel border in my terminal

	board := make([][]bool, width, width)

	for row := range board {
		board[row] = make([]bool, height, height)
	}

	for i := 0; i < maxFrames; i++ {

		fmt.Print("\033[H")
		px += vx
		py += vy

		if px <= 0 || px >= width-1 {
			vx *= -1
		}
		if py <= 0 || py >= height-1 {
			vy *= -1
		}

		// Asigning new values and deleting the old ones
		board[px][py], board[ppx][ppy] = true, false

		// se guardan las pociciones actuales
		ppx, ppy = px, py

		// screen.Clear()

		bufCap := (width*2 + 1) * height // this is to optimize the the buffer
		buf := make([]rune, 0, bufCap)
		buf = buf[:0]

		for y := range board[0] {
			for x := range board {
				char = printEmpty

				if board[x][y] {
					char = printBall
				}
				buf = append(buf, char)

			}

			buf = append(buf, '\n')

		}

		screen.MoveTopLeft()
		// fmt.Print("\033[H\033[2J")
		// fmt.Print("\r")
		time.Sleep(velocity)
		fmt.Println(string(buf))

	}

}
