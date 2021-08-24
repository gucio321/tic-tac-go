// Package giuwidget contains a giu implementation of game
package giuwidget

import (
	"image/color"
	"math"
	"strconv"

	"github.com/AllenDang/giu"

	"github.com/gucio321/tic-tac-go/pkg/ttgcore/ttggame"
	"github.com/gucio321/tic-tac-go/pkg/ttgcore/ttgletter"
	"github.com/gucio321/tic-tac-go/pkg/ttgcore/ttgplayer"
)

const id = "Tic-Tac-Go-game"

const (
	buttonW, buttonH = 100, 100
)

// GameWidget represents a giu implementation of tic-tac-go.
type GameWidget struct {
	w, h, chainLen int
	p1type, p2type ttgplayer.PlayerType
}

// Game creates GameWidget.
func Game(p1type, p2type ttgplayer.PlayerType, w, h, c int) *GameWidget {
	return &GameWidget{
		w:        w,
		h:        h,
		chainLen: c,
		p1type:   p1type,
		p2type:   p2type,
	}
}

func (g *GameWidget) getGame() (state *ttggame.Game) {
	if s := giu.Context.GetState(id); s == nil {
		state = ttggame.Create(g.p1type, g.p2type).SetBoardSize(g.w, g.h, g.chainLen)
		giu.Context.SetState(id, state)
	} else {
		var ok bool
		state, ok = s.(*ttggame.Game)
		if !ok {
			panic("Tic-Tac-Go: ttggame.(*Game).getGame (internal): unexpected state recovered from giu")
		}
	}

	return state
}

// Build builds the game.
func (g *GameWidget) Build() {
	game := g.getGame()

	// nolint:ifshort,staticcheck // will use it later
	isEnded, _ := game.Result()

	// nolint:staticcheck // TODO
	if isEnded {
		// build end layout
		// return
	}

	g.buildGameBoard(game)

	giu.Button("play new game").OnClick(func() {
		game.Dispose()
		go game.Run()
	}).Build()
}

func (g *GameWidget) buildGameBoard(game *ttggame.Game) {
	board := giu.Layout{}

	for y := 0; y < game.Board().Height(); y++ {
		line := giu.Layout{}

		for x := 0; x < game.Board().Width(); x++ {
			idx := y*game.Board().Width() + x
			s := game.Board().GetIndexState(idx)
			btn := giu.Button(s.String()+"##BoardIndex"+strconv.Itoa(idx)).
				Size(buttonW, buttonH).OnClick(func() {
				if game.IsUserActionRequired() {
					if s == ttgletter.LetterNone {
						game.TakeUserAction(idx)
					}
				}
			})

			var c color.RGBA

			switch s {
			case ttgletter.LetterX:
				c = color.RGBA{
					R: 0,
					G: math.MaxUint8,
					B: 0,
					A: math.MaxUint8,
				}
			case ttgletter.LetterO:
				c = color.RGBA{
					R: math.MaxUint8,
					G: 0,
					B: 0,
					A: math.MaxUint8,
				}
			}

			if gameEnd, l := game.Result(); gameEnd && l != ttgletter.LetterNone {
				_, winningCombo := game.Board().IsWinner(game.Board().ChainLength(), l)
				for _, i := range winningCombo {
					if i == idx {
						c = color.RGBA{
							R: 0,
							G: 0,
							B: math.MaxUint8,
							A: math.MaxUint8,
						}
					}
				}
			}

			styled := giu.Style().SetColor(giu.StyleColorText, c).To(btn)

			line = append(line, styled)
		}

		board = append(board, giu.Row(line...))
	}

	board.Build()
}
