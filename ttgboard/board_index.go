package ttgboard

// IdxState represents index's state
type Letter byte

const (
	LetterNone Letter = iota
	LetterX
	LetterO
)

func newBoardIndex() *Letter {
	result := LetterNone

	return &result
}

// SetState sets index state
func (b *Letter) SetState(state Letter) {
	*b = state
}

func (b *Letter) String() string {
	switch *b {
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

// IsFree return's true if index is free
func (b *Letter) IsNone() bool {
	return *b == LetterNone
}
