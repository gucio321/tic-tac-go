package ttgplayer

import (
	"testing"

	"github.com/gucio321/tic-tac-go/ttggame/ttgletter"
)

const playerString = "Player X"

func Test_NewPlayer(t *testing.T) {
	player := NewPlayer(PlayerPerson, ttgletter.LetterX, func() int { return 8 })

	if player.playerType != PlayerPerson {
		t.Fatal("Unexpected player created")
	}

	if player.letter != ttgletter.LetterX {
		t.Fatal("Unexpected player created")
	}

	if player.moveCb() != 8 {
		t.Fatal("Unexpected player created")
	}

	if player.name != playerString {
		t.Fatal("Unexpected player created")
	}
}

func Test_Move(t *testing.T) {
	player := NewPlayer(PlayerPerson, ttgletter.LetterX, func() int { return 8 })

	if player.Move() != 8 {
		t.Fatal("unexpected move done")
	}
}

func Test_Letter(t *testing.T) {
	player := NewPlayer(PlayerPerson, ttgletter.LetterX, func() int { return 8 })

	if player.Letter() != ttgletter.LetterX {
		t.Fatal("unexpected letter returned")
	}
}

func Test_Name(t *testing.T) {
	player := NewPlayer(PlayerPerson, ttgletter.LetterX, func() int { return 8 })

	if player.Name() != playerString {
		t.Fatal("unexpected name returned")
	}
}
