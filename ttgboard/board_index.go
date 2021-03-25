package ttgboard

// IdxState represents index's state
type IdxState int

const (
	IdxNone IdxState = iota
	IdxX
	IdxO
)

func (i IdxState) String() string {
	switch i {
	case IdxNone:
		return " "
	case IdxX:
		return "X"
	case IdxO:
		return "O"
	}

	return "?"
}

// BoardIndex represents board index
type BoardIndex struct {
	state IdxState
}

func newIndex() *BoardIndex {
	result := &BoardIndex{
		state: IdxNone,
	}

	return result
}

// SetState sets index state
func (b *BoardIndex) SetState(state IdxState) {
	b.state = state
}

func (b *BoardIndex) String() string {
	switch b.state {
	case IdxNone:
		return " "
	case IdxX:
		return "X"
	case IdxO:
		return "O"
	}

	// should not be reached
	return "?"
}

// IsFree return's true if index is free
func (b BoardIndex) IsFree() bool {
	return b.state == IdxNone
}
