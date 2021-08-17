package ttgpcplayer

import (
	"testing"

	"github.com/gucio321/tic-tac-go/pkg/ttgcore/ttgboard"
	"github.com/gucio321/tic-tac-go/pkg/ttgcore/ttgletter"
)

func Test_canWin(t *testing.T) {
	w, h, c := 3, 3, 3
	board := ttgboard.NewBoard(w, h, c)
	board.SetIndexState(0, ttgletter.LetterX)
	board.SetIndexState(2, ttgletter.LetterX)

	if i, ok := canWin(board, ttgletter.LetterX); i != 1 || !ok {
		t.Fatalf("canWin returned wrong values\n%s", board)
	}

	board.SetIndexState(1, ttgletter.LetterO)

	if _, ok := canWin(board, ttgletter.LetterX); ok {
		t.Fatalf("canWin returned true\n%s", board)
	}
}

func Test_canWinTwoMoves(t *testing.T) {
	w, h, c := 5, 5, 4
	board := ttgboard.NewBoard(w, h, c)
	board.SetIndexState(12, ttgletter.LetterX)
	board.SetIndexState(13, ttgletter.LetterX)

	i := canWinTwoMoves(board, ttgletter.LetterX)

	if len(i) != 1 {
		t.Fatalf("canWin returned wrong values\n%s", board)
	}

	if i[0] != 11 {
		t.Fatalf("canWin returned wrong values\n%s", board)
	}
}
