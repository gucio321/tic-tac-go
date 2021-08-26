package main

import (
	"github.com/AllenDang/giu"

	"github.com/gucio321/tic-tac-go/pkg/core/players/ttgplayer"
	"github.com/gucio321/tic-tac-go/pkg/giuwidget"
)

func main() {
	const (
		boardSize        = 3
		screenX, screenY = 640, 480
	)

	wnd := giu.NewMasterWindow("Tic-Tac-Go", screenX, screenY, 0)
	wnd.Run(func() {
		giu.SingleWindow().Layout(giuwidget.Game(ttgplayer.PlayerPerson, ttgplayer.PlayerPC, boardSize, boardSize, boardSize))
	})
}
