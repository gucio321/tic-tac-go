package ttgboard

import (
	"bufio"
	"os"

	"github.com/gucio321/tic-tac-go/ttgcommon"
)

// TTT represents TicTacToe game
type TTT struct {
	board   [][]*BoardIndex
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
		board:    make([][]*BoardIndex, h),
		reader:   bufio.NewReader(os.Stdin),
		width:    w,
		height:   h,
		chainLen: chainLen,
	}

	for i := 0; i < h; i++ {
		result.board[i] = make([]*BoardIndex, w)
		for j := 0; j < w; j++ {
			result.board[i][j] = newIndex()
		}
	}

	switch player1Type {
	case PlayerPC:
		result.player1 = newPlayer(player1Type, IdxX,
			func() (x, y int) {
				x, y = result.getPCMove(IdxX)
				return x, y
			},
		)
	case PlayerPerson:
		result.player1 = newPlayer(player1Type, IdxX, result.getPlayerMove)
	}

	switch player2Type {
	case PlayerPC:
		result.player2 = newPlayer(player2Type, IdxO,
			func() (x, y int) {
				x, y = result.getPCMove(IdxO)
				return x, y
			},
		)
	case PlayerPerson:
		result.player2 = newPlayer(player2Type, IdxO, result.getPlayerMove)
	}

	return result
}

func (t *TTT) isWinner(player IdxState) bool {
	b := ttgcommon.GetWinBoard(t.width, t.height, t.chainLen)

	for _, i := range b {
		indexes := make([]struct{ cords, x, y int }, t.chainLen)
		for c := 0; c < t.chainLen; c++ {
			indexes[c].cords = i[c]
			indexes[c].x, indexes[c].y = ttgcommon.IntToCords(t.width, t.height, i[c])
		}

		line := 0

		for _, c := range indexes {
			if t.board[c.y][c.x].state == player {
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
	for i := 0; i < t.height; i++ {
		for j := 0; j < t.width; j++ {
			if t.board[i][j].IsFree() {
				return false
			}
		}
	}

	return true
}

func (t *TTT) move(x, y int, letter IdxState) {
	t.board[y][x].SetState(letter)
}

// Run runs the game
func (t *TTT) Run() {
	var x, y int

	for {
		t.printBoard()
		x, y = t.player1.moveCb()
		t.move(x, y, t.player1.letter)

		if t.isWinner(t.player1.letter) {
			ttgcommon.Clear()
			t.printBoard()
			t.println(t.player1.name + " won")
			t.pressAnyKeyPrompt()

			break
		} else if t.isBoardFull() {
			ttgcommon.Clear()
			t.printBoard()
			t.println("DRAW")
			t.pressAnyKeyPrompt()

			break
		}

		t.printBoard()
		x, y = t.player2.moveCb()
		t.move(x, y, t.player2.letter)

		if t.isWinner(t.player2.letter) {
			ttgcommon.Clear()
			t.printBoard()
			t.println(t.player2.name + " won")
			t.pressAnyKeyPrompt()

			break
		} else if t.isBoardFull() {
			ttgcommon.Clear()
			t.printBoard()
			t.println("DRAW")
			t.pressAnyKeyPrompt()

			break
		}
	}
}
