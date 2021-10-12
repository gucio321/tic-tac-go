package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetWinBoard(t *testing.T) {
	w, h, l := 3, 3, 3
	board := Create(w, h, l)
	correctCombinations := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},

		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},

		{0, 4, 8},
		{2, 4, 6},
	}

	combinations := board.GetWinBoard(l)

	assert.Equal(t, correctCombinations, combinations, "Unexpected combos returned")
}

func Test_GetCorners(t *testing.T) {
	w, h, l := 4, 4, 3
	board := Create(w, h, l)

	// corners of board 4x4 should be: 0, 3, 11, 15
	correctCorners := []int{0, 3, 12, 15}

	corners := board.GetCorners()

	assert.Equal(t, correctCorners, corners, "wrong corner values returned")
}

func Test_GetOppositeCorner(t *testing.T) {
	a := assert.New(t)
	w, h, l := 3, 3, 3
	board := Create(w, h, l)

	a.Equal(8, board.GetOppositeCorner(0), "unexpected corner returned")
	a.Equal(2, board.GetOppositeCorner(6), "unexpected corner returned")
	a.Panics(func() { board.GetOppositeCorner(1) }, "non-corner passed but GetOppositeCorner didn't panicked")
}

func Test_GetSides(t *testing.T) {
	expected := []int{1, 2, 4, 7, 8, 11, 13, 14}
	w, h, l := 4, 4, 3
	board := Create(w, h, l)
	sides := board.GetSides()

	assert.Equal(t, expected, sides, "unexpected sides values returned")
}

func Test_GetCenterCorrectBoard(t *testing.T) {
	w, h, l := 3, 3, 3
	board := Create(w, h, l)

	expected := []int{4}
	center := board.GetCenter()

	assert.Equal(t, expected, center, "unexpected center value")
}

func Test_GetCenterWrongBoard(t *testing.T) {
	// 4x4 board doesn't have any center
	w, h, l := 4, 4, 3
	board := Create(w, h, l)
	center := board.GetCenter()

	assert.NotNil(t, center, "GetCenter on incorrect board returned nil")
	assert.Equal(t, center, []int{}, "GetCenter on incorrect board didn't returned correct value")
}

func Test_ConvertIndex(t *testing.T) {
	tests := []struct {
		name               string
		chainLen           int
		boardW, boardH     int
		fictionW, fictionH int
		index              int
		expectedValue      int
	}{
		{
			"Test 1",
			3,
			7, 7,
			3, 3,
			4,
			24,
		},
		{
			"Test 2",
			3,
			5, 5,
			3, 3,
			4,
			12,
		},
		{
			"Test 3",
			3,
			4, 4,
			2, 2,
			2,
			9,
		},
		{
			"Test 3",
			3,
			4, 5,
			2, 3,
			3,
			10,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			board := Create(test.boardW, test.boardH, test.chainLen)
			assert.Equal(tt, test.expectedValue, board.ConvertIndex(test.fictionW, test.fictionH, test.index))
		})
	}
}

func Test_ConvertIndex_incorrect_convertions(t *testing.T) {
	tests := []struct {
		name string
		realW, realH,
		fictionW, fictionH int
	}{
		{"convert to larger board 1", 3, 3, 5, 3},
		{"convert to larger board 2", 3, 3, 9, 9},
		{"odd and even numbers", 4, 4, 3, 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			assert.Panics(tt, func() { Create(test.realW, test.realH, 3).ConvertIndex(test.fictionW, test.fictionH, 0) },
				"invalid conversion didn't panicked")
		})
	}
}

func Test_IsEdgeIndex(t *testing.T) {
	w, h, l := 3, 3, 3
	board := Create(w, h, l)
	i := 4

	if board.IsEdgeIndex(i) {
		t.Fatal("center of 3x3 board isn't edge")
	}

	i = 2
	if !board.IsEdgeIndex(i) {
		t.Fatal("invalid edge value")
	}

	i = 5
	if !board.IsEdgeIndex(i) {
		t.Fatal("invalid edge value")
	}

	i = 6
	if !board.IsEdgeIndex(i) {
		t.Fatal("invalid edge value")
	}

	i = 8
	if !board.IsEdgeIndex(i) {
		t.Fatal("invalid edge value")
	}
}
