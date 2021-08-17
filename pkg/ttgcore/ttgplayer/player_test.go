package ttgplayer

import (
	"testing"

	"github.com/gucio321/tic-tac-go/pkg/ttgcore/ttgletter"
)

const playerString = "Player X"

func Test_NewPlayer(t *testing.T) {
	i := 0

	moveTest := &i

	player := NewPlayer(PlayerPerson, ttgletter.LetterX, func() { *moveTest = 8 })

	if player.playerType != PlayerPerson {
		t.Fatal("Unexpected player created")
	}

	if player.letter != ttgletter.LetterX {
		t.Fatal("Unexpected player created")
	}

	player.moveCb()

	if *moveTest != 8 {
		t.Fatal("Unexpected player created")
	}

	if player.name != playerString {
		t.Fatal("Unexpected player created")
	}
}

func Test_Move(t *testing.T) {
	i := 0

	moveTest := &i

	player := NewPlayer(PlayerPerson, ttgletter.LetterX, func() { *moveTest = 8 })

	player.Move()

	if *moveTest != 8 {
		t.Fatal("unexpected move done")
	}
}

func Test_Letter(t *testing.T) {
	player := NewPlayer(PlayerPerson, ttgletter.LetterX, nil)

	if player.Letter() != ttgletter.LetterX {
		t.Fatal("unexpected letter returned")
	}
}

func Test_Name(t *testing.T) {
	player := NewPlayer(PlayerPerson, ttgletter.LetterX, nil)

	if player.Name() != playerString {
		t.Fatal("unexpected name returned")
	}
}
