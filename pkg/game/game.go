// Package game contains a common implementation of game.
//
// Usage:
// - create game / set everything
// - use Board() to display board in implementation
// - run Run()
// -- take user action when IsUserActionRequired()
// - call Dispose() to reset
package game

import (
	"fmt"

	"github.com/gucio321/tic-tac-go/pkg/core/board"
	"github.com/gucio321/tic-tac-go/pkg/core/board/letter"
	"github.com/gucio321/tic-tac-go/pkg/core/pcplayer"
	"github.com/gucio321/tic-tac-go/pkg/core/players"
	"github.com/gucio321/tic-tac-go/pkg/core/players/player"
)

const (
	defaultBoardW,
	defaultBoardH,
	defaultChainLen = 3, 3, 3
)

// Game represents a Tic-Tac-Go game.
type Game struct {
	board *board.Board

	players *players.Players

	isRunning bool

	onContinue func()

	userAction         chan int
	userActionRequired bool

	resultCB func(letter.Letter)
}

// Create creates a game instance.
func Create(p1type, p2type player.Type) *Game {
	result := &Game{
		board:              board.Create(defaultBoardW, defaultBoardH, defaultChainLen),
		userAction:         make(chan int),
		onContinue:         func() {},
		userActionRequired: false,
		resultCB:           func(letter.Letter) {},
	}

	var p1Cb, p2Cb func(letter.Letter) int

	switch p1type {
	case player.PlayerPC:
		p1Cb = func(l letter.Letter) int { return pcplayer.GetPCMove(result.board, l) }
	case player.PlayerPerson:
		p1Cb = func(_ letter.Letter) int { return result.getUserAction() }
	}

	switch p2type {
	case player.PlayerPC:
		p2Cb = func(l letter.Letter) int { return pcplayer.GetPCMove(result.board, l) }
	case player.PlayerPerson:
		p2Cb = func(_ letter.Letter) int { return result.getUserAction() }
	}

	result.players = players.Create(p1type, p1Cb, p2type, p2Cb)

	return result
}

// setters

// SetBoardSize sets a size of board.
func (g *Game) SetBoardSize(w, h, c int) *Game {
	g.isRunningPanic("SetBoardSize")

	g.board = board.Create(w, h, c)

	return g
}

// OnContinue is called when player action is taken.
// it could be used to update board interface e.g. to redraw board
// in terminal.
func (g *Game) OnContinue(cb func()) *Game {
	g.isRunningPanic("OnContinue")

	// don't set nil callback
	if cb != nil {
		g.onContinue = cb
	}

	return g
}

// runners

// Board returns game board.
func (g *Game) Board() *board.Board {
	b := g.board

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
func (g *Game) Result(resultCB func(letter.Letter)) *Game {
	g.resultCB = resultCB
	return g
}

// CurrentPlayer returns a current player.
func (g *Game) CurrentPlayer() *player.Player {
	return g.players.Current()
}

// Run runs the game.
// NOTE: should call in a new go routime.
func (g *Game) Run() {
	if g.isRunning {
		panic("Tic-Tac-Go: game.(*Game).Run: invalid call of Run when game is running.")
	}

	// prevent user from calling setter functions
	g.isRunning = true

	go func() {
		// main loop
		for {
			g.onContinue()
			idx := g.players.Current().Move()

			// if loop was stopped by Dispose() or Stop(), exit the loop
			if !g.isRunning {
				return
			}

			g.Board().SetIndexState(idx, g.players.Current().Letter())

			if ok, _ := g.Board().IsWinner(g.players.Current().Letter()); ok {
				g.onContinue()
				g.isRunning = false
				g.resultCB(g.players.Current().Letter())

				return
			} else if g.Board().IsBoardFull() {
				g.onContinue()
				g.isRunning = false
				g.resultCB(g.players.Current().Letter())

				return
			}

			g.players.Next()
		}
	}()
}

// Dispose resets the game.
func (g *Game) Dispose() {
	g.Stop()
	g.Reset()
}

// Reset resets the game.
func (g *Game) Reset() {
	if g.isRunning {
		panic("Tic-Tac-Go: game.(*Game).Reset() call when game is running. Did you forgot to invoke (*Game).Stop()?")
	}

	*g.board = *board.Create(g.board.Width(), g.board.Height(), g.board.ChainLength())
}

// Stop savely stops the game loop invoked by (*Game).Run.
func (g *Game) Stop() {
	if !g.isRunning {
		return
	}

	g.isRunning = false

	// if awaiting user action, pass any invalid number to exit `Run` loop
	if g.IsUserActionRequired() {
		g.TakeUserAction(-1)
	}
}

// IsRunning returns true if Run loop was invoked.
func (g *Game) IsRunning() bool {
	return g.isRunning
}

// internal

// isRunningPanic is called when a method is not allowed when `isRunning`.
func (g *Game) isRunningPanic(methodName string) {
	if g.IsRunning() {
		panic(fmt.Sprintf("Tic-Tac-Go: game.(*Game).%s: invalid use of setter function after invoking Run", methodName))
	}
}

// getUserAction is set as a player callback when PlayerTypePerson.
func (g *Game) getUserAction() int {
	g.userActionRequired = true

	return <-g.userAction
}
