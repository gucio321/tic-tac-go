package ttgplayer

import (
	"testing"

	"github.com/gucio321/tic-tac-go/pkg/core/ttgletter"
)

const playerString = "Player X"

func Test_Create(t *testing.T) {
	player := Create(PlayerPerson, ttgletter.LetterX, func(_ ttgletter.Letter) int { return 8 })

	if player.playerType != PlayerPerson {
		t.Fatal("Unexpected player created")
	}

	if player.letter != ttgletter.LetterX {
		t.Fatal("Unexpected player created")
	}

	if player.moveCb(ttgletter.LetterX) != 8 {
		t.Fatal("Unexpected player created")
	}

	if player.name != playerString {
		t.Fatal("Unexpected player created")
	}
}

func Test_Move(t *testing.T) {
	player := Create(PlayerPerson, ttgletter.LetterX, func(_ ttgletter.Letter) int { return 8 })

	if player.Move() != 8 {
		t.Fatal("unexpected move done")
	}
}

func Test_Letter(t *testing.T) {
	player := Create(PlayerPerson, ttgletter.LetterX, nil)

	if player.Letter() != ttgletter.LetterX {
		t.Fatal("unexpected letter returned")
	}
}

func Test_Name(t *testing.T) {
	player := Create(PlayerPerson, ttgletter.LetterX, nil)

	if player.Name() != playerString {
		t.Fatal("unexpected name returned")
	}
}
