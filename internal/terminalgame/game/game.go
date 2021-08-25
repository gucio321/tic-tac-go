package game

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gucio321/tic-tac-go/internal/terminalgame/utils"
	"github.com/gucio321/tic-tac-go/pkg/ttgcore/ttggame"
	"github.com/gucio321/tic-tac-go/pkg/ttgcore/ttgletter"
	"github.com/gucio321/tic-tac-go/pkg/ttgcore/ttgplayers/ttgplayer"
)

// TTG represents TicTacToe game.
type TTG struct {
	reader *bufio.Reader
	*ttggame.Game
}

// NewTTG creates a ne TTG.
func NewTTG(w, h, chainLen byte, player1Type, player2Type ttgplayer.PlayerType) *TTG {
	result := &TTG{
		reader: bufio.NewReader(os.Stdin),
		Game:   ttggame.Create(player1Type, player2Type),
	}

	result.SetBoardSize(int(w), int(h), int(chainLen)).OnContinue(func() {
		utils.Clear()
		fmt.Println(result.Board())
	})

	return result
}

// Run runs the game.
func (t *TTG) Run() {
	go t.Game.Run()

	for {
		// handle user move
		if t.IsUserActionRequired() {
			t.TakeUserAction(t.getPlayerMove())
		}

		// handle game end
		if reached, result := t.Result(); reached {
			switch result {
			case ttgletter.LetterNone:
				fmt.Println("DRAW")
			default:
				fmt.Println(t.CurrentPlayer().Name() + " won")
			}

			t.pressAnyKeyPrompt()

			break
		}
	}
}
