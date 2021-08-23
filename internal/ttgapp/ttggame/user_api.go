package ttggame

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func (t *TTG) getPlayerMove() (i int) {
	for {
		fmt.Printf("Enter your move (1-%d): ", t.Board().Width()*t.Board().Height())

		text, err := t.reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		if text == "" {
			println("please enter number from 1 to 9")

			continue
		}

		text = strings.ReplaceAll(text, string('\n'), "")
		text = strings.ReplaceAll(text, "\r", "")

		num, err := strconv.Atoi(text)
		if err != nil {
			println("invalid index number")

			continue
		}

		if w, h := t.Board().Width(), t.Board().Height(); num <= 0 || num > w*h {
			println(fmt.Sprintf("You must enter number from 1 to %d", w*h))

			continue
		}

		num--

		if t.Board().IsIndexFree(num) {
			return num
		}

		println("This index is busy")
	}
}

func (t *TTG) pressAnyKeyPrompt() {
	print("\nPress any key to continue...")

	_, _ = t.reader.ReadString('\n')
}
