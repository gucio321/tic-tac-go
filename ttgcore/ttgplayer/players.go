package ttgplayer

import "github.com/gucio321/tic-tac-go/ttgcore/ttgletter"

// Players represents a pair of players.
type Players struct {
	player1,
	player2 *Player
	current    ttgletter.Letter
	onContinue func()
}

// Create creates a new players set.
func Create(player1Type PlayerType, cb1 func(), player2Type PlayerType, cb2 func()) *Players {
	result := &Players{
		player1: NewPlayer(player1Type, ttgletter.LetterX, cb1),
		player2: NewPlayer(player2Type, ttgletter.LetterO, cb2),
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
func (p *Players) Player1() *Player {
	return p.player1
}

// Player2 returns player2.
func (p *Players) Player2() *Player {
	return p.player2
}

// Current returns current player.
func (p *Players) Current() *Player {
	switch p.current {
	case p.player1.Letter():
		return p.player1
	case p.player2.Letter():
		return p.player2
	}

	return nil
}

// Next switch to the next player.
func (p *Players) Next() {
	if p.onContinue != nil {
		p.onContinue()
	}

	p.current = p.current.Opposite()
}
