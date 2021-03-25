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
	// for now, GetMiddles works only for 3x3 board
	correctMiddles := []int{1, 3, 5, 7}
	middles := GetMiddles()

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
