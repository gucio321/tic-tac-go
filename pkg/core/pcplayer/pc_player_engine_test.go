package pcplayer

import (
	"testing"

	"github.com/gucio321/tic-tac-go/pkg/core/board"
	"github.com/gucio321/tic-tac-go/pkg/core/board/letter"
)

func Test_canWin(t *testing.T) {
	w, h, c := 3, 3, 3
	b := board.Create(w, h, c)
	b.SetIndexState(0, letter.LetterX)
	b.SetIndexState(2, letter.LetterX)

	if ok, i := canWin(b, letter.LetterX); i == nil || i[0] != 1 || !ok {
		t.Fatalf("canWin returned wrong values\n%s", b)
	}

	b.SetIndexState(1, letter.LetterO)

	if ok, _ := canWin(b, letter.LetterX); ok {
		t.Fatalf("canWin returned true\n%s", b)
	}
}

func Test_canWinTwoMoves(t *testing.T) {
	w, h, c := 5, 5, 4
	b := board.Create(w, h, c)
	b.SetIndexState(12, letter.LetterX)
	b.SetIndexState(13, letter.LetterX)

	i := canWinTwoMoves(b, letter.LetterX)

	if len(i) != 1 {
		t.Fatalf("canWin returned wrong values\n%s", b)
	}

	if i[0] != 11 {
		t.Fatalf("canWin returned wrong values\n%s", b)
	}
}
