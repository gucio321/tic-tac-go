package game

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gucio321/tic-tac-go/internal/terminalgame/utils"
	"github.com/gucio321/tic-tac-go/pkg/core/ttgletter"
	"github.com/gucio321/tic-tac-go/pkg/core/ttgplayers/ttgplayer"
	"github.com/gucio321/tic-tac-go/pkg/game"
)

// TTG represents TicTacToe game.
type TTG struct {
	reader *bufio.Reader
	*game.Game
}

// NewTTG creates a ne TTG.
func NewTTG(w, h, chainLen byte, player1Type, player2Type ttgplayer.PlayerType) *TTG {
	result := &TTG{
		reader: bufio.NewReader(os.Stdin),
		Game:   game.Create(player1Type, player2Type),
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
