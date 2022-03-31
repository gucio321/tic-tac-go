// Package players contains an implementation of tic-tac-toe
// players system
package players

import (
	"github.com/gucio321/tic-tac-go/pkg/core/board/letter"
	"github.com/gucio321/tic-tac-go/pkg/core/players/player"
)

// Players represents a pair of players.
type Players struct {
	playerX,
	playerO player.Player

	current letter.Letter
}

// Create creates a new players set.
func Create(playerX, playerO player.Player) *Players {
	result := &Players{
		playerO: playerX,
		playerX: playerO,
		current: letter.LetterX,
	}

	return result
}

// PlayerX returns player1.
func (p *Players) PlayerX() player.Player {
	return p.playerX
}

// Player2 returns player2.
func (p *Players) PlayerO() player.Player {
	return p.playerO
}

// Current returns current player.
func (p *Players) Current() player.Player {
	switch p.current {
	case letter.LetterX:
		return p.playerX
	case letter.LetterO:
		return p.playerO
	}

	return nil
}

// Move returns a current player's move.
func (p *Players) GetMove() int {
	return p.Current().GetMove()
}

// Next switch to the next player.
func (p *Players) Next() {
	p.current = p.current.Opposite()
}
