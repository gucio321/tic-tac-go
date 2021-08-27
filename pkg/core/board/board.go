package board

import (
	"fmt"

	"github.com/gucio321/tic-tac-go/pkg/core/board/letter"
)

// Board represents game board.
type Board struct {
	board                   []*letter.Letter
	width, height, chainLen int
}

// Create creates a new board.
func Create(w, h, chainLen int) *Board {
	result := &Board{
		board:    make([]*letter.Letter, w*h),
		width:    w,
		height:   h,
		chainLen: chainLen,
	}

	for i := range result.board {
		result.board[i] = letter.Create()
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
func (b *Board) SetIndexState(i int, state letter.Letter) {
	if i > len(b.board) || i < 0 {
		panic(fmt.Sprintf("Tic-Tac-Go: board.(*Board).SetIndexState: index %d out of range", i))
	}

	b.board[i].SetState(state)
}

// GetIndexState returns index's state.
func (b *Board) GetIndexState(i int) letter.Letter {
	if i > len(b.board) || i < 0 {
		panic(fmt.Sprintf("Tic-Tac-Go: board.(*Board).GetIndexState: index %d out of range", i))
	}

	return *b.board[i]
}

// IsIndexFree returns true, if state of index given is None.
func (b *Board) IsIndexFree(i int) bool {
	if i > len(b.board) || i < 0 {
		panic(fmt.Sprintf("Tic-Tac-Go: board.(*Board).IsIndexFree: index %d out of range", i))
	}

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
		panic(fmt.Sprintf("Tic-Tac-Go: board.(*Board).Cut: cannot cut larger board from smaller: original board size is %dx%d, requested - %dx%d",
			b.width, b.height, w, h))
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
			i := b.CordsToInt(x, y)
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

// IsWinner returns true if the 'player' is a winner.
// In addition IsWinner will return a list of winning combination.
func (b *Board) IsWinner(player letter.Letter) (ok bool, i []int) {
	combos := b.GetWinBoard(b.chainLen)

	for _, i := range combos {
		line := 0

		for _, c := range i {
			if b.GetIndexState(c) == player {
				line++
			}
		}

		if line == b.chainLen {
			return true, i
		}
	}

	return false, nil
}

// IntToCords converts intager to X-Y cords.
func (b *Board) IntToCords(i int) (x, y int) {
	if i < 0 || i > b.width*b.height {
		panic("Tic-Tac-Go: board(*Board).IntToCords: index out of range")
	}

	for {
		if i-b.width >= 0 {
			y++

			i -= b.width
		} else {
			x = i

			break
		}
	}

	return
}

// CordsToInt converts coordinates on board to board index.
func (b *Board) CordsToInt(x, y int) int {
	if x*y < 0 || x > b.width || y > b.height {
		panic("Tic-Tac-Go: board.(*Board).CordsToInt: index out of range")
	}

	return y*b.width + x
}
