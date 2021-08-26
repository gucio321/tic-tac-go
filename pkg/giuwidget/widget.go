// Package giuwidget contains a giu implementation of game
package giuwidget

import (
	"image/color"
	"math"
	"strconv"

	"github.com/AllenDang/giu"

	"github.com/gucio321/tic-tac-go/pkg/core/board/letter"
	"github.com/gucio321/tic-tac-go/pkg/core/ttgplayers/ttgplayer"
	"github.com/gucio321/tic-tac-go/pkg/game"
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

func (g *GameWidget) getGame() (state *game.Game) {
	if s := giu.Context.GetState(id); s == nil {
		state = game.Create(g.p1type, g.p2type).SetBoardSize(g.w, g.h, g.chainLen)
		giu.Context.SetState(id, state)
	} else {
		var ok bool
		state, ok = s.(*game.Game)
		if !ok {
			panic("Tic-Tac-Go: game.(*Game).getGame (internal): unexpected state recovered from giu")
		}
	}

	return state
}

// Build builds the game.
func (g *GameWidget) Build() {
	gameInstance := g.getGame()

	// nolint:staticcheck // will use it later
	isEnded, _ := gameInstance.Result()

	// nolint:staticcheck // TODO
	if isEnded {
		// build end layout
		// return
	}

	g.buildGameBoard(gameInstance)

	giu.Button("play new game").OnClick(func() {
		gameInstance.Dispose()
		go gameInstance.Run()
	}).Build()
}

func (g *GameWidget) buildGameBoard(gameInstance *game.Game) {
	board := giu.Layout{}

	for y := 0; y < gameInstance.Board().Height(); y++ {
		line := giu.Layout{}

		for x := 0; x < gameInstance.Board().Width(); x++ {
			idx := y*gameInstance.Board().Width() + x
			s := gameInstance.Board().GetIndexState(idx)
			btn := giu.Button(s.String()+"##BoardIndex"+strconv.Itoa(idx)).
				Size(buttonW, buttonH).OnClick(func() {
				if gameInstance.IsUserActionRequired() {
					if s == letter.LetterNone {
						gameInstance.TakeUserAction(idx)
					}
				}
			})

			var c color.RGBA

			switch s {
			case letter.LetterX:
				c = color.RGBA{
					R: 0,
					G: math.MaxUint8,
					B: 0,
					A: math.MaxUint8,
				}
			case letter.LetterO:
				c = color.RGBA{
					R: math.MaxUint8,
					G: 0,
					B: 0,
					A: math.MaxUint8,
				}
			}

			if gameEnd, l := gameInstance.Result(); gameEnd && l != letter.LetterNone {
				_, winningCombo := gameInstance.Board().IsWinner(gameInstance.Board().ChainLength(), l)
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
