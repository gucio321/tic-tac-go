package ttgletter

// Letter represents board letters (x and o)
type Letter byte

// Letter types
const (
	LetterNone Letter = iota
	LetterX
	LetterO
)

// NewLetter creates a new letter
func NewLetter() *Letter {
	result := LetterNone

	return &result
}

// SetState sets index state
func (l *Letter) SetState(state Letter) {
	*l = state
}

// String returns Letter's string
func (l *Letter) String() string {
	switch *l {
	case LetterNone:
		return " "
	case LetterX:
		return "X"
	case LetterO:
		return "O"
	}

	// should not be reached
	return "?"
}

// IsNone return's true if Letter == LetterNone
func (l *Letter) IsNone() bool {
	return *l == LetterNone
}

// Opposite returns letter, which is an opposite to current
func (l *Letter) Opposite() Letter {
	switch *l {
	case LetterX:
		return LetterO
	case LetterO:
		return LetterX
	}

	return LetterNone
}
