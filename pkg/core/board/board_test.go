package board

import (
	"testing"

	"github.com/stretchr/testify/assert"

)

func Test_Create(t *testing.T) {
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

func Test_Width(t *testing.T) {
	w, h := 3, 3
	board := Create(w, h, 3)

	if board.Width() != w {
		t.Fatal("Unexpected with returned")
	}

	if board.Height() != h {
		t.Fatal("Unexpected board height returned")
	}
}

func Test_setIndexState(t *testing.T) {
	board := Create(3, 3, 3)

	board.SetIndexState(5, letter.LetterX)

	if board.board[5] != letter.LetterX {
		t.Fatal("unexpected index was set by board.setIndexState")
	}
}

func Test_getIndexState(t *testing.T) {
	board := Create(3, 3, 3)
	board.board[5] = letter.LetterX

	if l := board.GetIndexState(5); l != letter.LetterX {
		t.Fatal("unexpected index was returned by board.getIndexState")
	}
}

func Test_isIndexFree(t *testing.T) {
	board := Create(3, 3, 3)

	board.board[5] = letter.LetterX

	if board.IsIndexFree(5) {
		t.Fatal("isIndexFree returned unexpected value")
	}

	if !board.IsIndexFree(4) {
		t.Fatal("isIndexFree returned unexpected value")
	}
}

func Test_Copy(t *testing.T) {
	board := Create(3, 3, 3)
	board.SetIndexState(4, letter.LetterX)
	newBoard := board.Copy()

	assert.Equal(t, board, newBoard, "unexpected board copied")
}

func Test_Cut(t *testing.T) {
	board := Create(3, 3, 3)
	board.SetIndexState(4, letter.LetterX)
	result := board.Cut(1, 1)

	if len(result.board) != 1 {
		t.Fatal("unexpected board cut")
	}

	if result.board[0] != letter.LetterX {
		t.Fatal("unexpected board cut")
	}
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
