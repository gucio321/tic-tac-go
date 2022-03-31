// Package game contains a common implementation of game.
//
// Usage:
// - create game / set all callbacks
// - use Board() to display board in implementation
// - execute Run()
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

	onContinue   func()
	resultCB     func(winnerLetter letter.Letter, winnerName string)
	userActionCB func() int
}

// Create creates a game instance.
func Create(playerXType, playerOType PlayerType) *Game {
	result := &Game{
		isRunning:  false,
		board:      board.Create(defaultBoardW, defaultBoardH, defaultChainLen),
		onContinue: func() {},
		resultCB:   func(letter.Letter, string) {},
		userActionCB: func() int {
			panic("Tic-Tac-Go: game.(*Game): user action callback is not set!")
		},
	}

	var playerX, playerO player.Player

	switch playerXType {
	case PlayerTypePC:
		playerX = pcplayer.NewPCPlayer(result.board, letter.LetterX)
	case PlayerTypeHuman:
		playerX = newHumanPlayer(result.getUserAction, letter.LetterX)
	}

	switch playerOType {
	case PlayerTypePC:
		playerO = pcplayer.NewPCPlayer(result.board, letter.LetterO)
	case PlayerTypeHuman:
		playerO = newHumanPlayer(result.getUserAction, letter.LetterO)
	}

	result.players = players.Create(playerX, playerO)

	return result
}

// setters

// SetBoardSize sets a size of board.
func (g *Game) SetBoardSize(w, h, c int) *Game {
	g.isRunningPanic("SetBoardSize")

	*g.board = *board.Create(w, h, c)

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

// UserAction sets user action callback called when game needs to get user's action.
func (g *Game) UserAction(cb func() int) {
	g.userActionCB = cb
}

// Result returns true if game is ended. in addition, it returns its result.
// if LetterNone returned - it means that DRAW reached.
func (g *Game) Result(resultCB func(winnerLetter letter.Letter, winnerName string)) *Game {
	g.resultCB = resultCB

	return g
}

// runners

// Board returns game board.
func (g *Game) Board() *board.Board {
	g.notRunningPanic("Board")
	b := g.board

	return b
}

// CurrentPlayer returns a current player.
func (g *Game) CurrentPlayer() player.Player {
	g.notRunningPanic("CurrentPlayer")

	return g.players.CurrentPlayer()
}

// Run runs the game.
// NOTE: should call in a new go routine.
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
			idx := g.players.CurrentPlayer().GetMove()

			// if loop was stopped by Dispose() or Stop(), exit the loop
			if !g.isRunning {
				return
			}

			g.Board().SetIndexState(idx, g.players.Current())

			if ok, _ := g.Board().IsWinner(g.players.Current()); ok {
				g.onContinue()
				g.isRunning = false
				g.resultCB(g.players.Current(), g.players.CurrentPlayer().String())

				return
			} else if g.Board().IsBoardFull() {
				g.onContinue()
				g.isRunning = false
				g.resultCB(letter.LetterNone, "")

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
	g.players.RollFirstPlayer()
}

// Stop safely stops the game loop invoked by (*Game).Run.
func (g *Game) Stop() {
	if !g.isRunning {
		return
	}

	g.isRunning = false
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

// notRunningPanic is called when a method is not allowed when `isRunning`.
func (g *Game) notRunningPanic(methodName string) {
	if !g.IsRunning() {
		panic(fmt.Sprintf("Tic-Tac-Go: game.(*Game).%s: invalid use of in-game function before calling Run", methodName))
	}
}

// getUserAction is set as a player callback when PlayerTypePerson.
func (g *Game) getUserAction() int {
	return g.userActionCB()
}
