package ttgcommon

// BoardW, BoardH are board's width and height
const (
	BaseBoardW = 3
	BaseBoardH = 3
)

// GetWinBoard returns winning indexes list
func GetWinBoard() [8][3]int {
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
