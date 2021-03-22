package ttgapp

import (
	"github.com/AllenDang/giu"

	"github.com/gucio321/tic-tac-go/ttgmenu"
)

const (
	windowTitle  = "Tic-Tac-Go"
	windowWidth  = 500
	windowHeight = 500
)

type App struct {
	menu *ttgmenu.Menu
}

func Create() *App {
	result := &App{
		menu: ttgmenu.NewMenu(),
	}

	return result
}

func (a *App) Run() {
	wnd := giu.NewMasterWindow(windowTitle, windowWidth, windowHeight, 0, nil)
	wnd.Run(a.render)
}

func (a *App) render() {
	giu.SingleWindow("Game").Layout(
		a.menu.Build(),
	)
}
