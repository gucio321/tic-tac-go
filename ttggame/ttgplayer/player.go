package ttgplayer

import (
	"github.com/gucio321/tic-tac-go/ttggame/ttgletter"
)

// PlayerType represents players' types
type PlayerType int

// player types
const (
	PlayerPC PlayerType = iota
	PlayerPerson
)

func (p PlayerType) String() string {
	switch p {
	case PlayerPC:
		return "PC"
	case PlayerPerson:
		return "Player"
	}

	return "?"
}

// Player represents the game player
type Player struct {
	name       string
	playerType PlayerType
	letter     ttgletter.Letter
	moveCb     func() (i int)
}

// NewPlayer creates a new player
func NewPlayer(t PlayerType, letter ttgletter.Letter, cb func() (i int)) *Player {
	result := &Player{
		playerType: t,
		letter:     letter,
		moveCb:     cb,
		name:       t.String() + " " + letter.String(),
	}

	return result
}

// Move 'makes' player's move
func (p *Player) Move() int {
	return p.moveCb()
}

// Letter returns player's letter
func (p *Player) Letter() ttgletter.Letter {
	return p.letter
}

// Name returns player's name
func (p *Player) Name() string {
	return p.name
}
