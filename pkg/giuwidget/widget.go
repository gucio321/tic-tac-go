// Package giuwidget contains a giu implementation of game
package giuwidget

import (
	"image/color"
	"math"
	"strconv"

	"github.com/AllenDang/giu"
	"golang.org/x/image/colornames"

	"github.com/gucio321/tic-tac-go/pkg/core/board/letter"
	"github.com/gucio321/tic-tac-go/pkg/core/players/player"
)

const id = "Tic-Tac-Go-game"

const (
	buttonsSpacing = 1
)

// GameWidget represents a giu implementation of tic-tac-go.
type GameWidget struct {
	w, h, chainLen int
	p1type, p2type player.Type
}

// Game creates GameWidget.
func Game(p1type, p2type player.Type, w, h, c int) *GameWidget {
	return &GameWidget{
		w:        w,
		h:        h,
		chainLen: c,
		p1type:   p1type,
		p2type:   p2type,
	}
}

func (g *GameWidget) RunGame() {
	state := g.getState()

	state.displayBoard = true
	state.game.Run()
}

// Build builds the game.
func (g *GameWidget) Build() {
	state := g.getState()

	if state.displayBoard {
		g.buildGameBoard(state)
	}

	giu.Button("play new game").OnClick(func() {
		g.RunGame()
	}).Disabled(state.game.IsRunning()).Build()
}

func (g *GameWidget) buildGameBoard(state *gameState) {
	if state.game.IsRunning() {
		state.currentBoard = state.game.Board()
	}

	avilW, avilH := giu.GetAvailableRegion()
	boardW, boardH := state.currentBoard.Width(), state.currentBoard.Height()
	boardContainerSize := float32(math.Min(float64(avilW), float64(avilH)))
	buttonW, buttonH := (boardContainerSize-float32((boardW-1)*buttonsSpacing))/float32(boardW),
		(boardContainerSize-float32((boardH-1)*buttonsSpacing))/float32(boardH)

	board := giu.Layout{}

	for y := 0; y < state.currentBoard.Height(); y++ {
		line := giu.Layout{}

		for x := 0; x < state.currentBoard.Width(); x++ {
			idx := y*state.currentBoard.Width() + x
			s := state.currentBoard.GetIndexState(idx)
			btn := giu.Button(s.String()+"##BoardIndex"+strconv.Itoa(idx)).
				Size(buttonW, buttonH).OnClick(func() {
				if state.game.IsRunning() {
					if s == letter.LetterNone {
						state.buttonClick <- idx
					}

					return
				}

				state.Dispose()
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

			if state.gameEnded {
				for _, i := range state.winningCombo {
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

	giu.Layout{
		giu.Style().
			SetStyle(giu.StyleVarItemSpacing, buttonsSpacing, buttonsSpacing).
			SetStyle(giu.StyleVarFrameRounding, 0, 0).
			SetStyle(giu.StyleVarFrameBorderSize, 0, 0).
			SetStyle(giu.StyleVarChildBorderSize, 0, 0).
			SetColor(giu.StyleColorButton, colornames.Black).
			SetColor(giu.StyleColorButtonHovered, color.RGBA{20, 20, 20, 255}).
			SetColor(giu.StyleColorButtonActive, colornames.Black).
			SetColor(giu.StyleColorChildBg, colornames.White).
			SetFontSize(80).To(
			giu.Child().Size(boardContainerSize, boardContainerSize).Layout(
				board,
			),
		),
	}.Build()
}
