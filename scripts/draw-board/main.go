package main

import (
	"flag"
	"fmt"

	"github.com/gucio321/tic-tac-go/pkg/core/board"
)

const (
	defaultBoardSize = 3
	defaultChainLen  = 3
)

type flags struct {
	width    *int
	height   *int
	chainLen *int
}

func parseFlags() *flags {
	result := &flags{}
	result.width = flag.Int("width", defaultBoardSize, "width of the board")
	result.height = flag.Int("height", defaultBoardSize, "height of the board")
	result.chainLen = flag.Int("chain-len", defaultChainLen, "length of the chain to win")
	flag.Parse()

	return result
}

func main() {
	f := parseFlags()

	fmt.Println(
		board.Create(
			*f.width,
			*f.height,
			*f.chainLen,
		),
	)
}
