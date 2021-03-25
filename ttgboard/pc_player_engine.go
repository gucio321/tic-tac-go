package ttgboard

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gucio321/tic-tac-go/ttgcommon"
)

func (t *TTT) canWin(player Letter) (x, y int, result bool) {
	for i := 0; i < t.height; i++ {
		for j := 0; j < t.width; j++ {
			if !t.board[i][j].IsNone() {
				fmt.Println("continue")
				continue
			}

			t.board[i][j].SetState(player)

			if t.isWinner(player) {
				fmt.Println("can win")
				return j, i, true
			}

			t.board[i][j].SetState(LetterNone)
		}
	}

	return 0, 0, false
}

func (t *TTT) getPCMove(letter Letter) (x, y int) {
	type option struct{ X, Y int }

	pcLetter := letter
	playerLetter := pcLetter.Oposite()

	var options []option = nil

	rand.Seed(time.Now().UnixNano())

	// attack: try to win
	if x, y, ok := t.canWin(pcLetter); ok {
		return x, y
	}

	// defense: check, if user can win
	if x, y, ok := t.canWin(playerLetter); ok {
		return x, y
	}

	for _, i := range ttgcommon.GetCorners(t.width, t.height) {
		cornerY, cornerX := ttgcommon.IntToCords(t.width, t.height, i)
		if t.board[cornerX][cornerY].IsNone() {
			options = append(options, option{cornerY, cornerX})
		}
	}

	if options != nil {
		// nolint:gosec // it's ok
		result := options[rand.Intn(len(options))]

		return result.X, result.Y
	}

	// try to get center
	for _, i := range ttgcommon.GetCenter(t.width, t.height) {
		centerY, centerX := ttgcommon.IntToCords(t.width, t.height, i)
		if t.board[centerX][centerY].IsNone() {
			options = append(options, option{centerY, centerX})
		}
	}

	if options != nil {
		// nolint:gosec // it's ok
		result := options[rand.Intn(len(options))]

		return result.X, result.Y
	}

	for _, cords := range ttgcommon.GetMiddles() {
		middleY, middleX := ttgcommon.IntToCords(t.width, t.height, cords)
		if t.board[middleX][middleY].IsNone() {
			options = append(options, option{middleY, middleX})
		}
	}

	if options != nil {
		// nolint:gosec // it's ok
		result := options[rand.Intn(len(options))]

		return result.X, result.Y
	}

	log.Fatal("Cannot make move (board is full) and this fact wasn't detected")

	return 0, 0
}
