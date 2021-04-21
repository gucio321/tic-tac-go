package ttggame

import (
	"fmt"
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

		if t.isWinner(board, t.chainLen, player) {
			return i, true
		}
	}

	return 0, false
}

/*
This method should find situations like that:
chain lenght = 4
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
func (t *TTG) canWinTwoMoves(player ttgletter.Letter) (result []int) {
	minimalChainLen := t.chainLen - 2
	// processing this doesn't make sense
	if minimalChainLen == 1 {
		return nil
	}

	b := ttgcommon.GetWinBoard(t.board.Width(), t.board.Height(), minimalChainLen)

	options := make([][]int, 0)

	for _, i := range b {
		line := 0

		for _, c := range i {
			if t.board.GetIndexState(c) == player {
				line++
			}
		}

		if line == minimalChainLen {
			options = append(options, i)
		}
	}

	for _, option := range options {
		/*row := make([]int, len(option))
		for n := range option {
			row[n] = ttgcommon.ConvertIndex(
				board.Width(),
				board.Height(),
				t.board.Width(),
				t.board.Width(),
				option[n],
			)
		}
		*/
		switch option[0] {
		case option[1] - 1: // horizontal
			separator := 1
			if option[0]-separator < 0 {
				continue
			}

			if !t.board.IsIndexFree(option[0] - separator) {
				continue
			}

			if option[len(option)-1]+separator >= t.width*t.height {
				continue
			}

			if !t.board.IsIndexFree(option[len(option)-1] + separator) {
				continue
			}

			// special case
			if (option[0]+1-1)%t.width == 0 { // will move back to the previous line
				continue
			}

			if (option[len(option)-1]+1)%t.width == 1 { // will move back to the previous line
				continue
			}

			if index := option[0] - 2*separator; index > 0 {
				if t.board.IsIndexFree(index) &&
					(index+1)%t.width != 0 {
					result = append(result, index+separator)
				}
			}
			if index := option[len(option)-1] + 2*separator; index < t.width*t.height {
				if t.board.IsIndexFree(index) &&
					(index+1)%t.width != 1 {
					result = append(result, index-separator)
				}
			}
			// index := option[len(option)-1]
			fmt.Println(t.width, ((option[len(option)-1]+2)+1)%t.width)
		case option[1] - t.board.Width(): // vertical
		case option[1] - t.board.Width() + 1: // slant - forward
		case option[1] - t.board.Width() - 1: // slant - backword
		default: // shouldn't be reached
			fmt.Println("Nothing batches: ", option)
		}
	}

	return result
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

	for _, i := range t.canWinTwoMoves(pcLetter) {
		if t.board.IsIndexFree(i) {
			options = append(options, i)
		}
	}

	if options != nil {
		// nolint:gosec // it's ok
		result := options[rand.Intn(len(options))]

		return result
	}

	for _, i := range t.canWinTwoMoves(playerLetter) {
		if t.board.IsIndexFree(i) {
			options = append(options, i)
			fmt.Println(i)
		}
	}

	if options != nil {
		// nolint:gosec // it's ok
		result := options[rand.Intn(len(options))]

		return result
	}

	const doubbleRow = 2

	nw := t.width
	nh := t.height

	for nw != 0 && nh != 0 {
		for _, i := range ttgcommon.GetCorners(nw, nh) {
			if idx := ttgcommon.ConvertIndex(nw, nh, t.width, t.height, i); t.board.IsIndexFree(idx) {
				options = append(options, idx)
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
			if idx := ttgcommon.ConvertIndex(nw, nh, t.width, t.height, i); t.board.IsIndexFree(idx) {
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
