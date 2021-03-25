package ttgboard

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gucio321/tic-tac-go/ttgcommon"
)

func (t *TTT) printSeparator() {
	sep := "+"
	for i := 0; i < t.width; i++ {
		sep += "---+"
	}

	fmt.Println(sep)
}

func (t *TTT) printBoard() {
	ttgcommon.Clear()

	t.printSeparator()

	for y := 0; y < t.height; y++ {
		line := "| "

		for x := 0; x < t.width; x++ {
			i := ttgcommon.CordsToInt(t.width, t.height, x, y)
			line += (*t.board)[i].String()
			line += " | "
		}

		fmt.Println(line)
		t.printSeparator()
	}
}

func (t *TTT) getPlayerMove() (i int) {
	for {
		fmt.Printf("Enter your move (1-%d): ", t.width*t.height)

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

		if num <= 0 || num > t.width*t.height {
			println(fmt.Sprintf("You must enter number from 1 to %d", t.width*t.height))

			continue
		}

		num--

		if t.board.isIndexFree(num) {
			return num
		}

		println("This index is busy")
	}
}

func (t *TTT) pressAnyKeyPrompt() {
	print("\nPress any key to continue...")

	_, _ = t.reader.ReadString('\n')
}
