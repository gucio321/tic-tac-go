package main

import (
	"fmt"

	"github.com/gucio321/tic-tac-go/pkg/core/board"
	"github.com/gucio321/tic-tac-go/pkg/core/board/letter"
	"github.com/gucio321/tic-tac-go/pkg/core/pcplayer"
)

func main() {
	b := board.Create(3, 3, 3)
	b.SetIndexState(0, letter.LetterX)
	// b.SetIndexState(1, letter.LetterO)
	pcPlayer := pcplayer.NewPCPlayer(b, letter.LetterO)
	fmt.Println(pcPlayer.MinMax(b, letter.LetterO, 0, true))
}
