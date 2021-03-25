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

	t.println(sep)
}

func (t *TTT) printBoard() {
	ttgcommon.Clear()

	t.printSeparator()

	for i := 0; i < t.height; i++ {
		line := "| "
		for j := 0; j < t.width; j++ {
			line += t.board[i][j].String()
			line += " | "
		}

		t.println(line)
		t.printSeparator()
	}
}

func (t *TTT) getPlayerMove() (x, y int) {
	for {
		t.print(fmt.Sprintf("Enter your move (1-%d): ", t.width*t.height))

		text, err := t.reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		if text == "" {
			t.println("please enter number from 1 to 9")

			continue
		}

		text = strings.ReplaceAll(text, string('\n'), "")
		text = strings.ReplaceAll(text, "\r", "")

		num, err := strconv.Atoi(text)
		if err != nil {
			t.println("invalid index number")

			continue
		}

		if num <= 0 || num > t.width*t.height {
			t.println(fmt.Sprintf("You must enter number from 1 to %d", t.width*t.height))

			continue
		}

		num--

		x, y = ttgcommon.IntToCords(t.width, t.height, num)

		if t.board[y][x].IsNone() {
			return
		}

		t.println("This index is busy")
	}
}

func (t *TTT) pressAnyKeyPrompt() {
	t.print("\nPress any key to continue...")

	_, _ = t.reader.ReadString('\n')
}

func (t *TTT) print(msg ...string) {
	for _, m := range msg {
		fmt.Print(m)
	}
}

func (t *TTT) println(msg ...string) {
	t.print(msg...)
	fmt.Printf("\n")
}
