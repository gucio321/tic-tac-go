package game

import "github.com/gucio321/tic-tac-go/pkg/core/players/player"

var _ player.Player = &humanPlayer{}

type humanPlayer struct{}

func newHumanPlayer() *humanPlayer {
	return &humanPlayer{}
}

func (h *humanPlayer) GetMove() int {
	panic("not implemented")
}

func (h *humanPlayer) String() string {
	panic("not implemented")
}
