// Package ttgplayers contains an implementation of tic-tac-toe
// players system
package ttgplayers

import (
	"github.com/gucio321/tic-tac-go/pkg/ttgcore/ttgletter"
	"github.com/gucio321/tic-tac-go/pkg/ttgcore/ttgplayers/ttgplayer"
)

// Players represents a pair of players.
type Players struct {
	player1,
	player2 *ttgplayer.Player
	current    ttgletter.Letter
	onContinue func()
}

// Create creates a new players set.
func Create(player1Type ttgplayer.PlayerType, cb1 ttgplayer.PlayerCb, player2Type ttgplayer.PlayerType, cb2 ttgplayer.PlayerCb) *Players {
	result := &Players{
		player1: ttgplayer.NewPlayer(player1Type, ttgletter.LetterX, cb1),
		player2: ttgplayer.NewPlayer(player2Type, ttgletter.LetterO, cb2),
		current: ttgletter.LetterX,
	}

	return result
}

// OnContinue is executed when Next called.
func (p *Players) OnContinue(cb func()) *Players {
	p.onContinue = cb

	return p
}

// Player1 returns player1.
func (p *Players) Player1() *ttgplayer.Player {
	return p.player1
}

// Player2 returns player2.
func (p *Players) Player2() *ttgplayer.Player {
	return p.player2
}

// Current returns current player.
func (p *Players) Current() *ttgplayer.Player {
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
	if p.onContinue != nil {
		p.onContinue()
	}

	p.current = p.current.Opposite()
}