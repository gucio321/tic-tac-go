package ttgboard

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

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
// nolint:golint // this name is ok
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
	board   [ttgcommon.BoardH][ttgcommon.BoardW]*BoardIndex
	reader  *bufio.Reader
	player1 *player
	player2 *player
}

// NewTTT creates a ne TTT
func NewTTT(player1Type, player2Type PlayerType) *TTT {
	result := &TTT{
		reader: bufio.NewReader(os.Stdin),
	}

	switch player1Type {
	case PlayerPC:
		result.player1 = newPlayer(player1Type, IdxX, result.getPCMove)
	case PlayerPerson:
		result.player1 = newPlayer(player1Type, IdxX, result.getPlayerMove)
	}

	switch player2Type {
	case PlayerPC:
		result.player2 = newPlayer(player2Type, IdxO, result.getPCMove)
	case PlayerPerson:
		result.player2 = newPlayer(player2Type, IdxO, result.getPlayerMove)
	}

	var board [ttgcommon.BoardH][ttgcommon.BoardW]*BoardIndex

	for i := 0; i < ttgcommon.BoardH; i++ {
		for j := 0; j < ttgcommon.BoardW; j++ {
			board[i][j] = newIndex()
		}
	}

	result.board = board

	return result
}

func (t *TTT) printSeparator() {
	sep := "+"
	for i := 0; i < ttgcommon.BoardW; i++ {
		sep += "---+"
	}

	fmt.Println(sep)
}

func (t *TTT) printBoard() {
	ttgcommon.Clear()

	t.printSeparator()

	for i := 0; i < ttgcommon.BoardH; i++ {
		line := "| "
		for j := 0; j < ttgcommon.BoardW; j++ {
			line += t.board[i][j].String()
			line += " | "
		}

		fmt.Println(line)
		t.printSeparator()
	}
}

func (t *TTT) getPlayerMove() (x, y int) {
	for {
		fmt.Print("Enter your move (1-9): ")

		text, err := t.reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		if text == "" {
			fmt.Println("please enter number from 1 to 9")

			continue
		}

		text = strings.ReplaceAll(text, "\n", "")

		num, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("invalid index number")

			continue
		}

		if num <= 0 || num > ttgcommon.BoardW*ttgcommon.BoardH {
			fmt.Println("You must enter number from 1 to 9")
		}

		num--

		x, y = ttgcommon.IntToCords(num)

		if t.board[y][x].IsFree() {
			return
		}

		fmt.Println("This index is busy")
	}
}

func (t *TTT) isWinner(player IdxState) bool {
	b := ttgcommon.GetWinBoard()

	for _, i := range b {
		c1, c2, c3 := i[0], i[1], i[2]
		x1, y1 := ttgcommon.IntToCords(c1)
		x2, y2 := ttgcommon.IntToCords(c2)
		x3, y3 := ttgcommon.IntToCords(c3)

		if (t.board[y1][x1].state == player) &&
			t.board[y2][x2].state == player &&
			t.board[y3][x3].state == player {
			return true
		}
	}

	return false
}

func (t *TTT) canWin(player IdxState) (x, y int, result bool) {
	for i := 0; i < ttgcommon.BoardH; i++ {
		for j := 0; j < ttgcommon.BoardW; j++ {
			if !t.board[i][j].IsFree() {
				continue
			}

			t.board[i][j].state = player

			if t.isWinner(player) {
				return j, i, true
			}

			t.board[i][j].state = IdxNone
		}
	}

	return 0, 0, false
}

func (t *TTT) isBoardFull() bool {
	for i := 0; i < ttgcommon.BoardH; i++ {
		for j := 0; j < ttgcommon.BoardW; j++ {
			if t.board[i][j].IsFree() {
				return false
			}
		}
	}

	return true
}

func (t *TTT) getPCMove() (x, y int) {
	type option struct{ X, Y int }

	var options []option = nil

	rand.Seed(time.Now().UnixNano())

	// attack: try to win
	if x, y, ok := t.canWin(IdxO); ok {
		return x, y
	}

	// defense: check, if user can win
	if x, y, ok := t.canWin(IdxX); ok {
		return x, y
	}

	for _, i := range ttgcommon.GetCorners() {
		if t.board[i.Y][i.X].IsFree() {
			options = append(options, option{i.X, i.Y})
		}
	}

	if options != nil {
		// nolint:gosec // it's ok
		result := options[rand.Intn(len(options))]

		return result.X, result.Y
	}

	// try to get center
	if t.board[1][1].IsFree() {
		return 1, 1
	}

	for _, i := range ttgcommon.GetMiddles() {
		if t.board[i.Y][i.X].IsFree() {
			options = append(options, option{i.X, i.Y})
		}
	}

	if options != nil {
		// nolint:gosec // it's ok
		result := options[rand.Intn(len(options))]

		return result.X, result.Y
	}

	log.Fatal("Cannot make move (board is full)")

	return 0, 0
}

func (t *TTT) move(x, y int, letter IdxState) {
	t.board[y][x].SetState(letter)
}

func (t *TTT) pressAnyKey() {
	fmt.Print("\nPress any key to continue...")

	_, _ = t.reader.ReadString('\n')
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
			fmt.Println(t.player1.name + " won")
			t.pressAnyKey()

			break
		} else if t.isBoardFull() {
			ttgcommon.Clear()
			t.printBoard()
			fmt.Println("DRAW")
			t.pressAnyKey()

			break
		}

		t.printBoard()
		x, y = t.player2.moveCb()
		t.move(x, y, t.player2.letter)

		if t.isWinner(t.player2.letter) {
			ttgcommon.Clear()
			t.printBoard()
			fmt.Println(t.player2.name + " won")
			t.pressAnyKey()

			break
		} else if t.isBoardFull() {
			ttgcommon.Clear()
			t.printBoard()
			fmt.Println("DRAW")
			t.pressAnyKey()

			break
		}
	}
}
