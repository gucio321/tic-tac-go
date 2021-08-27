package letter

// Letter represents board letters (x and o).
type Letter byte

// Letter types.
const (
	LetterNone Letter = iota
	LetterX
	LetterO
)

// String returns Letter's string.
func (l Letter) String() string {
	lookup := map[Letter]string{
		LetterNone: " ",
		LetterX:    "X",
		LetterO:    "O",
	}

	result, ok := lookup[l]
	if !ok {
		panic("Tic-Tac-Go: letter.(*Letter).String: unexpected letter value")
	}

	return result
}

// Opposite returns letter, which is an opposite to current.
func (l Letter) Opposite() Letter {
	switch l {
	case LetterX:
		return LetterO
	case LetterO:
		return LetterX
	}

	return LetterNone
}
