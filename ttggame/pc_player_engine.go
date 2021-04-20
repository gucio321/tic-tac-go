package ttggame

import (
	"log"
	"math/rand"
	"time"

	"github.com/gucio321/tic-tac-go/ttgcommon"
	"github.com/gucio321/tic-tac-go/ttggame/ttgletter"
)

func (t *TTG) canWin(player ttgletter.Letter) (i int, result bool) {
	for i := 0; i < t.width*t.height; i++ {
		if !t.board.IsIndexFree(i) {
			continue
		}

		board := t.board.Copy()

		board.SetIndexState(i, player)

		if t.isWinner(board, player) {
			return i, true
		}
	}

	return 0, false
}

func (t *TTG) getPCMove(letter ttgletter.Letter) (i int) {
	pcLetter := letter
	playerLetter := pcLetter.Opposite()

	var options []int = nil

	rand.Seed(time.Now().UnixNano())

	// attack: try to win
	if i, ok := t.canWin(pcLetter); ok {
		return i
	}

	// defense: check, if user can win
	if i, ok := t.canWin(playerLetter); ok {
		return i
	}

	for _, i := range ttgcommon.GetCorners(t.width, t.height) {
		if t.board.IsIndexFree(i) {
			options = append(options, i)
		}
	}

	if options != nil {
		// nolint:gosec // it's ok
		result := options[rand.Intn(len(options))]

		return result
	}

	// try to get center
	for _, i := range ttgcommon.GetCenter(t.width, t.height) {
		if t.board.IsIndexFree(i) {
			options = append(options, i)
		}
	}

	if options != nil {
		// nolint:gosec // it's ok
		result := options[rand.Intn(len(options))]

		return result
	}

	for _, i := range ttgcommon.GetMiddles(t.width, t.height) {
		if t.board.IsIndexFree(i) {
			options = append(options, i)
		}
	}

	if options != nil {
		// nolint:gosec // it's ok
		result := options[rand.Intn(len(options))]

		return result
	}

	log.Fatal("Cannot make move (board is full) and this fact wasn't detected")

	return 0
}
