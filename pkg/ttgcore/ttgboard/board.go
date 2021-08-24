package ttgboard

import (
	"log"

	"github.com/gucio321/tic-tac-go/pkg/ttgcommon"
	"github.com/gucio321/tic-tac-go/pkg/ttgcore/ttgletter"
)

// Board represents game board.
type Board struct {
	board                   []*ttgletter.Letter
	width, height, chainLen int
}

// Create creates a new board.
func Create(w, h, chainLen int) *Board {
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

// Width returns board's width.
func (b *Board) Width() int {
	return b.width
}

// Height returns board's height.
func (b *Board) Height() int {
	return b.height
}

// ChainLength returns length of chain coded in board.
func (b *Board) ChainLength() int {
	return b.chainLen
}

// SetIndexState set index's state.
func (b *Board) SetIndexState(i int, state ttgletter.Letter) {
	b.board[i].SetState(state)
}

// GetIndexState returns index's state.
func (b *Board) GetIndexState(i int) ttgletter.Letter {
	return *b.board[i]
}

// IsIndexFree returns true, if state of index given is None.
func (b *Board) IsIndexFree(i int) bool {
	return b.board[i].IsNone()
}

// Copy returns board copy.
func (b *Board) Copy() *Board {
	newBoard := Create(b.width, b.height, b.chainLen)
	for i := range newBoard.board {
		newBoard.SetIndexState(i, b.GetIndexState(i))
	}

	return newBoard
}

// Cut cuts a smaller board from a larger.
func (b *Board) Cut(w, h int) *Board {
	if w > b.width || h > b.height {
		log.Fatal("cannot cat larger board from smaller")
	}

	result := Create(w, h, b.chainLen)
	for i := range result.board {
		result.SetIndexState(i, b.GetIndexState(b.ConvertIndex(w, h, i)))
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

// IsBoardFull returns true, if there is no more space on the board.
func (b *Board) IsBoardFull() bool {
	for _, i := range b.board {
		if i.IsNone() {
			return false
		}
	}

	return true
}

// IsWinner returns true if the 'player' is a winner using specified 'chainLen'.
func (b *Board) IsWinner(chainLen int, player ttgletter.Letter) (ok bool, i []int) {
	combos := b.GetWinBoard(chainLen)

	for _, i := range combos {
		line := 0

		for _, c := range i {
			if b.GetIndexState(c) == player {
				line++
			}
		}

		if line == chainLen {
			return true, i
		}
	}

	return false, nil
}
