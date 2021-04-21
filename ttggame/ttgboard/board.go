package ttgboard

import (
	"github.com/gucio321/tic-tac-go/ttgcommon"
	"github.com/gucio321/tic-tac-go/ttggame/ttgletter"
)

// Board represents game board
type Board struct {
	board                   []*ttgletter.Letter
	width, height, chainLen int
}

// NewBoard creates a new board
func NewBoard(w, h, chainLen int) *Board {
	result := &Board{
		board:    make([]*ttgletter.Letter, w*h),
		width:    w,
		height:   h,
		chainLen: chainLen,
	}

	for i := range result.board {
		result.board[i] = ttgletter.NewLetter()
	}

	return result
}

// SetIndexState set index's state
func (b *Board) SetIndexState(i int, state ttgletter.Letter) {
	b.board[i].SetState(state)
}

// GetIndexState returns index's state
func (b *Board) GetIndexState(i int) ttgletter.Letter {
	return *b.board[i]
}

// IsIndexFree returns true, if state of index given is None
func (b *Board) IsIndexFree(i int) bool {
	return b.board[i].IsNone()
}

// Copy returns board copy
func (b *Board) Copy() *Board {
	newBoard := NewBoard(b.width, b.height, b.chainLen)
	for i := range newBoard.board {
		newBoard.SetIndexState(i, b.GetIndexState(i))
	}

	return newBoard
}

func (b *Board) Cut(w, h int) *Board {
	result := NewBoard(w, h, b.chainLen)
	for i := range result.board {
		result.SetIndexState(i, b.GetIndexState(ttgcommon.ConvertIndex(w, h, b.width, b.height, i)))
	}
	return result
}

func (b *Board) separator() string {
	sep := "+"
	for i := 0; i < b.width; i++ {
		sep += "---+"
	}

	return sep
}

func (b *Board) String() string {
	ttgcommon.Clear()

	s := b.separator()
	s += "\n"

	for y := 0; y < b.height; y++ {
		line := "| "

		for x := 0; x < b.width; x++ {
			i := ttgcommon.CordsToInt(b.width, b.height, x, y)
			line += b.board[i].String()
			line += " | "
		}

		s += line + "\n" + b.separator() + "\n"
	}

	return s
}
