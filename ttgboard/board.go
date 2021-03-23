package ttgboard

import (
	"bufio"
	"os"

	"github.com/gucio321/tic-tac-go/ttgcommon"
)

// IdxState represents index's state
type IdxState int

func (i IdxState) String() string {
	switch i {
	case IdxNone:
		return " "
	case IdxX:
		return "X"
	case IdxO:
		return "O"
	}

	return "?"
}

// Indt.player2.name + "s
const (
	IdxNone IdxState = iota
	IdxX
	IdxO
)

// PlayerType represents players' types
type PlayerType int

// player types
const (
	PlayerPC PlayerType = iota
	PlayerPerson
)

func (p PlayerType) String() string {
	switch p {
	case PlayerPC:
		return "PC"
	case PlayerPerson:
		return "Player"
	}

	return "?"
}

// BoardIndex represents board index
type BoardIndex struct {
	state IdxState
}

func newIndex() *BoardIndex {
	result := &BoardIndex{
		state: IdxNone,
	}

	return result
}

// SetState sets index state
func (b *BoardIndex) SetState(state IdxState) {
	b.state = state
}

func (b *BoardIndex) String() string {
	switch b.state {
	case IdxNone:
		return " "
	case IdxX:
		return "X"
	case IdxO:
		return "O"
	}

	// should not be reached
	return "?"
}

// IsFree return's true if index is free
func (b BoardIndex) IsFree() bool {
	return b.state == IdxNone
}

type player struct {
	name       string
	playerType PlayerType
	letter     IdxState
	moveCb     func() (x, y int)
}

func newPlayer(t PlayerType, letter IdxState, cb func() (x, y int)) *player {
	result := &player{
		playerType: t,
		letter:     letter,
		moveCb:     cb,
		name:       t.String() + " " + letter.String(),
	}

	return result
}

// TTT represents TicTacToe game
type TTT struct {
	board   [][]*BoardIndex
	reader  *bufio.Reader
	player1 *player
	player2 *player
	width,
	height int
}

// NewTTT creates a ne TTT
func NewTTT(w, h int, player1Type, player2Type PlayerType) *TTT {
	result := &TTT{
		board:  make([][]*BoardIndex, h),
		reader: bufio.NewReader(os.Stdin),
		width:  w,
		height: h,
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
	b := ttgcommon.GetWinBoard(t.width, t.height, 3)

	for _, i := range b {
		c1, c2, c3 := i[0], i[1], i[2]
		x1, y1 := ttgcommon.IntToCords(t.width, t.height, c1)
		x2, y2 := ttgcommon.IntToCords(t.width, t.height, c2)
		x3, y3 := ttgcommon.IntToCords(t.width, t.height, c3)

		if (t.board[y1][x1].state == player) &&
			t.board[y2][x2].state == player &&
			t.board[y3][x3].state == player {
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
