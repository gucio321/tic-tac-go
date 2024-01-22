package giuwidget

import (
	"github.com/AllenDang/giu"

	"github.com/gucio321/tic-tac-go/pkg/core/board"
	"github.com/gucio321/tic-tac-go/pkg/game"
)

var _ giu.Disposable = &gameState{}

type gameState struct {
	w, h, chainLen           int32
	game                     *game.Game
	buttonClick              chan int
	gameEnded                bool
	winningCombo             []int
	currentBoard             *board.Board
	displayBoard             bool
	playerXType, playerOType int32
}

// Dispose implements giu.Disposable.
func (s *gameState) Dispose() {
	s.game.Dispose()
	s.gameEnded = false
	s.winningCombo = nil
	s.displayBoard = false
}

func (g *GameWidget) getState() (state *gameState) {
	s := giu.Context.GetState(id)
	if s == nil {
		giu.Context.SetState(id, g.newState())

		return g.getState()
	}

	var ok bool

	state, ok = s.(*gameState)
	if !ok {
		panic("Tic-Tac-Go: game.(*Game).getGame (internal): unexpected state recovered from giu")
	}

	return state
}

func (g *GameWidget) newState() *gameState {
	state := &gameState{
		w:           defaultBoardSize,
		h:           defaultBoardSize,
		chainLen:    defaultBoardSize,
		gameEnded:   false,
		buttonClick: make(chan int),
	}

	return state
}
