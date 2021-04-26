// Package ttgpcplayer provides megods for simple-AI logic
// used in Tic-Tac-Go for calculating PC-player's move.
package ttgpcplayer

import (
	"log"
	"math/rand"
	"time"

	"github.com/gucio321/tic-tac-go/ttgcommon"
	"github.com/gucio321/tic-tac-go/ttgcore/ttgboard"
	"github.com/gucio321/tic-tac-go/ttgcore/ttgletter"
)

func canWin(baseBoard *ttgboard.Board, player ttgletter.Letter) (i int, result bool) {
	for i := 0; i < baseBoard.Width()*baseBoard.Height(); i++ {
		if !baseBoard.IsIndexFree(i) {
			continue
		}

		board := baseBoard.Copy()

		board.SetIndexState(i, player)

		if ok, _ := board.IsWinner(board.ChainLength(), player); ok {
			return i, true
		}
	}

	return 0, false
}

/*
This method should find situations like that:
chain length = 4
+---+---+---+---+---+
|   |   | o |   |   |
+---+---+---+---+---+
|   |   |   |   | o |
+---+---+---+---+---+
|   | x | x |   |   |
+---+---+---+---+---+
|   |   |   |   |   |
+---+---+---+---+---+
|   |   |   | o |   |
+---+---+---+---+---+
let's look at bove board: when we'll make our move at left side of X-chain (14)
the O-player will not be able to keep X from winning.
+---+---+---+---+---+
|   |   | o |   |   |
+---+---+---+---+---+
|   |   |   |   | o |
+---+---+---+---+---+
|   | x | x | X |   |
+---+---+---+---+---+
|   |   |   |   |   |
+---+---+---+---+---+
|   |   |   | o |   |
+---+---+---+---+---+
O-player lost.
*/
func canWinTwoMoves(board *ttgboard.Board, player ttgletter.Letter) (result []int) {
	// nolint:gomnd // look a scheme above - in the second one, the chain is by 2 less than max
	minimalChainLen := board.ChainLength() - 2
	b := ttgcommon.GetWinBoard(board.Width(), board.Height(), minimalChainLen)
	options := make([][]int, 0)

	for _, i := range b {
		line := 0

		for _, c := range i {
			if board.GetIndexState(c) == player {
				line++
			}
		}

		if line == minimalChainLen {
			options = append(options, i)
		}
	}

	b = ttgcommon.GetWinBoard(board.Width(), board.Height(), board.ChainLength()+1)
	for _, i := range b {
		for _, o := range options {
			if i[1] == o[0] && i[2] == o[1] {
				result = append(result, i[len(i)-2])
			} else if i[2] == o[0] && i[3] == o[1] {
				result = append(result, i[1])
			}
		}
	}

	return result
}

// GetPCMove calculates move for PC player on given board
// nolint:gocognit,gocyclo // it is ok
func GetPCMove(board *ttgboard.Board, letter ttgletter.Letter) (i int) {
	pcLetter := letter
	playerLetter := pcLetter.Opposite()

	var options []int = nil

	rand.Seed(time.Now().UnixNano())

	// attack: try to win
	if i, ok := canWin(board, pcLetter); ok {
		return i
	}

	// defense: check, if user can win
	if i, ok := canWin(board, playerLetter); ok {
		return i
	}

	for _, i := range canWinTwoMoves(board, pcLetter) {
		if board.IsIndexFree(i) {
			options = append(options, i)
		}
	}

	if options != nil {
		// nolint:gosec // it's ok
		result := options[rand.Intn(len(options))]

		return result
	}

	for _, i := range canWinTwoMoves(board, playerLetter) {
		if board.IsIndexFree(i) {
			options = append(options, i)
		}
	}

	if options != nil {
		// nolint:gosec // it's ok
		result := options[rand.Intn(len(options))]

		return result
	}

	const doubbleRow = 2

	nw := board.Width()
	nh := board.Height()

	for nw != 0 && nh != 0 {
		for _, i := range ttgcommon.GetCorners(nw, nh) {
			if idx := ttgcommon.ConvertIndex(nw, nh, board.Width(), board.Height(), i); board.IsIndexFree(idx) {
				options = append(options, idx)
			}
		}

		if options != nil {
			// nolint:gosec // it's ok
			result := options[rand.Intn(len(options))]

			return result
		}

		// try to get center
		for _, i := range ttgcommon.GetCenter(board.Width(), board.Height()) {
			if board.IsIndexFree(i) {
				options = append(options, i)
			}
		}

		if options != nil {
			// nolint:gosec // it's ok
			result := options[rand.Intn(len(options))]

			return result
		}

		for _, i := range ttgcommon.GetMiddles(nw, nh) {
			if idx := ttgcommon.ConvertIndex(nw, nh, board.Width(), board.Height(), i); board.IsIndexFree(idx) {
				options = append(options, idx)
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
