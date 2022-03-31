// Package pcplayer provides methods for simple-AI logic
// used in Tic-Tac-Go for calculating PC-player's move.
package pcplayer

import (
	"math/rand"
	"time"

	"github.com/gucio321/tic-tac-go/pkg/core/board"
	"github.com/gucio321/tic-tac-go/pkg/core/board/letter"
)

// nolint:gochecknoinits // need to set rand seed
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
let's look at the board above: when we'll make our move at right side of X-chain (14)
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
	result = make([]int, 0)

	// nolint:gomnd // look a scheme above - in the second one, the chain is by 2 less than max
	minimalChainLen := gameBoard.ChainLength() - 2
	if minimalChainLen <= 0 {
		return
	}

	potentiallyAvailableChains := gameBoard.GetWinBoard(gameBoard.ChainLength() + 1)

searching:
	for _, potentialPlace := range potentiallyAvailableChains {
		if !gameBoard.IsIndexFree(potentialPlace[0]) || !gameBoard.IsIndexFree(potentialPlace[len(potentialPlace)-1]) {
			continue
		}

		var gaps []int

		for i := 1; i < len(potentialPlace)-1; i++ {
			switch gameBoard.GetIndexState(potentialPlace[i]) {
			case letter.LetterNone:
				gaps = append(gaps, potentialPlace[i])
			case player.Opposite(): // operation already blocked
				continue searching
			}
		}

		if len(gaps) == 1 {
			result = append(result, gaps...)
		}
	}

	return result
}

// GetPCMove calculates move for PC player on given board.
// Steps:
// - try to win
// - stop opponent from winning
// - try to win in 2-moves perspective
// - stop opponent from winning in 2-moves perspective
// - for current board and for each smaller board (w - 2 h - 2)
//   - take opposite (to any corner taken by pc) corner
//   - take opponent's opposite corner
//   - take center
//   - take random side
// nolint:gocognit,gocyclo // https://github.com/gucio321/tic-tac-go/issues/154
func GetPCMove(gameBoard *board.Board, pcLetter letter.Letter) (i int) {
	playerLetter := pcLetter.Opposite()

	// attack: try to win now
	if ok, indexes := canWin(gameBoard, pcLetter); ok {
		options := getAvailableOptions(gameBoard, indexes)

		if len(options) > 0 {
			return getRandomNumber(options)
		}
	}

	// defense: check, if user can win
	if ok, indexes := canWin(gameBoard, playerLetter); ok {
		options := getAvailableOptions(gameBoard, indexes)

		if len(options) > 0 {
			return getRandomNumber(options)
		}
	}

	options := getAvailableOptions(gameBoard, canWinTwoMoves(gameBoard, pcLetter))
	if len(options) > 0 {
		return getRandomNumber(options)
	}

	options = getAvailableOptions(gameBoard, canWinTwoMoves(gameBoard, playerLetter))
	if len(options) > 0 {
		return getRandomNumber(options)
	}

	corners := gameBoard.GetCorners()
	pcOppositeCorners := make([]int, 0)
	playerOppositeCorners := make([]int, 0)

	for _, i := range corners {
		if gameBoard.IsIndexFree(i) {
			options = append(options, i)

			continue
		}

		o := gameBoard.GetOppositeCorner(i)

		if !gameBoard.IsIndexFree(o) {
			continue
		}

		switch s := gameBoard.GetIndexState(i); s {
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

	if len(options) > 0 {
		return getRandomNumber(options)
	}

	// try to get center
	for _, i := range gameBoard.GetCenter() {
		if gameBoard.IsIndexFree(i) {
			options = append(options, i)
		}
	}

	if len(options) > 0 {
		return getRandomNumber(options)
	}

	for _, i := range gameBoard.GetSides() {
		if gameBoard.IsIndexFree(i) {
			options = append(options, i)
		}
	}

	if len(options) > 0 {
		return getRandomNumber(options)
	}

	const smallerBoard = 2
	if newW, newH := gameBoard.Width()-smallerBoard, gameBoard.Height()-smallerBoard; newW > 0 && newH > 0 {
		return gameBoard.ConvertIndex(newW, newH, GetPCMove(gameBoard.Cut(newW, newH), pcLetter))
	}

	panic("Tic-Tac-Go: pcplayer.GetPCMove(...): cannot determinate pc move - board is full")
}

func getRandomNumber(numbers []int) int {
	// nolint:gosec // it's ok
	result := numbers[rand.Intn(len(numbers))]

	return result
}

func getAvailableOptions(gameBoard *board.Board, candidates []int) (available []int) {
	available = make([]int, 0)

	for _, i := range candidates {
		if gameBoard.IsIndexFree(i) {
			available = append(available, i)
		}
	}

	return available
}
