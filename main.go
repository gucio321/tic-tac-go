package main

import (
	_ "embed"

	app "github.com/gucio321/tic-tac-go/ttgapp"
)

//go:embed README.md
// nolint:gochecknoglobals // go embed requires global variable.
// THIS IS NOT  A LINT
var readme []byte

func main() {
	instance := app.New(readme)
	instance.Run()
}
