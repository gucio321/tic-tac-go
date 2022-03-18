package giuwidget

import (
	"github.com/AllenDang/giu"
	"github.com/gucio321/tic-tac-go/pkg/core/board/letter"
	"github.com/gucio321/tic-tac-go/pkg/game"
)

var _ giu.Disposable = &gameState{}

type gameState struct {
	game         *game.Game
	buttonClick  chan int
	gameEnded    bool
	winningCombo []int
}

// Dispose implements giu.Disposable
func (s *gameState) Dispose() {
	s.game.Dispose()
	s.gameEnded = false
	s.winningCombo = nil
}

func (g *GameWidget) getState() (state *gameState) {
	if s := giu.Context.GetState(id); s == nil {
		giu.Context.SetState(id, g.newState())
		return g.getState()
	} else {
		var ok bool
		state, ok = s.(*gameState)
		if !ok {
			panic("Tic-Tac-Go: game.(*Game).getGame (internal): unexpected state recovered from giu")
		}
	}

	return state
}

func (g *GameWidget) newState() *gameState {
	state := &gameState{
		game:        game.Create(g.p1type, g.p2type).SetBoardSize(g.w, g.h, g.chainLen),
		buttonClick: make(chan int),
	}

	state.game.Result(func(l letter.Letter) {
		_, state.winningCombo = state.game.Board().GetWinner()
		state.gameEnded = true
	})

	state.game.UserAction(func() int {
		return <-state.buttonClick
	})

	return state
}
