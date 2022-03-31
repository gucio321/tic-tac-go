package game

import (
	"github.com/gucio321/tic-tac-go/pkg/core/board/letter"
	"github.com/gucio321/tic-tac-go/pkg/core/players/player"
)

var _ player.Player = &humanPlayer{}

type humanPlayer struct {
	callback func() int
	letter   letter.Letter
}

func newHumanPlayer(cb func() int, l letter.Letter) *humanPlayer {
	return &humanPlayer{
		callback: cb,
		letter:   l,
	}
}

func (h *humanPlayer) GetMove() int {
	return h.callback()
}

func (h *humanPlayer) String() string {
	return "Player " + h.letter.String()
}
