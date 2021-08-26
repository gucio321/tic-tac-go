// Package players contains an implementation of tic-tac-toe
// players system
package players

import (
	"github.com/gucio321/tic-tac-go/pkg/core/board/letter"
	"github.com/gucio321/tic-tac-go/pkg/core/players/player"
)

// Players represents a pair of players.
type Players struct {
	player1,
	player2 *player.Player

	current letter.Letter
}

// Create creates a new players set.
func Create(player1Type player.PlayerType, cb1 player.PlayerCb, player2Type player.PlayerType, cb2 player.PlayerCb) *Players {
	result := &Players{
		player1: player.Create(player1Type, letter.LetterX, cb1),
		player2: player.Create(player2Type, letter.LetterO, cb2),
		current: letter.LetterX,
	}

	return result
}

// Player1 returns player1.
func (p *Players) Player1() *player.Player {
	return p.player1
}

// Player2 returns player2.
func (p *Players) Player2() *player.Player {
	return p.player2
}

// Current returns current player.
func (p *Players) Current() *player.Player {
	switch p.current {
	case p.player1.Letter():
		return p.player1
	case p.player2.Letter():
		return p.player2
	}

	return nil
}

// Move returns a current player's move.
func (p *Players) Move() int {
	return p.Current().Move()
}

// Next switch to the next player.
func (p *Players) Next() {
	p.current = p.current.Opposite()
}
