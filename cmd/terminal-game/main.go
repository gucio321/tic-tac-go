package main

import (
	app "github.com/gucio321/tic-tac-go/internal/ttgapp"
)

func main() {
	instance := app.New(nil)
	instance.Run()
}
