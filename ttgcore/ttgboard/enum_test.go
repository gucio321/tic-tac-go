package ttgboard

import (
	"testing"
)

func Test_GetWinBoard(t *testing.T) {
	w, h, l := 3, 3, 3
	board := NewBoard(w, h, l)
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

	if len(correctCombinations) != len(combinations) {
		t.Fatal("Unexpected board returned")
	}

	for i := range combinations {
		if len(correctCombinations[i]) != len(combinations[i]) {
			t.Fatal("Unexpected board returned")
		}

		for j := range combinations[i] {
			if correctCombinations[i][j] != combinations[i][j] {
				t.Fatal("Unexpected board returned")
			}
		}
	}
}

func Test_GetCorners(t *testing.T) {
	w, h, l := 4, 4, 3
	board := NewBoard(w, h, l)

	// corners of board 4x4 should be: 0, 3, 11, 15
	correctCorners := []int{0, 3, 12, 15}

	corners := board.GetCorners()

	if len(corners) != len(correctCorners) {
		t.Fatal("Unexpected board corners returned")
	}

	for i := range corners {
		if correctCorners[i] != corners[i] {
			t.Fatal("Unexpected board corners returned")
		}
	}
}

func Test_GetOppositeCorner(t *testing.T) {
	w, h, l := 3, 3, 3
	board := NewBoard(w, h, l)

	expected := 8
	given := board.GetOppositeCorner(0)

	if expected != given {
		t.Fatal("Unexpected corner value returned")
	}

	expected = 2
	given = board.GetOppositeCorner(6)

	if expected != given {
		t.Fatal("Unexpected corner value returned")
	}
}

func Test_GetMiddles(t *testing.T) {
	correctMiddles := []int{1, 2, 4, 7, 8, 11, 13, 14}
	w, h, l := 4, 4, 3
	board := NewBoard(w, h, l)
	sides := board.GetSides()

	if len(sides) != len(correctMiddles) {
		t.Fatal("invalid board middles returned")
	}

	for i := range sides {
		if sides[i] != correctMiddles[i] {
			t.Fatal("invalid board middles returned")
		}
	}
}

func Test_GetCenterCorrectBoard(t *testing.T) {
	w, h, l := 3, 3, 3
	board := NewBoard(w, h, l)
	correctCenter := []int{4}
	center := board.GetCenter()

	if len(center) != len(correctCenter) {
		t.Fatal("Unexpected board center returned")
	}

	for i := range center {
		if center[i] != correctCenter[i] {
			t.Fatal("Unexpected board center returned")
		}
	}
}

func Test_GetCenterWrongBoard(t *testing.T) {
	// 4x4 board doesn't have any center
	w, h, l := 4, 4, 3
	board := NewBoard(w, h, l)
	center := board.GetCenter()

	if len(center) > 0 {
		t.Fatal("Unexpected board center returned")
	}
}

func Test_ConvertIndex(t *testing.T) {
	l := 3

	fw, fh := 3, 3
	rw, rh := 7, 7
	board := NewBoard(rw, rh, l)
	idx := 4       // the center of 3x3 board
	expected := 24 // should be 24 on 7x7 board
	returned := board.ConvertIndex(fw, fh, idx)

	if returned != expected {
		t.Fatalf("Returned index isn't equal to expected (%d != %d)", returned, expected)
	}

	rw, rh = 5, 5
	board = NewBoard(rw, rh, l)
	expected = 12 // should be 12 on 5x5 board
	returned = board.ConvertIndex(fw, fh, idx)

	if returned != expected {
		t.Fatalf("Returned index isn't equal to expected (%d != %d)", returned, expected)
	}

	fw, fh = 2, 2
	rw, rh = 4, 4
	board = NewBoard(rw, rh, l)
	idx = 2
	expected = 9
	returned = board.ConvertIndex(fw, fh, idx)

	if returned != expected {
		t.Fatalf("Returned index isn't equal to expected (%d != %d)", returned, expected)
	}

	fw, fh = 2, 3
	rw, rh = 4, 5
	board = NewBoard(rw, rh, l)
	idx = 3
	expected = 10
	returned = board.ConvertIndex(fw, fh, idx)

	if returned != expected {
		t.Fatalf("Returned index isn't equal to expected (%d != %d)", returned, expected)
	}
}

func Test_IsEdgeIndex(t *testing.T) {
	w, h, l := 3, 3, 3
	board := NewBoard(w, h, l)
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
