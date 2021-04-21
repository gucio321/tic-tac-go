package ttgboard

import (
	"testing"

	"github.com/gucio321/tic-tac-go/ttggame/ttgletter"
)

func Test_NewBoard(t *testing.T) {
	correctBoard := Board{
		board:  make([]*ttgletter.Letter, 9),
		width:  3,
		height: 3,
	}
	for i := range correctBoard.board {
		correctBoard.board[i] = ttgletter.NewLetter()
	}

	board := NewBoard(3, 3, 3)

	if len(board.board) != len(correctBoard.board) {
		t.Fatal("Invalid board created")
	}

	for i := range board.board {
		if *board.board[i] != *correctBoard.board[i] {
			t.Fatal("Invalid board created")
		}
	}
}

func Test_setIndexState(t *testing.T) {
	board := NewBoard(3, 3, 3)

	board.SetIndexState(5, ttgletter.LetterX)

	if *board.board[5] != ttgletter.LetterX {
		t.Fatal("unexpected index was set by board.setIndexState")
	}
}

func Test_getIndexState(t *testing.T) {
	board := NewBoard(3, 3, 3)
	*board.board[5] = ttgletter.LetterX

	if l := board.GetIndexState(5); l != ttgletter.LetterX {
		t.Fatal("unexpected index was returned by board.getIndexState")
	}
}

func Test_isIndexFree(t *testing.T) {
	board := NewBoard(3, 3, 3)

	*board.board[5] = ttgletter.LetterX

	if board.IsIndexFree(5) {
		t.Fatal("isIndexFree returned unexpected value")
	}

	if !board.IsIndexFree(4) {
		t.Fatal("isIndexFree returned unexpected value")
	}
}

func Test_Copy(t *testing.T) {
	board := NewBoard(3, 3, 3)
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
	board := NewBoard(3, 3, 3)
	board.SetIndexState(4, ttgletter.LetterX)
	result := board.Cut(1, 1)

	if len(result.board) != 1 {
		t.Fatal("unexpected board cut")
	}

	if *result.board[0] != ttgletter.LetterX {
		t.Fatal("unexpected board cut")
	}
}
