package gameimpl

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gucio321/tic-tac-go/pkg/core/board/letter"

	"github.com/gucio321/go-clear"
	"github.com/gucio321/terminalmenu/pkg/menuutils"

	"github.com/gucio321/tic-tac-go/pkg/game"
)

// TTG represents TicTacToe game.
type TTG struct {
	reader *bufio.Reader
	*game.Game
}

// NewTTG creates a new TTG.
func NewTTG(w, h, chainLen byte, playerXType, playerOType game.PlayerType) *TTG {
	result := &TTG{
		reader: bufio.NewReader(os.Stdin),
		Game:   game.Create(playerXType, playerOType),
	}

	result.SetBoardSize(int(w), int(h), int(chainLen)).OnContinue(func() {
		if err := clear.Clear(); err != nil {
			log.Fatal(err)
		}

		fmt.Println(result.Board())
	})

	return result
}

// Run runs the game.
func (t *TTG) Run() {
	endGame := make(chan bool, 1)

	t.Game.Result(func(l letter.Letter, name string) {
		// handle game end
		switch l {
		case letter.LetterNone:
			fmt.Println("DRAW")
		default:
			fmt.Println(name + " won")
		}

		if err := menuutils.PromptEnter("Press ENTER to continue "); err != nil {
			log.Fatal(err)
		}

		endGame <- true
	})

	t.Game.UserAction(t.getPlayerMove)

	t.Game.Run()

	<-endGame
}
