package main

import (
	_ "embed"

	game "github.com/gucio321/tic-tac-go/ttgmenu"
)

//go:embed README.md
// nolint:gochecknoglobals // go embed requires global variable.
// THIS IS NOT  A LINT
var readme []byte

func main() {
	app := game.New(readme)
	app.Run()
}
