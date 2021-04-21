package ttggame

import (
	"log"
	"math/rand"
	"time"

	"github.com/gucio321/tic-tac-go/ttgcommon"
	"github.com/gucio321/tic-tac-go/ttggame/ttgletter"
)

// func (t *TTG) canWinOneMove(player ttgletter.Letter) (i int, result bool) {
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

func (t *TTG) canWinTwoMoves(player ttgletter.Letter) (i int, result bool) {
	for i := 0; i < t.width*t.height; i++ {
		/*if t.board.GetIndexState(i) != player {
			continue
		}

		if ttgcommon.IsEdgeIndex(t.width, t.height, i) {
			continue
		}*/

		// board := t.board.Copy().Cut(t.width-2, t.height-2)

		// minimalChainLen := t.chainLen - 2
		// ...

		/*if t.isWinner(board, player) {
			return i, true
		}*/
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

	const doubbleRow = 2

	nw := t.width
	nh := t.height

	for nw != 0 && nh != 0 {
		for _, i := range ttgcommon.GetCorners(nw, nh) {
			if t.board.IsIndexFree(ttgcommon.ConvertIndex(t.width, t.height, nw, nh, i)) {
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

		for _, i := range ttgcommon.GetMiddles(nw, nh) {
			if t.board.IsIndexFree(ttgcommon.ConvertIndex(t.width, t.height, nw, nh, i)) {
				options = append(options, i)
			}
		}

		if options != nil {
			// nolint:gosec // it's ok
			result := options[rand.Intn(len(options))]

			return result
		}

		nw -= doubbleRow
		nh -= doubbleRow
	}

	log.Fatal("Cannot make move (board is full) and this fact wasn't detected")

	return 0
}
