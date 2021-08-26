package ttgplayer

import (
	"github.com/gucio321/tic-tac-go/pkg/core/board/letter"
)

// PlayerCb is a player move callback.
type PlayerCb func(letter.Letter) int

// PlayerType represents players' types.
type PlayerType int

// player types.
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

// Player represents the game player.
type Player struct {
	name       string
	playerType PlayerType
	letter     letter.Letter
	moveCb     PlayerCb
}

// Create creates a new player.
func Create(t PlayerType, playerLetter letter.Letter, cb PlayerCb) *Player {
	result := &Player{
		playerType: t,
		letter:     playerLetter,
		moveCb:     cb,
		name:       t.String() + " " + playerLetter.String(),
	}

	return result
}

// Move 'makes' player's move.
func (p *Player) Move() int {
	if p.moveCb == nil {
		panic("ttgplayer.(*Player).Move(): moveCb cannot be nil!")
	}

	return p.moveCb(p.Letter())
}

// Letter returns player's letter.
func (p *Player) Letter() letter.Letter {
	return p.letter
}

// Name returns player's name.
func (p *Player) Name() string {
	return p.name
}

// Type returns player's type.
func (p *Player) Type() PlayerType {
	return p.playerType
}
