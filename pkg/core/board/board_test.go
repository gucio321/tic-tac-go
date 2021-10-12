package board

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gucio321/tic-tac-go/pkg/core/board/letter"
)

func Test_Board_Create(t *testing.T) {
	correctBoard := &Board{
		board:    make([]letter.Letter, 9),
		width:    3,
		height:   3,
		chainLen: 3,
	}

	for i := range correctBoard.board {
		correctBoard.board[i] = letter.LetterNone
	}

	board := Create(3, 3, 3)

	assert.Equal(t, correctBoard, board, "Unexpected board created")
}

func Test_Board_getting_size(t *testing.T) {
	a := assert.New(t)
	w, h, c := 4, 3, 2
	board := Create(w, h, c)

	a.Equal(w, board.Width(), "unexpected board width")
	a.Equal(h, board.Height(), "unexpected board height")
	a.Equal(c, board.ChainLength(), "unexpected board height")
}

func Test_Board_SetIndexState(t *testing.T) {
	board := Create(3, 3, 3)

	board.SetIndexState(5, letter.LetterX)

	assert.Equal(t, letter.LetterX, board.board[5], "Index state was set incorrectly")
	assert.Panics(t, func() { board.SetIndexState(20, letter.LetterO) }, "Setting state of unegzisting index didn't panicked")
}

func Test_GetIndexState(t *testing.T) {
	board := Create(3, 3, 3)
	board.board[5] = letter.LetterX

	assert.Equal(t, letter.LetterX, board.GetIndexState(5), "Unexpected index state")
	assert.Panics(t, func() { board.GetIndexState(20) }, "getting state of unegzisging index didn't panicked")
}

func Test_isIndexFree(t *testing.T) {
	a := assert.New(t)
	board := Create(3, 3, 3)

	board.board[5] = letter.LetterX

	a.False(board.IsIndexFree(5), "IsIndexFree returned unexpected value")
	a.True(board.IsIndexFree(4), "IsIndexFree returned unexpected value")
	a.Panics(func() { board.IsIndexFree(20) }, "IsIndexFree returned unexpected value")
}

func Test_Copy(t *testing.T) {
	board := Create(3, 3, 3)
	board.SetIndexState(4, letter.LetterX)
	newBoard := board.Copy()

	assert.Equal(t, board, newBoard, "unexpected board copied")
}

func Test_Cut(t *testing.T) {
	a := assert.New(t)
	board := Create(3, 3, 3)
	board.SetIndexState(4, letter.LetterX)
	result := board.Cut(1, 1)

	a.Equal(1, len(result.board), "wrong board cut")
	a.Equal(letter.LetterX, result.board[0], "wrong board cut")

	a.Panics(func() { board.Cut(20, 20) }, "cutting larger board from smaller didn't panicked")
}

func Test_IsBoardFull(t *testing.T) {
	board := Create(2, 2, 2)
	if board.IsBoardFull() {
		t.Fatal("unexbected value returned by isBoardFull method")
	}

	board.SetIndexState(0, letter.LetterX)
	board.SetIndexState(1, letter.LetterO)
	board.SetIndexState(2, letter.LetterO)
	board.SetIndexState(3, letter.LetterX)

	if !board.IsBoardFull() {
		t.Fatal("unexbected value returned by isBoardFull method")
	}
}

func Test_IntToCords(t *testing.T) {
	// standard 3x3 board
	/*
		+---+---+---+
		| 0 | 1 | 2 |
		+---+---+---+
		| 3 | 4 | 5 |
		+---+---+---+
		| 6 | 7 | 8 |
		+---+---+---+
	*/
	w, h, c := 3, 3, 3
	b := Create(w, h, c)
	// so index 3 should have y = 1, x = 0
	i := 3
	x, y := b.IntToCords(i)

	if y != 1 || x != 0 {
		t.Fatalf("IntToCords(%d, %d, %d) returned unexpected values x: %d, y: %d", w, h, i, x, y)
	}

	// index 7 should have y = 2, x = 1
	i = 7
	x, y = b.IntToCords(i)

	if y != 2 || x != 1 {
		t.Fatalf("IntToCords(%d, %d, %d) returned unexpected values x: %d, y: %d", w, h, i, x, y)
	}
}
