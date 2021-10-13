package letter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_String(t *testing.T) {
	tests := []struct {
		id     string
		letter Letter
		str    string
	}{
		{"letter X", LetterX, "X"},
		{"letter O", LetterO, "O"},
		{"letter None", LetterNone, " "},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(tst *testing.T) {
			assert.Equal(tst, tt.letter.String(), tt.str, "incorrect letter string")
		})
	}
}

func Test_Letter_incorrect_letter(t *testing.T) {
	assert.Panics(t, func() { _ = Letter(5).String() }, "Getting string of incorrect letter didn't panicked")
}

func Test_Opposite(t *testing.T) {
	tests := []struct {
		id       string
		letter   Letter
		opposite Letter
	}{
		{"letter none", LetterNone, LetterNone},
		{"letter X", LetterX, LetterO},
		{"letter O", LetterO, LetterX},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(tst *testing.T) {
			assert.Equal(tst, tt.opposite, tt.letter.Opposite(), "unexpected opposite letter.")
		})
	}
}
