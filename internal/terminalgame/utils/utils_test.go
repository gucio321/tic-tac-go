package utils

import (
	"testing"
)

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
