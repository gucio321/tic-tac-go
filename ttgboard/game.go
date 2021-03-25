package ttgboard

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gucio321/tic-tac-go/ttgcommon"
)

// TTT represents TicTacToe game
type TTT struct {
	board   *board
	reader  *bufio.Reader
	player1 *player
	player2 *player
	width,
	height int
	chainLen int
}

// NewTTT creates a ne TTT
func NewTTT(w, h, chainLen int, player1Type, player2Type PlayerType) *TTT {
	result := &TTT{
		board:    newBoard(w, h),
		reader:   bufio.NewReader(os.Stdin),
		width:    w,
		height:   h,
		chainLen: chainLen,
	}

	player1Letter := LetterX
	player2Letter := LetterO

	switch player1Type {
	case PlayerPC:
		result.player1 = newPlayer(player1Type, player1Letter,
			func() (i int) {
				i = result.getPCMove(player1Letter)
				return i
			},
		)
	case PlayerPerson:
		result.player1 = newPlayer(player1Type, player1Letter, result.getPlayerMove)
	}

	switch player2Type {
	case PlayerPC:
		result.player2 = newPlayer(player2Type, player2Letter,
			func() (i int) {
				i = result.getPCMove(player2Letter)
				return i
			},
		)
	case PlayerPerson:
		result.player2 = newPlayer(player2Type, LetterO, result.getPlayerMove)
	}

	return result
}

func (t *TTT) isWinner(player Letter) bool {
	b := ttgcommon.GetWinBoard(t.width, t.height, t.chainLen)

	for _, i := range b {
		line := 0

		for _, c := range i {
			if t.board.getIndexState(c) == player {
				line++
			}
		}

		if line == t.chainLen {
			return true
		}
	}

	return false
}

func (t *TTT) isBoardFull() bool {
	for i := 0; i < t.width*t.height; i++ {
		if t.board.isIndexFree(i) {
			return false
		}
	}

	return true
}

// Run runs the game
func (t *TTT) Run() {
	for {
		t.printBoard()
		i := t.player1.moveCb()
		t.board.setIndexState(i, t.player1.letter)

		if t.isWinner(t.player1.letter) {
			ttgcommon.Clear()
			t.printBoard()
			fmt.Println(t.player1.name + " won")
			t.pressAnyKeyPrompt()

			break
		} else if t.isBoardFull() {
			ttgcommon.Clear()
			t.printBoard()
			fmt.Println("DRAW")
			t.pressAnyKeyPrompt()

			break
		}

		t.printBoard()
		i = t.player2.moveCb()
		t.board.setIndexState(i, t.player2.letter)

		if t.isWinner(t.player2.letter) {
			ttgcommon.Clear()
			t.printBoard()
			fmt.Println(t.player2.name + " won")
			t.pressAnyKeyPrompt()

			break
		} else if t.isBoardFull() {
			ttgcommon.Clear()
			t.printBoard()
			fmt.Println("DRAW")
			t.pressAnyKeyPrompt()

			break
		}
	}
}
