## Description

a simple pseudo-AI engine for calculating a move on Tic-Tac-Toe based board.
Like whole repository, this engine was written in [golang](https://golang.org)


## Why not MaxMin algorithm?

There is a lot of Tic-Tac-Toe implementations, which uses
an implementation of [MinMax algorithm](https://en.wikipedia.org/wiki/Minimax).
This algorithm might be very effective,
but it has one major disadventige: 
MinMax is very demanding (in terms of hardware).
It loops itself and gets very lot of RAM.
This fact, makes the algorithm useless on larger than 3x3 boards

## Example

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
