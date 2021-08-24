package main

import (
	"github.com/AllenDang/giu"

	"github.com/gucio321/tic-tac-go/pkg/giuwidget"
	"github.com/gucio321/tic-tac-go/pkg/ttgcore/ttgplayer"
)

func main() {
	wnd := giu.NewMasterWindow("Tic-Tac-Go", 640, 480, 0)
	wnd.Run(func() {
		giu.SingleWindow().Layout(giuwidget.Game(ttgplayer.PlayerPerson, ttgplayer.PlayerPC, 3, 3, 3))
	})
}
