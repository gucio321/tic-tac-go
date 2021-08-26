package ttgboard

import (
	"testing"

	"github.com/gucio321/tic-tac-go/pkg/core/ttgletter"
)

func Test_Create(t *testing.T) {
	correctBoard := Board{
		board:  make([]*ttgletter.Letter, 9),
		width:  3,
		height: 3,
	}
	for i := range correctBoard.board {
		correctBoard.board[i] = ttgletter.Create()
	}

	board := Create(3, 3, 3)

	if len(board.board) != len(correctBoard.board) {
		t.Fatal("Invalid board created")
	}

	for i := range board.board {
		if *board.board[i] != *correctBoard.board[i] {
			t.Fatal("Invalid board created")
		}
	}
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

	board.SetIndexState(5, ttgletter.LetterX)

	if *board.board[5] != ttgletter.LetterX {
		t.Fatal("unexpected index was set by board.setIndexState")
	}
}

func Test_getIndexState(t *testing.T) {
	board := Create(3, 3, 3)
	*board.board[5] = ttgletter.LetterX

	if l := board.GetIndexState(5); l != ttgletter.LetterX {
		t.Fatal("unexpected index was returned by board.getIndexState")
	}
}

func Test_isIndexFree(t *testing.T) {
	board := Create(3, 3, 3)

	*board.board[5] = ttgletter.LetterX

	if board.IsIndexFree(5) {
		t.Fatal("isIndexFree returned unexpected value")
	}

	if !board.IsIndexFree(4) {
		t.Fatal("isIndexFree returned unexpected value")
	}
}

func Test_Copy(t *testing.T) {
	board := Create(3, 3, 3)
	board.SetIndexState(4, ttgletter.LetterX)
	newBoard := board.Copy()

	if len(board.board) != len(newBoard.board) {
		t.Fatal("Unexpected board copied")
	}

	for i := range board.board {
		if *board.board[i] != *newBoard.board[i] {
			t.Fatal("Unexpected board copied")
		}
	}
}

func Test_Cut(t *testing.T) {
	board := Create(3, 3, 3)
	board.SetIndexState(4, ttgletter.LetterX)
	result := board.Cut(1, 1)

	if len(result.board) != 1 {
		t.Fatal("unexpected board cut")
	}

	if *result.board[0] != ttgletter.LetterX {
		t.Fatal("unexpected board cut")
	}
}

func Test_IsBoardFull(t *testing.T) {
	board := Create(2, 2, 2)
	if board.IsBoardFull() {
		t.Fatal("unexbected value returned by isBoardFull method")
	}

	board.SetIndexState(0, ttgletter.LetterX)
	board.SetIndexState(1, ttgletter.LetterO)
	board.SetIndexState(2, ttgletter.LetterO)
	board.SetIndexState(3, ttgletter.LetterX)

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
