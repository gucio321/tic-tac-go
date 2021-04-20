package ttgboard

import (
	"testing"

	"github.com/gucio321/tic-tac-go/ttggame/ttgletter"
)

func Test_NewBoard(t *testing.T) {
	correctBoard := make([]*ttgletter.Letter, 9)
	for i := range correctBoard {
		correctBoard[i] = ttgletter.NewLetter()
	}

	board := NewBoard(9)

	if len(*board) != len(correctBoard) {
		t.Fatal("Invalid board created")
	}

	for i := range *board {
		if *(*board)[i] != *correctBoard[i] {
			t.Fatal("Invalid board created")
		}
	}
}

func Test_setIndexState(t *testing.T) {
	board := NewBoard(9)

	board.SetIndexState(5, ttgletter.LetterX)

	if *(*board)[5] != ttgletter.LetterX {
		t.Fatal("unexpected index was set by board.setIndexState")
	}
}

func Test_getIndexState(t *testing.T) {
	board := NewBoard(9)
	*(*board)[5] = ttgletter.LetterX

	if l := board.GetIndexState(5); l != ttgletter.LetterX {
		t.Fatal("unexpected index was returned by board.getIndexState")
	}
}

func Test_isIndexFree(t *testing.T) {
	board := NewBoard(9)

	*(*board)[5] = ttgletter.LetterX

	if board.IsIndexFree(5) {
		t.Fatal("isIndexFree returned unexpected value")
	}

	if !board.IsIndexFree(4) {
		t.Fatal("isIndexFree returned unexpected value")
	}
}

func Test_Copy(t *testing.T) {
	board := NewBoard(9)
	board.SetIndexState(4, ttgletter.LetterX)
	newBoard := board.Copy()

	if len(*board) != len(*newBoard) {
		t.Fatal("Unexpected board copied")
	}

	for i := range *board {
		if *(*board)[i] != *(*newBoard)[i] {
			t.Fatal("Unexpected board copied")
		}
	}
}

func Test_Cut(t *testing.T) {
	board := NewBoard(9)
	board.SetIndexState(4, ttgletter.LetterX)
	result := board.Cut(3, 3, 1, 1)
	if len(*result) != 1 {
		t.Fatal("unexpected board cut")
	}

	if *(*result)[0] != ttgletter.LetterX {
		t.Fatal("unexpected board cut")
	}
}
