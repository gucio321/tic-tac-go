package ttgboard

import (
	"testing"
)

func Test_NewBoard(t *testing.T) {
	correctBoard := make([]*Letter, 9)
	for i := range correctBoard {
		correctBoard[i] = newBoardIndex()
	}

	board := newBoard(3, 3)

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
	board := newBoard(3, 3)
	board.setIndexState(5, LetterX)
	if *(*board)[5] != LetterX {
		t.Fatal("unexpected index was set by board.setIndexState")
	}
}

func Test_getIndexState(t *testing.T) {
	board := newBoard(3, 3)
	*(*board)[5] = LetterX
	l := board.getIndexState(5)
	if l != LetterX {
		t.Fatal("unexpected index was returned by board.getIndexState")
	}
}

func Test_isIndexFree(t *testing.T) {
	board := newBoard(3, 3)
	*(*board)[5] = LetterX
	if board.isIndexFree(5) {
		t.Fatal("isIndexFree returned unexpected value")
	}

	if !board.isIndexFree(4) {
		t.Fatal("isIndexFree returned unexpected value")
	}
}
