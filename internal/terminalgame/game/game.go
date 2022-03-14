package game

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gucio321/tic-tac-go/internal/terminalgame/utils"
	"github.com/gucio321/tic-tac-go/pkg/core/board/letter"
	"github.com/gucio321/tic-tac-go/pkg/core/players/player"
	"github.com/gucio321/tic-tac-go/pkg/game"
)

// TTG represents TicTacToe game.
type TTG struct {
	reader *bufio.Reader
	*game.Game
}

// NewTTG creates a ne TTG.
func NewTTG(w, h, chainLen byte, player1Type, player2Type player.Type) *TTG {
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
	endGame := make(chan bool, 1)
	t.Game.Result(func(l letter.Letter) {
		// handle game end
		switch l {
		case letter.LetterNone:
			fmt.Println("DRAW")
		default:
			fmt.Println(t.CurrentPlayer().Name() + " won")
		}

		t.pressAnyKeyPrompt()
		endGame <- true
	})

	t.Game.UserAction(t.getPlayerMove)

	t.Game.Run()

	<-endGame
}
