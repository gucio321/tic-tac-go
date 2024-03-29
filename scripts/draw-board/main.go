// Package main - this package could be used to draw a game board
// specified with command-line arguments
package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"

	"github.com/gucio321/tic-tac-go/pkg/core/board/letter"

	"github.com/gucio321/tic-tac-go/pkg/core/board"
)

const (
	defaultBoardSize = 3
	defaultChainLen  = 3
)

var errIncorrectLetter = errors.New("incorrect number")

var _ flag.Value = &lettersValue{}

type lettersValue []int

func (l *lettersValue) String() string {
	return fmt.Sprintf("%v", *l)
}

func (l *lettersValue) Set(value string) (err error) {
	for _, letter := range value {
		if letter < '0' || letter > '9' {
			return fmt.Errorf("%w %v", errIncorrectLetter, letter)
		}
	}

	x, err := strconv.Atoi(value)
	if err != nil {
		return fmt.Errorf("converting string to int: %w", err)
	}

	*l = append(*l, x)

	return nil
}

type flags struct {
	width    *int
	height   *int
	chainLen *int
	x        lettersValue
	o        lettersValue
}

func parseFlags() *flags {
	result := &flags{}
	result.width = flag.Int("width", defaultBoardSize, "width of the board")
	result.height = flag.Int("height", defaultBoardSize, "height of the board")
	result.chainLen = flag.Int("chain-len", defaultChainLen, "length of the chain to win")
	flag.Var(&result.x, "x", "indexes with X")
	flag.Var(&result.o, "o", "indexes with O")
	flag.Parse()

	return result
}

func main() {
	f := parseFlags()

	b := board.Create(
		*f.width,
		*f.height,
		*f.chainLen,
	)

	for _, x := range f.x {
		b.SetIndexState(x, letter.LetterX)
	}

	for _, o := range f.o {
		b.SetIndexState(o, letter.LetterO)
	}

	fmt.Println(b)
}
