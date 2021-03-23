package ttgcommon

// BoardW, BoardH are board's width and height
const (
	BaseBoardW = 3
	BaseBoardH = 3
)

// GetWinBoard returns winning indexes list
func GetWinBoard(w, h, l int) [][]int {
	// for w = h:
	// n = (w-l+1)*h + (h-l+1) * w + 2 * ((w or h)-l+1)
	// generally (if s = w = h) n = (s-l+1)*s + (s-l+1) *w + 2 * (s - l + 1)
	winningIndexes := make([][]int, 0)

	// horizontal indexes
	for row := 0; row < h; row++ {
		for rowIdx := 0; rowIdx+l <= w; rowIdx++ {
			line := make([]int, l)
			for idx := range line {
				line[idx] = row*w + rowIdx + idx
			}

			winningIndexes = append(winningIndexes, line)
		}
	}

	// vertical indexes
	for col := 0; col < w; col++ {
		for colIdx := 0; colIdx+l <= h; colIdx++ {
			line := make([]int, l)
			for idx := range line {
				line[idx] = col + colIdx*w + idx*w
			}

			winningIndexes = append(winningIndexes, line)
		}
	}

	for x := 0; x < h; x++ {
		for xIdx := 0; (x*w+xIdx*w+xIdx)+((l-1)*w+(l-1)) <= h*w-1; xIdx++ {
			line := make([]int, l)
			for idx := range line {
				line[idx] = (x*w + xIdx) + (idx*w + idx)
			}

			winningIndexes = append(winningIndexes, line)
		}
	}

	for bx := 0; bx < h; bx++ {
		for bxIdx := 0; bx*w+(bxIdx+l)*w <= h*w-1; bxIdx++ {
			line := make([]int, l)
			for idx := range line {
				line[idx] = (bx * w) + (bxIdx+idx)*w + (l - idx - 1)
			}

			winningIndexes = append(winningIndexes, line)
		}
	}

	return winningIndexes
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
