![CircleCI](https://img.shields.io/circleci/build/github/gucio321/tic-tac-go/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/gucio321/tic-tac-go)](https://goreportcard.com/report/github.com/gucio321/tic-tac-go)
[![GoDoc](https://pkg.go.dev/badge/github.com/gucio321/tic-tac-go?utm_source=godoc)](https://pkg.go.dev/mod/github.com/gucio321/tic-tac-go)

## About

Tic-Tac-Go is a simple, command line implementation
of tic-tac-toe game written in [Golang](https://golang.org)

## Requirements

to run the game you only need to install [golang](https://golang.org)

### Installation

To install the game, first download it: `go get github.com/gucio321/tic-tac-go`
and let's GO!
Since now, an executale binary will be present in `$GOPATH/bin/`

### How to run?

After installation, just execute `go run github.com/gucio321/tic-tac-go`
or `$GOPATH/bin/tic-tac-go`

You can also download the source by `git clone https://github.com/gucio321/tic-tac-go`
and then:

```sh
cd tic-tac-go
go get -d ./...
go run .
```

## Examples usage of ttggame/ttgpcplayer

This simple example should show, how to use an AI game engine:

```golang
package main

import (
	"fmt"

	"github.com/gucio321/tic-tac-go/ttggame/ttgboard"
	"github.com/gucio321/tic-tac-go/ttggame/ttgpcplayer"
	"github.com/gucio321/tic-tac-go/ttggame/ttgletter"
)

func main() {
	// create board
	width, height := 3, 3
	chainLen := 3
	board := ttgboard.NewBoard(width, height, chainLen)

	// fill board
	board.SetIndexState(0, ttgletter.LetterX)
	board.SetIndexState(4, ttgletter.LetterO)
	board.SetIndexState(8, ttgletter.LetterX)
	board.SetIndexState(6, ttgletter.LetterX)

	fmt.Println(board)

	// make move using AI engine
	i := ttgpcplayer.GetPCMove(board, ttgletter.LetterO)
	board.SetIndexState(i, ttgletter.LetterO)

	fmt.Println(board)
}
```

## Screenshots

![menu](docs/menu.png)

![gameplay](docs/gameplay.png)

![help](docs/help.png)

## See also

there is a few wrappers of this game. see:

*  [tic-tac-go in Ebiten](https://github.com/gucio321/ttg-gui)
*  [tic-tac-go using DearImgui with GIU](https://github.com/gucio321/ttg-giu)
