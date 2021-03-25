package ttgcommon

import (
	"testing"
)

func Test_IntToCords(t *testing.T) {
	// standard 3x3 board
	/*
		+---+---+---+
		| 0 | 1 | 2 |
		+---+---+---+
		| 3 | 4 | 5 |
		+---+---+---+
		| 6 | 7 | 8 |
		+---+---+---+
	*/
	w, h := 3, 3
	// so index 3 should have y = 1, x = 0
	i := 3
	x, y := IntToCords(w, h, i)

	if y != 1 || x != 0 {
		t.Fatalf("IntToCords(%d, %d, %d) returned unexpected values x: %d, y: %d", w, h, i, x, y)
	}

	// index 7 should have y = 2, x = 1
	i = 7
	x, y = IntToCords(w, h, i)

	if y != 2 || x != 1 {
		t.Fatalf("IntToCords(%d, %d, %d) returned unexpected values x: %d, y: %d", w, h, i, x, y)
	}
}

func Test_SplitIntoLinesWithMaxWidth(t *testing.T) {
	testString := "this is a long string, which need to be splited"
	l := 22
	splited := SplitIntoLinesWithMaxWidth(testString, l)

	if len(splited) != 3 {
		t.Fatal("String wasn't splited correctly")
	}

	for i := range splited {
		if len(splited[i]) > l {
			t.Fatal("String wasn't splited correctly")
		}
	}
}
