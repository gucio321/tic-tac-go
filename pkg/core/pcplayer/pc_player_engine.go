// Package pcplayer provides methods for simple-AI logic
// used in Tic-Tac-Go for calculating PC-player's move.
package pcplayer

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"

	"github.com/gucio321/tic-tac-go/internal/logger"

	"github.com/gucio321/tic-tac-go/pkg/core/board"
	"github.com/gucio321/tic-tac-go/pkg/core/board/letter"
	"github.com/gucio321/tic-tac-go/pkg/core/players/player"
)

type AlgType byte

const (
	AlgOriginal AlgType = iota
	AlgMinMax
)

var _ player.Player = &PCPlayer{}

// PCPlayer is a simple-AI logic used in Tic-Tac-Go for calculating PC-player's move.
type PCPlayer struct {
	b        *board.Board
	pcLetter letter.Letter
	algType  AlgType
}

// NewPCPlayer creates new PCPlayer instance.
func NewPCPlayer(b *board.Board, pcLetter letter.Letter, algType AlgType) *PCPlayer {
	return &PCPlayer{
		b:        b,
		pcLetter: pcLetter,
		algType:  algType,
	}
}

func (p *PCPlayer) String() string {
	return "PC " + p.pcLetter.String()
}

// GetMove calculates move for PC player on given board.
// Steps:
// - try to win
// - stop opponent from winning
// - try to win in 2-moves perspective
// - stop opponent from winning in 2-moves perspective
// - for current board and for each smaller board (w - 2 h - 2)
//   - take opposite (to any corner taken by pc) corner
//   - take opponent's opposite corner
//   - take center
//   - take random side
func (p PCPlayer) GetMove() (i int) {
	logger.Infof("Calculating Move for PC Player")

	switch p.algType {
	case AlgOriginal:
		return p.getPCMove(p.b)
	case AlgMinMax:
		return p.minMax(p.b, 10)
	default:
		panic(fmt.Sprintf("Unknown algorithm type: %v", p.algType))
	}
}

// THis is a min-max algorithm implementation.
// This algorithm predicts all possible solutions and chooses the best one.
// After writting this I found out the followint:
// 1. This is really ineffective: It is playable on 3x3 board, but on 4x4 it
// freezes my 12-core, 16GB RAM machine. (I will try to add MaxDepth (after reaching this it will just randomize the move)
// and maybe I'll try to optimize it so that it doesn't call recursively if not needed (solution worse than current worst))
// 2. This is a bit theoritical conclusion but: In theory of 3x3 tic-tac-toe game, the best 2nd move (if 1st player took corner)
// should be taking the center. However after looking at algorithnm's behaviour it turns out
// that taking the center will not lead to the fastest winning opportunity. Conclusion: the algorithm should be
// improved to consider "unblockable wins" and "draws"
//
// UPDATE 1: I've added maxDepth parameter. Now user can specify how many moves ahead the algorithm should predict.
// UPDATE 2: Number of calls to mm is boardArea^maxDepth. So DONT EVEN TRY IT FOR 4x4 board with maxDepth ~10 (1099511627776 callss)!!!
func (p *PCPlayer) minMax(gameBoard *board.Board, maxDepth int) (i int) {
	//progress := int(math.Pow(float64(gameBoard.Width()*gameBoard.Height()), float64(maxDepth)))
	m := &sync.Mutex{}
	cw := new(bool)
	move := new(int)
	bestDepth := new(int) // this one is for internal state of mm - will hold lowest depth for which canWin was true
	p.mm(gameBoard, p.pcLetter, 0, maxDepth, m, bestDepth, move, cw)
	fmt.Println("Best Depth", *bestDepth)
	// now if can't get best move get random from possible
	if !*cw {
		logger.Infof("Randomize move")
		for i := 0; i < gameBoard.Width()*gameBoard.Height(); i++ {
			if !gameBoard.IsIndexFree(i) {
				continue
			}
			*move = i
			break
		}
	}
	return *move
}

func (a *PCPlayer) mm(
	gameBoard *board.Board, l letter.Letter, currentDepth int,
	maxDepth int, mutex *sync.Mutex, best *int, move *int, couldWin *bool,
) {
	//logger.Debugf("mm: call for %s (depth: %d)\n%s", l, currentDepth, gameBoard)
	mutex.Lock()
	if *best <= currentDepth && *couldWin {
		//logger.Debugf("mm: depth %d: best better than current (%d > %d)", currentDepth, *best, currentDepth)
		mutex.Unlock()

		return
	}

	mutex.Unlock()

	wg := sync.WaitGroup{}
	for i := 0; i < gameBoard.Width()*gameBoard.Height(); i++ {
		if !gameBoard.IsIndexFree(i) {
			continue
		}

		cp := gameBoard.Copy()
		cp.SetIndexState(i, l)
		if winner, u := cp.IsWinner(l); winner {
			logger.Debugf("Can win at %d (combo %v) Depth %d", i, u, currentDepth)
			mutex.Lock()
			if !*couldWin || (*couldWin && *best > currentDepth) {
				//logger.Warnf("Found move %d at depth %d (Potential winner is %s)", m, currentDepth, l)
				//logger.Debugf("mm depth %d: updated best move to %d (on depth %d)", currentDepth, cpMove, cpDepth)
				*best = currentDepth
				*move = i
				*couldWin = true
			}
			mutex.Unlock()
		}

		//logger.Debugf("re-running for opposite letter")
		wg.Add(1)
		go func() {
			a.mm(cp, l.Opposite(), currentDepth+1, maxDepth, mutex, best, move, couldWin)
			wg.Done()
		}()
	}

	wg.Wait()
}

//nolint:gocyclo,funlen // https://github.com/gucio321/tic-tac-go/issues/154
func (p *PCPlayer) getPCMove(gameBoard *board.Board) (i int) {
	logger.Debugf("Game board is\n %s", gameBoard)
	logger.Debugf("PC letter is %s", p.pcLetter)

	playerLetter := p.pcLetter.Opposite()

	logger.Debugf("opponent's letter is %s", playerLetter)

	// attack: try to win now
	logger.Debug("Performing attack")

	if ok, indexes := p.canWin(gameBoard, p.pcLetter); ok {
		logger.Debugf("Can win now, indexes: %v", indexes)
		options := p.getAvailableOptions(gameBoard, indexes)

		if len(options) > 0 {
			return p.getRandomNumber(options)
		}
	}

	// defense: check, if user can win
	logger.Debug("Performing defense")

	if ok, indexes := p.canWin(gameBoard, playerLetter); ok {
		logger.Debugf("Player can win now, indexes: %v", indexes)
		options := p.getAvailableOptions(gameBoard, indexes)

		if len(options) > 0 {
			return p.getRandomNumber(options)
		}
	}

	logger.Debug("Performing attack (2 moves)")

	options := p.getAvailableOptions(gameBoard, p.canWinTwoMoves(gameBoard, p.pcLetter))
	if len(options) > 0 {
		logger.Debugf("Player can win in two moves, indexes: %v", options)

		return p.getRandomNumber(options)
	}

	logger.Debug("Performing defense (2 moves)")

	options = p.getAvailableOptions(gameBoard, p.canWinTwoMoves(gameBoard, playerLetter))
	if len(options) > 0 {
		logger.Debugf("Player can win in two moves, indexes: %v", options)

		return p.getRandomNumber(options)
	}

	corners := gameBoard.GetCorners()
	pcOppositeCorners := make([]int, 0)
	playerOppositeCorners := make([]int, 0)

	for _, i := range corners {
		if gameBoard.IsIndexFree(i) {
			options = append(options, i)

			continue
		}

		o := gameBoard.GetOppositeCorner(i)

		if !gameBoard.IsIndexFree(o) {
			continue
		}

		switch s := gameBoard.GetIndexState(i); s {
		case p.pcLetter:
			pcOppositeCorners = append(pcOppositeCorners, o)
		case playerLetter:
			playerOppositeCorners = append(playerOppositeCorners, o)
		}
	}

	if len(pcOppositeCorners) != 0 {
		return p.getRandomNumber(pcOppositeCorners)
	}

	if len(playerOppositeCorners) != 0 {
		return p.getRandomNumber(playerOppositeCorners)
	}

	if len(options) > 0 {
		return p.getRandomNumber(options)
	}

	// try to get center
	for _, i := range gameBoard.GetCenter() {
		if gameBoard.IsIndexFree(i) {
			options = append(options, i)
		}
	}

	if len(options) > 0 {
		return p.getRandomNumber(options)
	}

	for _, i := range gameBoard.GetSides() {
		if gameBoard.IsIndexFree(i) {
			options = append(options, i)
		}
	}

	if len(options) > 0 {
		return p.getRandomNumber(options)
	}

	const smallerBoard = 2
	if newW, newH := gameBoard.Width()-smallerBoard, gameBoard.Height()-smallerBoard; newW > 0 && newH > 0 {
		return gameBoard.ConvertIndex(newW, newH, p.getPCMove(gameBoard.Cut(newW, newH)))
	}

	panic("Tic-Tac-Go: pcplayer.GetPCMove(...): cannot determinate pc move - board is full")
}

func (p *PCPlayer) canWin(baseBoard *board.Board, playerLetter letter.Letter) (canWin bool, results []int) {
	results = make([]int, 0)

	for i := 0; i < baseBoard.Width()*baseBoard.Height(); i++ {
		if !baseBoard.IsIndexFree(i) {
			continue
		}

		fictionBoard := baseBoard.Copy()

		fictionBoard.SetIndexState(i, playerLetter)

		if ok, _ := fictionBoard.IsWinner(playerLetter); ok {
			results = append(results, i)
		}
	}

	return len(results) > 0, results
}

/*
This method should find situations like that:
chain length = 4
+---+---+---+---+---+
|   |   | o |   |   |
+---+---+---+---+---+
|   |   |   |   | o |
+---+---+---+---+---+
|   | x | x |   |   |
+---+---+---+---+---+
|   |   |   |   |   |
+---+---+---+---+---+
|   |   |   | o |   |
+---+---+---+---+---+
let's look at the board above: when we'll make our move at right side of X-chain (14)
the O-player will not be able to keep X from winning.
+---+---+---+---+---+
|   |   | o |   |   |
+---+---+---+---+---+
|   |   |   |   | o |
+---+---+---+---+---+
|   | x | x | X |   |
+---+---+---+---+---+
|   |   |   |   |   |
+---+---+---+---+---+
|   |   |   | o |   |
+---+---+---+---+---+
O-player lost.
*/
func (p *PCPlayer) canWinTwoMoves(gameBoard *board.Board, playerLetter letter.Letter) (result []int) {
	result = make([]int, 0)

	//nolint:gomnd // look a scheme above - in the second one, the chain is by 2 less than max
	minimalChainLen := gameBoard.ChainLength() - 2
	if minimalChainLen <= 0 {
		return
	}

	potentiallyAvailableChains := gameBoard.GetWinBoard(gameBoard.ChainLength() + 1)

searching:
	for _, potentialPlace := range potentiallyAvailableChains {
		if !gameBoard.IsIndexFree(potentialPlace[0]) || !gameBoard.IsIndexFree(potentialPlace[len(potentialPlace)-1]) {
			continue
		}

		var gaps []int

		for i := 1; i < len(potentialPlace)-1; i++ {
			switch gameBoard.GetIndexState(potentialPlace[i]) {
			case letter.LetterNone:
				gaps = append(gaps, potentialPlace[i])
			case playerLetter.Opposite(): // operation already blocked
				continue searching
			}
		}

		if len(gaps) == 1 {
			result = append(result, gaps...)
		}
	}

	return result
}

func (p *PCPlayer) getRandomNumber(numbers []int) int {
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(int64(len(numbers))))
	if err != nil {
		panic(fmt.Sprintf("Reading random number: %v", err))
	}

	return numbers[randomNumber.Int64()]
}

func (p *PCPlayer) getAvailableOptions(gameBoard *board.Board, candidates []int) (available []int) {
	available = make([]int, 0)

	for _, i := range candidates {
		if gameBoard.IsIndexFree(i) {
			available = append(available, i)
		}
	}

	return available
}
