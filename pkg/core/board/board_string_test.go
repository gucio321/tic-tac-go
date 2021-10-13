package board

import (
	"fmt"
	"testing"

	"github.com/gucio321/tic-tac-go/pkg/core/board/letter"
	"github.com/stretchr/testify/assert"
)

func Test_Board_separator(t *testing.T) {
	tests := []struct {
		name     string
		width    int
		expected string
	}{
		{"one", 1, "+---+"},
		{"three (the most standard)", 3, "+---+---+---+"},
		{"four", 4, "+---+---+---+---+"},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			b := Create(test.width, 1, 1)
			assert.Equal(tt, test.expected, b.separator())
		})
	}
}

func Test_Board_String(t *testing.T) {
	tests := []struct {
		name     string
		board    *Board
		expected string
	}{
		{
			"empty 3x3 board", &Board{
				width:  3,
				height: 3,
				board: []letter.Letter{
					0, 0, 0,
					0, 0, 0,
					0, 0, 0,
				},
			},
			fmt.Sprint(
				"+---+---+---+\n",
				"|   |   |   |\n",
				"+---+---+---+\n",
				"|   |   |   |\n",
				"+---+---+---+\n",
				"|   |   |   |\n",
				"+---+---+---+\n",
			),
		},
		{
			"filled 3x3 board", &Board{
				width:  3,
				height: 3,
				board: []letter.Letter{
					1, 2, 0,
					0, 2, 2,
					0, 1, 1,
				},
			},
			fmt.Sprint(
				"+---+---+---+\n",
				"| X | O |   |\n",
				"+---+---+---+\n",
				"|   | O | O |\n",
				"+---+---+---+\n",
				"|   | X | X |\n",
				"+---+---+---+\n",
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			assert.Equal(tt, test.expected, test.board.String())
		})
	}
}
