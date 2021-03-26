package ttgboard

import (
	"github.com/gucio321/tic-tac-go/ttggame/ttgletter"
)

// Board represents game board
type Board []*ttgletter.Letter

// NewBoard creates a new board
func NewBoard(size int) *Board {
	result := &Board{}
	*result = make([]*ttgletter.Letter, size)

	for i := range *result {
		(*result)[i] = ttgletter.NewLetter()
	}

	return result
}

// SetIndexState set index's state
func (b *Board) SetIndexState(i int, state ttgletter.Letter) {
	(*b)[i].SetState(state)
}

// GetIndexState returns index's state
func (b *Board) GetIndexState(i int) ttgletter.Letter {
	return *(*b)[i]
}

// IsIndexFree returns true, if state of index given is None
func (b *Board) IsIndexFree(i int) bool {
	return (*b)[i].IsNone()
}

// Copy returns board copy
func (b *Board) Copy() *Board {
	newBoard := NewBoard(len(*b))
	for i := range *newBoard {
		*(*newBoard)[i] = *(*b)[i]
	}

	return newBoard
}
