![CircleCI](https://img.shields.io/circleci/build/github/gucio321/tic-tac-go/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/gucio321/tic-tac-go)](https://goreportcard.com/report/github.com/gucio321/tic-tac-go)
[![GoDoc](https://pkg.go.dev/badge/github.com/gucio321/tic-tac-go?utm_source=godoc)](https://pkg.go.dev/mod/github.com/gucio321/tic-tac-go)
[![codecov](https://codecov.io/gh/gucio321/tic-tac-go/branch/master/graph/badge.svg)](https://codecov.io/gh/gucio321/tic-tac-go)

<image align="left" src="./logo.png">
<h1>Tic-Tac-Go</h1>
is an implementation
of the tic-tac-toe game written in <a href="https://go.dev">Golang</a>
In addition it implements a <i>relatively simple</i>
AI logic for PC players. For more details see
<a href="./pkg/core/pcplayer">here</a>
<br clear="all" />


# Installation

## Requirements

to run the game you just need to have
[GO programming language](https://golang.org) installed.

You may also want to use graphical version of the game,
so I suggest following
[GIU installation instruction](https://github.com/AllenDang/giu#install)

## Installing binaries

To install the game with golang api,
first download it: `go get github.com/gucio321/tic-tac-go`
and let's GO!
Since now, an executable will be present in `$GOPATH/bin/` directory.

### So how to run now?

After installation, just run
`go run github.com/gucio321/tic-tac-go/cmd/terminal-game` for
simple console game implementation or
`go run github.com/gucio321/tic-tac-go/cmd/giu-game` for
advanced graphical one.

## Well, but I'd like to know more about source code!

You can also download the source by running
`git clone https://github.com/gucio321/tic-tac-go`
and then, to set up the project:

```sh
cd tic-tac-go
go get -d ./...
```

## Screenshots

![tic tac go in terminal](docs/in_terminal.png)

![tic tac go with DearImgui using GIU](docs/in_giu.png)

## Motivation

When I'm learning a new programming language, I write a game
like this one to check myself. Because I liked [golang](https://golang.org)
and decided to share and improve my work.
