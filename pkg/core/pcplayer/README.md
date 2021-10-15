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

	"github.com/gucio321/tic-tac-go/pkg/core/board"
	"github.com/gucio321/tic-tac-go/pkg/core/board/letter"
	"github.com/gucio321/tic-tac-go/pkg/core/pcplayer"
)

func main() {
	// create board
	width, height := 3, 3
	chainLen := 3
	board := board.Create(width, height, chainLen)

	// fill board
	board.SetIndexState(0, letter.LetterX)
	board.SetIndexState(4, letter.LetterO)
	board.SetIndexState(8, letter.LetterX)
	board.SetIndexState(6, letter.LetterX)

	fmt.Println(board)

	// make move using AI engine
	i := pcplayer.GetPCMove(board, letter.LetterO)
	board.SetIndexState(i, letter.LetterO)

	fmt.Println(board)
}
```
