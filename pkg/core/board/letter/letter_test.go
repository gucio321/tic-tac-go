package ttgletter

import (
	"testing"
)

func Test_newBoardIndex(t *testing.T) {
	letter := Create()

	if *letter != LetterNone {
		t.Fatal("Unexpected letter index created")
	}
}

func Test_SetState(t *testing.T) {
	letter := Create()

	letter.SetState(LetterX)

	if *letter != LetterX {
		t.Fatal("Unexpected state was set")
	}
}

func Test_String(t *testing.T) {
	letter := Create()

	*letter = LetterX

	if letter.String() != "X" {
		t.Fatal("unexpected string returned")
	}
}

func Test_IsNone(t *testing.T) {
	l := Create()
	if !l.IsNone() {
		t.Fatal("letter isn't none, but should be")
	}

	*l = LetterX

	if l.IsNone() {
		t.Fatal("leter is none, but shouldn't")
	}
}

func Test_Opposite(t *testing.T) {
	l := Create()

	if l.Opposite() != LetterNone {
		t.Fatal("opposite to letter none should be letter none, but isn't")
	}

	*l = LetterX

	if l.Opposite() != LetterO {
		t.Fatal("Letter.Opposite returned unexpected value")
	}
}