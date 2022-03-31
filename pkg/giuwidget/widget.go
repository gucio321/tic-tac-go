// Package giuwidget contains a giu implementation of game
package giuwidget

import (
	"image/color"
	"math"
	"strconv"

	"github.com/AllenDang/giu"
	"golang.org/x/image/colornames"

	"github.com/gucio321/tic-tac-go/pkg/core/board/letter"
	"github.com/gucio321/tic-tac-go/pkg/game"
)

const id = "Tic-Tac-Go-game"

const (
	buttonsSpacing   = 3
	defaultBoardSize = 3
	inputIntW        = 40
	menuFontSize     = 30
	headerFontSize   = 80
)

// GameWidget represents a giu implementation of tic-tac-go.
type GameWidget struct {
	playerXType, playerOType game.PlayerType
}

// Game creates GameWidget.
func Game(playerXType, playerOType game.PlayerType) *GameWidget {
	return &GameWidget{
		playerXType: playerXType,
		playerOType: playerOType,
	}
}

func (g *GameWidget) runGame() {
	state := g.getState()

	state.game.SetBoardSize(int(state.w), int(state.h), int(state.chainLen))
	state.displayBoard = true
	state.game.Run()
}

// Build builds the game.
func (g *GameWidget) Build() {
	// nolint:ifshort // https://github.com/golangci/golangci-lint/issues/2662
	state := g.getState()

	if state.displayBoard {
		g.buildGameBoard(state)

		return
	}

	giu.Layout{
		giu.Align(giu.AlignCenter).To(
			giu.Style().SetFontSize(headerFontSize).To(
				giu.Row(
					giu.Style().SetColor(giu.StyleColorText, colornames.Red).To(
						giu.Label("TIC"),
					),
					giu.Label("-"),
					giu.Style().SetColor(giu.StyleColorText, colornames.Blue).To(
						giu.Label("TAC"),
					),
					giu.Label("-"),
					giu.Style().SetColor(giu.StyleColorText, colornames.Green).To(
						giu.Label("GO"),
					),
				),
			),
			giu.Row(
				giu.Style().SetFontSize(menuFontSize).To(
					giu.Label("Width: "),
				),
				giu.Style().SetFontSize(menuFontSize).To(
					giu.InputInt(&state.w).Size(inputIntW),
				),
			),
			giu.Row(
				giu.Style().SetFontSize(menuFontSize).To(
					giu.Label("Heigh: "),
				),
				giu.Style().SetFontSize(menuFontSize).To(
					giu.InputInt(&state.h).Size(inputIntW),
				),
			),
			giu.Row(
				giu.Style().SetFontSize(menuFontSize).To(
					giu.Label("Chain Length: "),
				),
				giu.Style().SetFontSize(menuFontSize).To(
					giu.InputInt(&state.chainLen).Size(inputIntW),
				),
			),
			giu.Style().SetFontSize(menuFontSize).To(
				giu.Button("START GAME").OnClick(func() {
					g.runGame()
				}),
			),
		),
	}.Build()
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

		board = append(board,
			giu.Row(line...),
		)
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
			SetFontSize(buttonH).To(
			giu.Align(giu.AlignCenter).To(
				giu.Child().Size(boardContainerSize, boardContainerSize).Layout(
					board,
				),
			),
		),
	}.Build()
}
