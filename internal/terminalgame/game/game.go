package game

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gucio321/go-clear"
	"github.com/gucio321/terminalmenu/pkg/menuutils"

	"github.com/gucio321/tic-tac-go/pkg/core/players/player"
	"github.com/gucio321/tic-tac-go/pkg/game"
)

// TTG represents TicTacToe game.
type TTG struct {
	reader *bufio.Reader
	*game.Game
}

// NewTTG creates a new TTG.
func NewTTG(w, h, chainLen byte, player1Type, player2Type player.Type) *TTG {
	result := &TTG{
		reader: bufio.NewReader(os.Stdin),
		Game:   game.Create(player1Type, player2Type),
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

	t.Game.Result(func(p *player.Player) {
		// handle game end
		switch p {
		case nil:
			fmt.Println("DRAW")
		default:
			fmt.Println(p.Name() + " won")
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
