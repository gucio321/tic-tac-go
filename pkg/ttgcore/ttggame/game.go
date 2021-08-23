// Package ttggame contains a common implementation of game.
//
// Usage:
// - create game / set everything
// - use Board() to display board
// - run Run()
package ttggame

import (
	"fmt"
	"sync"

	"github.com/gucio321/tic-tac-go/pkg/ttgcore/ttgboard"
	"github.com/gucio321/tic-tac-go/pkg/ttgcore/ttgletter"
	"github.com/gucio321/tic-tac-go/pkg/ttgcore/ttgpcplayer"
	"github.com/gucio321/tic-tac-go/pkg/ttgcore/ttgplayer"
)

const (
	defaultBoardW,
	defaultBoardH,
	defaultChainLen = 3, 3, 3
)

// Game represents a Tic-Tac-Go game.
type Game struct {
	board      *ttgboard.Board
	boardMutex *sync.Mutex

	players *ttgplayer.Players

	isRunning bool

	onContinue func()

	userAction         chan int
	userActionRequired bool

	gameOver bool
	winner   ttgletter.Letter
}

// Create creates a game instance.
func Create(p1type, p2type ttgplayer.PlayerType) *Game {
	result := &Game{
		board:              ttgboard.NewBoard(defaultBoardW, defaultBoardH, defaultChainLen),
		boardMutex:         &sync.Mutex{},
		userAction:         make(chan int),
		winner:             ttgletter.LetterNone,
		onContinue:         func() {},
		userActionRequired: false,
	}

	var p1Cb, p2Cb func(ttgletter.Letter) int

	switch p1type {
	case ttgplayer.PlayerPC:
		p1Cb = func(l ttgletter.Letter) int { return ttgpcplayer.GetPCMove(result.board, l) }
	case ttgplayer.PlayerPerson:
		p1Cb = func(_ ttgletter.Letter) int { return result.getUserAction() }
	}

	switch p2type {
	case ttgplayer.PlayerPC:
		p2Cb = func(l ttgletter.Letter) int { return ttgpcplayer.GetPCMove(result.board, l) }
	case ttgplayer.PlayerPerson:
		p2Cb = func(_ ttgletter.Letter) int { return result.getUserAction() }
	}

	result.players = ttgplayer.Create(p1type, p1Cb, p2type, p2Cb)

	return result
}

// setters

// SetBoardSize sets a size of board.
func (g *Game) SetBoardSize(w, h, c int) *Game {
	if g.isRunning {
		isRunningPanic("SetBoardSize")
	}

	g.board = ttgboard.NewBoard(w, h, c)

	return g
}

// OnContinue is called when player action is taken.
// it could be used to update board interface e.g. to redraw board
// in terminal.
func (g *Game) OnContinue(cb func()) *Game {
	if g.isRunning {
		isRunningPanic("OnContinue")
	}

	// don't set nil callback
	if cb == nil {
		return g
	}

	g.onContinue = cb

	return g
}

// runners

// Board returns game board.
func (g *Game) Board() *ttgboard.Board {
	g.boardMutex.Lock()
	b := g.board
	g.boardMutex.Unlock()

	return b
}

// IsUserActionRequired returns true if caller needs to ask user for his move.
func (g *Game) IsUserActionRequired() bool {
	return g.userActionRequired
}

// TakeUserAction should be used to pass player move.
func (g *Game) TakeUserAction(idx int) {
	g.userActionRequired = false
	g.userAction <- idx
}

// Result returns true if game is ended. in addition it returns its result.
// if LetterNone returned - it means that DRAW reached.
func (g *Game) Result() (bool, ttgletter.Letter) {
	return g.gameOver, g.winner
}

// CurrentPlayer returns a current player.
func (g *Game) CurrentPlayer() *ttgplayer.Player {
	return g.players.Current()
}

// Run runs the game.
// NOTE: should call in a new go routime.
func (g *Game) Run() {
	// prevent user from calling setter functions
	g.isRunning = true

	// main loop
	for {
		g.onContinue()
		g.Board().SetIndexState(g.players.Current().Move(), g.players.Current().Letter())

		if ok, _ := g.Board().IsWinner(g.Board().ChainLength(), g.players.Current().Letter()); ok {
			g.onContinue()
			g.winner = g.players.Current().Letter()
			g.isRunning = false
			g.gameOver = true

			return
		} else if g.Board().IsBoardFull() {
			g.onContinue()
			g.winner = ttgletter.LetterNone
			g.gameOver = true
			g.isRunning = false

			return
		}

		g.players.Next()
	}
}

// Dispose resets the game.
func (g *Game) Dispose() {
	if g.isRunning {
		panic("Tic-Tac-Go: ttggame.(*Game).Dispose: aborted")
	}

	*g = *Create(g.players.Player1().Type(), g.players.Player2().Type())
}

// internal

// isRunningPanic is called when a method is not allowed when `isRunning`.
func isRunningPanic(methodName string) {
	panic(fmt.Sprintf("Tic-Tac-Go: ttggame.(*Game).%s: invalid use of setter function after invoking Run", methodName))
}

// getUserAction is set as a player callback when PlayerTypePerson.
func (g *Game) getUserAction() int {
	g.userActionRequired = true

	return <-g.userAction
}
