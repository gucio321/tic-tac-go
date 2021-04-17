package main

import (
	_ "embed"
)

//go:embed README.md
// nolint:gochecknoglobals // go embed requires global variable.
// THIS IS NOT  A LINT
var readme []byte

func main() {
	app.Run()
}
