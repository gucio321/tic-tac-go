// Package pcplayer provides megods for simple-AI logic
// used in Tic-Tac-Go for calculating PC-player's move.
package pcplayer

import (
	"math/rand"
	"time"

	"github.com/gucio321/tic-tac-go/pkg/core/board"
	"github.com/gucio321/tic-tac-go/pkg/core/board/letter"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func canWin(baseBoard *board.Board, player letter.Letter) (canWin bool, results []int) {
	results = make([]int, 0)
	for i := 0; i < baseBoard.Width()*baseBoard.Height(); i++ {
		if !baseBoard.IsIndexFree(i) {
			continue
		}

		fictionBoard := baseBoard.Copy()

		fictionBoard.SetIndexState(i, player)

		if ok, _ := fictionBoard.IsWinner(player); ok {
			results = append(results, i)
		}
	}

	return len(results) > 0, results
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
func canWinTwoMoves(gameBoard *board.Board, player letter.Letter) (result []int) {
	// nolint:gomnd // look a scheme above - in the second one, the chain is by 2 less than max
	minimalChainLen := gameBoard.ChainLength() - 2
	if minimalChainLen < 2 { // nolint:gomnd // processing this values doesn't make sense with chain smaller than 3
		return nil
	}

	b := gameBoard.GetWinBoard(minimalChainLen)
	options := make([][]int, 0)

	for _, i := range b {
		line := 0

		for _, c := range i {
			if gameBoard.GetIndexState(c) == player {
				line++
			}
		}

		if line == minimalChainLen {
			options = append(options, i)
		}
	}

	b = gameBoard.GetWinBoard(gameBoard.ChainLength() + 1)
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
func GetPCMove(gameBoard *board.Board, pcLetter letter.Letter) (i int) {
	playerLetter := pcLetter.Opposite()

	var options []int

	// attack: try to win now
	if ok, indexes := canWin(gameBoard, pcLetter); ok {
		for _, i := range indexes {
			if gameBoard.IsIndexFree(i) {
				options = append(options, i)
			}
		}

		if options != nil {
			return getRandomNumber(options)
		}
	}

	// defense: check, if user can win
	if ok, indexes := canWin(gameBoard, playerLetter); ok {
		for _, i := range indexes {
			if gameBoard.IsIndexFree(i) {
				options = append(options, i)
			}
		}

		if options != nil {
			return getRandomNumber(options)
		}
	}

	for _, i := range canWinTwoMoves(gameBoard, pcLetter) {
		if gameBoard.IsIndexFree(i) {
			options = append(options, i)
		}
	}

	if options != nil {
		return getRandomNumber(options)
	}

	for _, i := range canWinTwoMoves(gameBoard, playerLetter) {
		if gameBoard.IsIndexFree(i) {
			options = append(options, i)
		}
	}

	if options != nil {
		return getRandomNumber(options)
	}

	const doubbleRow = 2

	nw := gameBoard.Width()
	nh := gameBoard.Height()

	for nw != 0 && nh != 0 {
		fictionBoard := board.Create(nw, nh, 0)
		corners := fictionBoard.GetCorners()
		pcOppositeCorners := make([]int, 0)
		playerOppositeCorners := make([]int, 0)

		for _, i := range corners {
			idx := gameBoard.ConvertIndex(nw, nh, i)
			if gameBoard.IsIndexFree(idx) {
				options = append(options, idx)

				continue
			}

			o := gameBoard.ConvertIndex(nw, nh, fictionBoard.GetOppositeCorner(i))

			if !gameBoard.IsIndexFree(o) {
				continue
			}

			switch s := gameBoard.GetIndexState(idx); s {
			case pcLetter:
				pcOppositeCorners = append(pcOppositeCorners, o)
			case playerLetter:
				playerOppositeCorners = append(playerOppositeCorners, o)
			}
		}

		if len(pcOppositeCorners) != 0 {
			return getRandomNumber(pcOppositeCorners)
		}

		if len(playerOppositeCorners) != 0 {
			return getRandomNumber(playerOppositeCorners)
		}

		if options != nil {
			return getRandomNumber(options)
		}

		// try to get center
		for _, i := range gameBoard.GetCenter() {
			if gameBoard.IsIndexFree(i) {
				options = append(options, i)
			}
		}

		if options != nil {
			return getRandomNumber(options)
		}

		for _, i := range fictionBoard.GetSides() {
			if idx := gameBoard.ConvertIndex(nw, nh, i); gameBoard.IsIndexFree(idx) {
				options = append(options, idx)
			}
		}

		if options != nil {
			return getRandomNumber(options)
		}

		nw -= doubbleRow
		nh -= doubbleRow
	}

	panic("Tic-Tac-Go: pcplayer.GetPCMove(...): cannot determinate pc move - board is full")
}

func getRandomNumber(numbers []int) int {
	// nolint:gosec // it's ok
	result := numbers[rand.Intn(len(numbers))]

	return result
}
