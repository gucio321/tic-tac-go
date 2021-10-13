package player

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gucio321/tic-tac-go/pkg/core/board/letter"
)

func Test_PlayerType_String_invalid_type(t *testing.T) {
	assert.Panics(t, func() { _ = Type(5).String() }, "Calling string method of inocorrecty player's type didn't panicked")
}

func Test_Create(t *testing.T) {
	const num = 8

	a := assert.New(t)
	player := Create(PlayerPerson, letter.LetterX, func(_ letter.Letter) int { return num })

	a.Equal(PlayerPerson, player.playerType, "Unexpected player created")
	a.Equal(letter.LetterX, player.letter, "Unexpected player created")
	a.Equal(num, player.moveCb(player.letter), "Unexpected player created")
}

func Test_Move(t *testing.T) {
	tests := []struct {
		name   string
		number int
	}{
		{"Standard test", 8},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			player := Create(PlayerPerson, letter.LetterX, func(_ letter.Letter) int { return test.number })

			assert.Equal(tt, test.number, player.Move(), "unexpected move done")
		})
	}
}

func Test_Move_nil_callback(t *testing.T) {
	assert.Panics(t, func() {
		player := Create(PlayerPerson, letter.LetterX, nil)
		player.Move()
	}, "calling player.Move with nil callback didn't panicked")
}

func Test_Letter(t *testing.T) {
	player := Create(PlayerPerson, letter.LetterX, nil)

	assert.Equal(t, letter.LetterX, player.Letter(), "unexpected letter returned")
}

func Test_Name(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		pt       Type
		pl       letter.Letter
	}{
		{"person x", "Player X", PlayerPerson, letter.LetterX},
		{"pc o", "PC O", PlayerPC, letter.LetterO},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			player := Create(test.pt, test.pl, nil)

			assert.Equal(tt, test.expected, player.Name(), "unexpected name returned")
		})
	}
}

func Test_Type(t *testing.T) {
	player := Create(PlayerPerson, letter.LetterX, nil)

	assert.Equal(t, PlayerPerson, player.Type(), "Unexpected type of the player returned")
}
