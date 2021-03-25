package ttgboard

import (
	"log"
	"math/rand"
	"time"

	"github.com/gucio321/tic-tac-go/ttgcommon"
)

func (t *TTG) canWin(player Letter) (i int, result bool) {
	for i := 0; i < t.width*t.height; i++ {
		if !t.board.isIndexFree(i) {
			continue
		}

		t.board.setIndexState(i, player)

		if t.isWinner(player) {
			return i, true
		}

		t.board.setIndexState(i, LetterNone)
	}

	return 0, false
}

func (t *TTG) getPCMove(letter Letter) (i int) {
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
		if t.board.isIndexFree(i) {
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
		if t.board.isIndexFree(i) {
			options = append(options, i)
		}
	}

	if options != nil {
		// nolint:gosec // it's ok
		result := options[rand.Intn(len(options))]

		return result
	}

	for _, i := range ttgcommon.GetMiddles() {
		if t.board.isIndexFree(i) {
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
