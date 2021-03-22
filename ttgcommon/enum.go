package ttgcommon

import (
	"fmt"
)

// BoardW, BoardH are board's width and height
const (
	BaseBoardW = 3
	BaseBoardH = 3
)

// GetWinBoard returns winning indexes list
func GetWinBoard(w, h, l int) [8][3]int {
	// for w = h:
	// n = (w-l+1)*h + (h-l+1) * w + 2 * ((w or h)-l+1)
	// generally (if s = w = h) n = (s-l+1)*s + (s-l+1) *w + 2 * (s - l + 1)
	w, h = 3, 4
	numberOfCombinations := (w-l+1)*h + (h-l+1)*w + 2*(w-l+1)
	winningIndexes := make([][]int, numberOfCombinations)
	for n := range winningIndexes {
		winningIndexes[n] = make([]int, l)
	}

	// horizontal indexes
	idx := 0
	for row := 0; row < h; row++ {
		for rowIdx := 0; rowIdx+l <= w; rowIdx++ {
			line := make([]int, l)
			for idx := range line {
				line[idx] = row*w + rowIdx + idx
			}
			winningIndexes[idx] = line
			idx++
		}
	}

	// vertical indexes
	for col := 0; col < w; col++ {
		for colIdx := 0; colIdx+l <= h; colIdx++ {
			line := make([]int, l)
			for idx := range line {
				line[idx] = col + colIdx*w + idx*w
			}
			winningIndexes[idx] = line
			idx++
		}
	}

	fmt.Println("indexes", winningIndexes)

	return [8][3]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},

		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},

		{0, 4, 8},
		{2, 4, 6},
	}
}

// GetCorners returns board's corners
func GetCorners() (result [4]struct{ X, Y int }) {
	result = [4]struct{ X, Y int }{
		{0, 0},
		{0, 2},
		{2, 0},
		{2, 2},
	}

	return
}

// GetMiddles returns middles of board's edges
func GetMiddles() (result [4]struct{ X, Y int }) {
	result = [4]struct{ X, Y int }{
		{1, 0},
		{0, 1},
		{2, 1},
		{1, 2},
	}

	return
}
