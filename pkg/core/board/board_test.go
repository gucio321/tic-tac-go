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
	const chainLen = 3

	tests := []struct {
		name                 string
		w, h                 int
		source               int
		expectedX, expectedY int
	}{
		{"Test 1", 3, 3, 3, 0, 1},
		{"Test 2", 3, 3, 7, 1, 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			b := Create(test.w, test.h, chainLen)
			x, y := b.IntToCords(test.source)

			assert.Equal(tt, test.expectedX, x, "unexpected index converted")
			assert.Equal(tt, test.expectedY, y, "unexpected index converted")
		})
	}
}

func Test_IntToCords_incorrect_cords(t *testing.T) {
	assert.Panics(t, func() {
		Create(2, 2, 2).IntToCords(20)
	}, "calling IntToCords with too large index didn't panicked")
}

func Test_CordsToInt(t *testing.T) {
	tests := []struct {
		name             string
		w, h             int
		sourceX, sourceY int
		expected         int
	}{
		{"Test 1", 3, 3, 0, 0, 0},
		{"Test 2", 3, 3, 1, 2, 7},
		{"Test 3", 3, 4, 1, 2, 7},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			assert.Equal(tt,
				Create(test.w, test.h, 1).CordsToInt(test.sourceX, test.sourceY),
				test.expected, "unexpected result")
		})
	}
}

func Test_CordsToInt_incorrect_cords(t *testing.T) {
	assert.Panics(t, func() {
		Create(2, 2, 2).CordsToInt(20, 20)
	}, "calling IntToCords with too large index didn't panicked")
}
