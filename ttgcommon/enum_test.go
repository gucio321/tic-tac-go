package ttgcommon

import (
	"testing"
)

func Test_GetWinBoard(t *testing.T) {
	w, h, l := 3, 3, 3
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

	combinations := GetWinBoard(w, h, l)

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
	w, h := 4, 4
	// corners of board 4x4 should be: 0, 3, 11, 15
	correctCorners := []int{0, 3, 12, 15}

	corners := GetCorners(w, h)

	if len(corners) != len(correctCorners) {
		t.Fatal("Unexpected board corners returned")
	}

	for i := range corners {
		if correctCorners[i] != corners[i] {
			t.Fatal("Unexpected board corners returned")
		}
	}
}

func Test_GetMiddles(t *testing.T) {
	correctMiddles := []int{1, 2, 4, 7, 8, 11, 13, 14}
	w, h := 4, 4
	middles := GetMiddles(w, h)

	if len(middles) != len(correctMiddles) {
		t.Fatal("invalid board middles returned")
	}

	for i := range middles {
		if middles[i] != correctMiddles[i] {
			t.Fatal("invalid board middles returned")
		}
	}
}

func Test_GetCenterCorrectBoard(t *testing.T) {
	w, h := 3, 3
	correctCenter := []int{4}
	center := GetCenter(w, h)

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
	w, h := 4, 4
	center := GetCenter(w, h)

	if len(center) > 0 {
		t.Fatal("Unexpected board center returned")
	}
}

func Test_ConvertIndex(t *testing.T) {
	fw, fh := 3, 3
	rw, rh := 7, 7
	idx := 4       // the center of 3x3 board
	expected := 24 // should be 24 on 7x7 board
	returned := ConvertIndex(fw, fh, rw, rh, idx)

	if returned != expected {
		t.Fatalf("Returned index isn't equal to expected (%d != %d)", returned, expected)
	}

	rw, rh = 5, 5
	expected = 12 // should be 12 on 5x5 board
	returned = ConvertIndex(fw, fh, rw, rh, idx)

	if returned != expected {
		t.Fatalf("Returned index isn't equal to expected (%d != %d)", returned, expected)
	}

	fw, fh = 2, 2
	rw, rh = 4, 4
	idx = 2
	expected = 9
	returned = ConvertIndex(fw, fh, rw, rh, idx)

	if returned != expected {
		t.Fatalf("Returned index isn't equal to expected (%d != %d)", returned, expected)
	}

	fw, fh = 2, 3
	rw, rh = 4, 5
	idx = 3
	expected = 10
	returned = ConvertIndex(fw, fh, rw, rh, idx)

	if returned != expected {
		t.Fatalf("Returned index isn't equal to expected (%d != %d)", returned, expected)
	}
}
