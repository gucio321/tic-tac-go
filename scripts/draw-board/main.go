package main

import (
	"flag"
	"fmt"

	"github.com/gucio321/tic-tac-go/pkg/core/board"
)

type flags struct {
	width    *int
	height   *int
	chainLen *int
}

func parseFlags() *flags {
	result := &flags{}
	result.width = flag.Int("width", 3, "width of the board")
	result.height = flag.Int("height", 3, "height of the board")
	result.chainLen = flag.Int("chain-len", 3, "length of the chain to win")
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
