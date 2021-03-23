package ttgboard

import (
	"log"
	"math/rand"
	"time"

	"github.com/gucio321/tic-tac-go/ttgcommon"
)

func (t *TTT) canWin(player IdxState) (x, y int, result bool) {
	for i := 0; i < t.height; i++ {
		for j := 0; j < t.width; j++ {
			if !t.board[i][j].IsFree() {
				continue
			}

			t.board[i][j].state = player

			if t.isWinner(player) {
				return j, i, true
			}

			t.board[i][j].state = IdxNone
		}
	}

	return 0, 0, false
}

func (t *TTT) getPCMove(letter IdxState) (x, y int) {
	type option struct{ X, Y int }

	pcLetter := letter

	var playerLetter IdxState

	switch pcLetter {
	case IdxX:
		playerLetter = IdxO
	case IdxO:
		playerLetter = IdxX
	}

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
		if t.board[cornerX][cornerY].IsFree() {
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
		if t.board[centerX][centerY].IsFree() {
			options = append(options, option{centerY, centerX})
		}
	}

	if options != nil {
		// nolint:gosec // it's ok
		result := options[rand.Intn(len(options))]

		return result.X, result.Y
	}

	for _, i := range ttgcommon.GetMiddles() {
		if t.board[i.Y][i.X].IsFree() {
			options = append(options, option{i.X, i.Y})
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
