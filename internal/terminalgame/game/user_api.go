package game

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/gucio321/terminalmenu/pkg/menuutils"
)

func (t *TTG) getPlayerMove() (i int) {
	for {
		num, err := menuutils.GetNumber(fmt.Sprintf("Enter your move (1-%d): ", t.Board().Width()*t.Board().Height()))

		switch {
		case err == nil:
			// noop
		case errors.Is(err, strconv.ErrSyntax):
			if readErr := menuutils.PromptEnter("Please enter correct number"); readErr != nil {
				log.Fatal(readErr)
			}

			continue
		default:
			log.Fatal(err)
		}

		if w, h := t.Board().Width(), t.Board().Height(); num <= 0 || num > w*h {
			if err := menuutils.PromptEnter(fmt.Sprintf("You must enter number from 1 to %d", w*h)); err != nil {
				log.Fatal(err)
			}

			continue
		}

		num--

		if t.Board().IsIndexFree(num) {
			return num
		}

		if err := menuutils.PromptEnter("This index is busy"); err != nil {
			log.Fatal(err)
		}
	}
}
