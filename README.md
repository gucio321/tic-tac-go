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

## Screenshots

![menu](docs/menu.png)

![gameplay](docs/gameplay.png)

![help](docs/help.png)

## See also

there is a few wrappers of this game. see:

*  [tic-tac-go in Ebiten](https://github.com/gucio321/ttg-gui)
*  [tic-tac-go using DearImgui with GIU](https://github.com/gucio321/ttg-giu)

## Motivation

When I'm learning a new programming language, I'm writtin a game
like that to check myself. Because I liked the [golang](https://golang.org),
I decided to share and improve my work.
