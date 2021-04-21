package ttggame

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gucio321/tic-tac-go/ttgcommon"
	"github.com/gucio321/tic-tac-go/ttggame/ttgboard"
	"github.com/gucio321/tic-tac-go/ttggame/ttgletter"
	"github.com/gucio321/tic-tac-go/ttggame/ttgplayer"
)

// TTG represents TicTacToe game
type TTG struct {
	board   *ttgboard.Board
	reader  *bufio.Reader
	player1 *ttgplayer.Player
	player2 *ttgplayer.Player
	width,
	height,
	chainLen int
}

// NewTTG creates a ne TTG
func NewTTG(w, h, chainLen byte, player1Type, player2Type ttgplayer.PlayerType) *TTG {
	result := &TTG{
		board:    ttgboard.NewBoard(int(w), int(h), int(chainLen)),
		reader:   bufio.NewReader(os.Stdin),
		width:    int(w),
		height:   int(h),
		chainLen: int(chainLen),
	}

	player1Letter := ttgletter.LetterX
	player2Letter := ttgletter.LetterO

	switch player1Type {
	case ttgplayer.PlayerPC:
		result.player1 = ttgplayer.NewPlayer(player1Type, player1Letter,
			func() (i int) {
				i = result.getPCMove(player1Letter)
				return i
			},
		)
	case ttgplayer.PlayerPerson:
		result.player1 = ttgplayer.NewPlayer(player1Type, player1Letter, result.getPlayerMove)
	}

	switch player2Type {
	case ttgplayer.PlayerPC:
		result.player2 = ttgplayer.NewPlayer(player2Type, player2Letter,
			func() (i int) {
				i = result.getPCMove(player2Letter)
				return i
			},
		)
	case ttgplayer.PlayerPerson:
		result.player2 = ttgplayer.NewPlayer(player2Type, ttgletter.LetterO, result.getPlayerMove)
	}

	return result
}

func (t *TTG) isWinner(board *ttgboard.Board, chainLen int, player ttgletter.Letter) bool {
	b := ttgcommon.GetWinBoard(board.Width(), board.Height(), chainLen)

	for _, i := range b {
		line := 0

		for _, c := range i {
			if board.GetIndexState(c) == player {
				line++
			}
		}

		if line == chainLen {
			return true
		}
	}

	return false
}

func (t *TTG) isBoardFull() bool {
	for i := 0; i < t.width*t.height; i++ {
		if t.board.IsIndexFree(i) {
			return false
		}
	}

	return true
}

// Run runs the game
func (t *TTG) Run() {
	for {
		fmt.Println(t.board)
		i := t.player1.Move()
		t.board.SetIndexState(i, t.player1.Letter())

		if t.isWinner(t.board, t.chainLen, t.player1.Letter()) {
			ttgcommon.Clear()
			fmt.Println(t.board)
			fmt.Println(t.player1.Name() + " won")
			t.pressAnyKeyPrompt()

			break
		} else if t.isBoardFull() {
			ttgcommon.Clear()
			fmt.Println(t.board)
			fmt.Println("DRAW")
			t.pressAnyKeyPrompt()

			break
		}

		fmt.Println(t.board)
		i = t.player2.Move()
		t.board.SetIndexState(i, t.player2.Letter())

		if t.isWinner(t.board, t.chainLen, t.player2.Letter()) {
			ttgcommon.Clear()
			fmt.Println(t.board)
			fmt.Println(t.player2.Name() + " won")
			t.pressAnyKeyPrompt()

			break
		} else if t.isBoardFull() {
			ttgcommon.Clear()
			fmt.Println(t.board)
			fmt.Println("DRAW")
			t.pressAnyKeyPrompt()

			break
		}
	}
}
