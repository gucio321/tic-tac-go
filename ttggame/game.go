package ttggame

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gucio321/tic-tac-go/ttgcommon"
	"github.com/gucio321/tic-tac-go/ttggame/ttgboard"
	"github.com/gucio321/tic-tac-go/ttggame/ttgletter"
	"github.com/gucio321/tic-tac-go/ttggame/ttgpcplayer"
	"github.com/gucio321/tic-tac-go/ttggame/ttgplayer"
)

// TTG represents TicTacToe game
type TTG struct {
	board   *ttgboard.Board
	reader  *bufio.Reader
	players *ttgplayer.Players
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

	var cb1, cb2 func()

	switch player1Type {
	case ttgplayer.PlayerPC:
		cb1 = func() {
			result.board.SetIndexState(
				ttgpcplayer.GetPCMove(result.board, player1Letter),
				player1Letter,
			)
		}
	case ttgplayer.PlayerPerson:
		cb1 = func() {
			result.board.SetIndexState(
				result.getPlayerMove(),
				player1Letter,
			)
		}
	}

	switch player2Type {
	case ttgplayer.PlayerPC:
		cb2 = func() {
			result.board.SetIndexState(
				ttgpcplayer.GetPCMove(result.board, player2Letter),
				player2Letter,
			)
		}
	case ttgplayer.PlayerPerson:
		cb2 = func() {
			result.board.SetIndexState(
				result.getPlayerMove(),
				player2Letter,
			)
		}
	}

	result.players = ttgplayer.Create(player1Type, cb1, player2Type, cb2)

	return result
}

// Run runs the game
func (t *TTG) Run() {
	for {
		fmt.Println(t.board)
		t.players.Current().Move()

		if t.board.IsWinner(t.chainLen, t.players.Current().Letter()) {
			ttgcommon.Clear()
			fmt.Println(t.board)
			fmt.Println(t.players.Current().Name() + " won")
			t.pressAnyKeyPrompt()

			break
		} else if t.board.IsBoardFull() {
			ttgcommon.Clear()
			fmt.Println(t.board)
			fmt.Println("DRAW")
			t.pressAnyKeyPrompt()

			break
		}

		// switch to next player
		t.players.Next()
	}
}
