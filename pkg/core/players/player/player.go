package player

import (
	"fmt"

	"github.com/gucio321/tic-tac-go/pkg/core/board/letter"
)

// Callback is a player move callback.
type Callback func(letter.Letter) int

// Type represents players' types.
type Type int

// player types.
const (
	PlayerPC Type = iota
	PlayerPerson
)

func (p Type) String() string {
	lookup := map[Type]string{
		PlayerPC:     "PC",
		PlayerPerson: "Player",
	}

	result, ok := lookup[p]

	if !ok {
		panic(fmt.Sprintf("Tic-Tac-Go: player.Type.String(): unknown player type %d", p))
	}

	return result
}

// Player represents the game player.
type Player struct {
	playerType Type
	letter     letter.Letter
	moveCb     Callback
}

// Create creates a new player.
func Create(t Type, playerLetter letter.Letter, cb Callback) *Player {
	result := &Player{
		playerType: t,
		letter:     playerLetter,
		moveCb:     cb,
	}

	return result
}

// Move 'makes' player's move.
func (p *Player) Move() int {
	if p.moveCb == nil {
		panic("player.(*Player).Move(): moveCb cannot be nil!")
	}

	return p.moveCb(p.Letter())
}

// Letter returns player's letter.
func (p *Player) Letter() letter.Letter {
	return p.letter
}

// Name returns player's name.
func (p *Player) Name() string {
	return fmt.Sprintf("%s %s", p.playerType, p.letter)
}

// Type returns player's type.
func (p *Player) Type() Type {
	return p.playerType
}
