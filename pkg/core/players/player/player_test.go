package player

import (
	"testing"

	"github.com/gucio321/tic-tac-go/pkg/core/board/letter"
	"github.com/stretchr/testify/assert"
)

const playerString = "Player X"

func Test_Create(t *testing.T) {
	const num = 8
	a := assert.New(t)
	player := Create(PlayerPerson, letter.LetterX, func(_ letter.Letter) int { return num })

	a.Equal(PlayerPerson, player.playerType, "Unexpected player created")
	a.Equal(letter.LetterX, player.letter, "Unexpected player created")
	a.Equal(num, player.moveCb(player.letter), "Unexpected player created")
	a.Equal(playerString, player.name, "Unexpected player created")
}

func Test_Move(t *testing.T) {
	const num = 8
	player := Create(PlayerPerson, letter.LetterX, func(_ letter.Letter) int { return num })

	assert.Equal(t, num, player.Move(), "unexpected move done")
}

func Test_Letter(t *testing.T) {
	player := Create(PlayerPerson, letter.LetterX, nil)

	assert.Equal(t, letter.LetterX, player.Letter(), "unexpected letter returned")
}

func Test_Name(t *testing.T) {
	player := Create(PlayerPerson, letter.LetterX, nil)

	assert.Equal(t, playerString, player.Name(), "unexpected name returned")
}

func Test_Type(t *testing.T) {
	player := Create(PlayerPerson, letter.LetterX, nil)

	assert.Equal(t, PlayerPerson, player.Type(), "Unexpected type of the player returned")
}
