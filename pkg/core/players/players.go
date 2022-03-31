// Package players contains an implementation of tic-tac-toe
// players system
package players

import (
	"crypto/rand"
	"fmt"
	"math/big"

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
	}

	result.RollFirstPlayer()

	return result
}

// PlayerX returns X player.
func (p *Players) PlayerX() player.Player {
	return p.playerX
}

// PlayerO returns O player.
func (p *Players) PlayerO() player.Player {
	return p.playerO
}

// Current returns current player's letter.
func (p *Players) Current() letter.Letter {
	return p.current
}

// CurrentPlayer returns current player.
func (p *Players) CurrentPlayer() player.Player {
	switch p.current {
	case letter.LetterX:
		return p.playerX
	case letter.LetterO:
		return p.playerO
	}

	return nil
}

// GetMove returns a current player's move.
func (p *Players) GetMove() int {
	return p.CurrentPlayer().GetMove()
}

// Next switch to the next player.
func (p *Players) Next() {
	p.current = p.current.Opposite()
}

// RollFirstPlayer sets random player as current.
func (p *Players) RollFirstPlayer() {
	dict := map[int64]letter.Letter{
		0: letter.LetterX,
		1: letter.LetterO,
	}

	randomNumber, err := rand.Int(rand.Reader, big.NewInt(int64(len(dict))))
	if err != nil {
		panic(fmt.Sprintf("Reading random number: %v", err))
	}

	p.current = dict[randomNumber.Int64()]
}
