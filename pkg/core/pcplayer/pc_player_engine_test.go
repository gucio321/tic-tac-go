package pcplayer

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gucio321/tic-tac-go/pkg/core/board"
	"github.com/gucio321/tic-tac-go/pkg/core/board/letter"
)

//nolint:funlen // this is test func, and it's ok
func TestGetPCMove(t *testing.T) {
	type args struct {
		gameBoard *board.Board
		pcLetter  letter.Letter
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		// mechanic (on standard board)
		{
			name: "get corners",
			args: args{
				gameBoard: board.Create(3, 3, 3).
					SetIndexState(2, letter.LetterX),
				pcLetter: letter.LetterX,
			},
			want: []int{0, 2, 6, 8},
		},
		{
			name: "get center",
			args: args{
				gameBoard: board.Create(3, 3, 3).
					SetIndexState(0, letter.LetterX).
					SetIndexState(1, letter.LetterO).
					SetIndexState(2, letter.LetterX).
					SetIndexState(6, letter.LetterO).
					SetIndexState(7, letter.LetterX).
					SetIndexState(8, letter.LetterO),
			},
			want: []int{4},
		},
		{
			name: "get sides",
			args: args{
				gameBoard: board.Create(3, 3, 3).
					SetIndexState(0, letter.LetterO).
					SetIndexState(2, letter.LetterX).
					SetIndexState(4, letter.LetterX).
					SetIndexState(6, letter.LetterX).
					SetIndexState(8, letter.LetterO),
			},
			want: []int{1, 3, 5, 7},
		},
		{
			name: "get opposite corner (PCs)",
			args: args{
				gameBoard: board.Create(3, 3, 3).
					SetIndexState(2, letter.LetterX),
				pcLetter: letter.LetterX,
			},
			want: []int{6},
		},
		{
			name: "get opposite corner (opponent)",
			args: args{
				gameBoard: board.Create(3, 3, 3).
					SetIndexState(8, letter.LetterO),
				pcLetter: letter.LetterX,
			},
			want: []int{0},
		},
		{
			name: "can win two moves (PC) (5x5 board; chain len 4)",
			args: args{
				gameBoard: board.Create(5, 5, 4).
					SetIndexState(11, letter.LetterX).
					SetIndexState(12, letter.LetterX),
				pcLetter: letter.LetterX,
			},
			want: []int{13},
		},
		{
			name: "can win two moves (opponent) (5x5 board; chain len 4)",
			args: args{
				gameBoard: board.Create(5, 5, 4).
					SetIndexState(5, letter.LetterO).
					SetIndexState(10, letter.LetterO),
				pcLetter: letter.LetterX,
			},
			want: []int{15},
		},

		// behaviors
		{
			name: "Empty Board 3x3, first move",
			args: args{
				gameBoard: board.Create(3, 3, 3),
				pcLetter:  letter.LetterX,
			},
			want: []int{0, 2, 6, 8},
		},
		{
			name: "3x3 Board, pc can win",
			args: args{
				gameBoard: board.Create(3, 3, 3).
					SetIndexState(0, letter.LetterX).
					SetIndexState(8, letter.LetterX),
				pcLetter: letter.LetterX,
			},
			want: []int{4},
		},
		{
			name: "3x3 Board, pc and opponent can win",
			args: args{
				gameBoard: board.Create(3, 3, 3).
					SetIndexState(0, letter.LetterX).
					SetIndexState(2, letter.LetterX).
					SetIndexState(3, letter.LetterO).
					SetIndexState(5, letter.LetterO),
				pcLetter: letter.LetterX,
			},
			want: []int{1},
		},
		{
			name: "3x3 Board, opponent can win",
			args: args{
				gameBoard: board.Create(3, 3, 3).
					SetIndexState(3, letter.LetterO).
					SetIndexState(5, letter.LetterO),
				pcLetter: letter.LetterX,
			},
			want: []int{4},
		},

		{
			name: "Empty Board 4x4 (chain len 4), first move",
			args: args{
				gameBoard: board.Create(4, 4, 4),
				pcLetter:  letter.LetterX,
			},
			want: []int{0, 3, 12, 15},
		},
		{
			name: "4x4 Board (chain len 4), pc can win",
			args: args{
				gameBoard: board.Create(4, 4, 4).
					SetIndexState(3, letter.LetterX).
					SetIndexState(6, letter.LetterX).
					SetIndexState(9, letter.LetterX),
				pcLetter: letter.LetterX,
			},
			want: []int{12},
		},
		{
			name: "4x4 Board (chain len 4), pc and opponent can win",
			args: args{
				gameBoard: board.Create(4, 4, 4).
					SetIndexState(0, letter.LetterX).
					SetIndexState(1, letter.LetterX).
					SetIndexState(2, letter.LetterX).
					SetIndexState(5, letter.LetterO).
					SetIndexState(6, letter.LetterO).
					SetIndexState(7, letter.LetterO),
				pcLetter: letter.LetterX,
			},
			want: []int{3},
		},
		{
			name: "4x4 (chain len 4) Board, opponent can win",
			args: args{
				gameBoard: board.Create(4, 4, 4).
					SetIndexState(1, letter.LetterO).
					SetIndexState(5, letter.LetterO).
					SetIndexState(9, letter.LetterO),
				pcLetter: letter.LetterX,
			},
			want: []int{13},
		},

		{
			name: "Empty Board 4x4 (chain len 3), first move",
			args: args{
				gameBoard: board.Create(4, 4, 3),
				pcLetter:  letter.LetterX,
			},
			want: []int{0, 3, 12, 15},
		},
		{
			name: "4x4 Board (chain len 3), pc can win",
			args: args{
				gameBoard: board.Create(4, 4, 3).
					SetIndexState(10, letter.LetterX).
					SetIndexState(14, letter.LetterX),
				pcLetter: letter.LetterX,
			},
			want: []int{6},
		},
		{
			name: "4x4 Board (chain len 3), pc and opponent can win",
			args: args{
				gameBoard: board.Create(4, 4, 3).
					SetIndexState(6, letter.LetterX).
					SetIndexState(9, letter.LetterX).
					SetIndexState(2, letter.LetterO).
					SetIndexState(7, letter.LetterO).
					SetIndexState(11, letter.LetterO),
				pcLetter: letter.LetterX,
			},
			want: []int{3, 12},
		},
		{
			name: "4x4 (chain len 3) Board, opponent can win",
			args: args{
				gameBoard: board.Create(4, 4, 3).
					SetIndexState(1, letter.LetterO).
					SetIndexState(11, letter.LetterO),
				pcLetter: letter.LetterX,
			},
			want: []int{6},
		},
		{
			name: "inside board",
			args: args{
				gameBoard: board.Create(4, 4, 4).
					SetIndexState(0, letter.LetterX).
					SetIndexState(1, letter.LetterO).
					SetIndexState(2, letter.LetterO).
					SetIndexState(3, letter.LetterX).
					SetIndexState(4, letter.LetterX).
					SetIndexState(7, letter.LetterO).
					SetIndexState(8, letter.LetterX).
					SetIndexState(11, letter.LetterO).
					SetIndexState(12, letter.LetterO).
					SetIndexState(13, letter.LetterX).
					SetIndexState(14, letter.LetterX).
					SetIndexState(15, letter.LetterO),
				pcLetter: letter.LetterX,
			},
			want: []int{5, 6, 9, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewPCPlayer(tt.args.gameBoard, tt.args.pcLetter).GetMove()
			assert.Truef(t, contains(tt.want, got), "GetPCMove() returned unexpected result: expected %v, got %v", tt.want, got)
		})
	}
}

func TestGetPCMove_FullBoard(t *testing.T) {
	gameBoard := board.Create(4, 4, 4).
		SetIndexState(0, letter.LetterX).
		SetIndexState(1, letter.LetterO).
		SetIndexState(2, letter.LetterX).
		SetIndexState(3, letter.LetterO).
		SetIndexState(4, letter.LetterX).
		SetIndexState(5, letter.LetterO).
		SetIndexState(6, letter.LetterX).
		SetIndexState(7, letter.LetterO).
		SetIndexState(8, letter.LetterX).
		SetIndexState(9, letter.LetterO).
		SetIndexState(10, letter.LetterX).
		SetIndexState(11, letter.LetterO).
		SetIndexState(12, letter.LetterX).
		SetIndexState(13, letter.LetterO).
		SetIndexState(14, letter.LetterX).
		SetIndexState(15, letter.LetterO)

	assert.Panics(t, func() { NewPCPlayer(gameBoard, letter.LetterX).GetMove() }, "GetPCMove on full board didn't panicked")
	assert.Panics(t, func() { NewPCPlayer(gameBoard, letter.LetterO).GetMove() }, "GetPCMove on full board didn't panicked")
}

//nolint:funlen // tests function; it is ok
func Test_canWin(t *testing.T) {
	type args struct {
		baseBoard *board.Board
	}

	tests := []struct {
		name         string
		args         args
		wantCanWinX  bool
		wantResultsX []int
		wantCanWinO  bool
		wantResultsO []int
	}{
		{
			name: "Empty board 3x3",
			args: args{
				baseBoard: board.Create(3, 3, 3),
			},
			wantCanWinX:  false,
			wantResultsX: []int{},
			wantCanWinO:  false,
			wantResultsO: []int{},
		},
		{
			name: "3x3 board: X can win",
			args: args{
				baseBoard: board.Create(3, 3, 3).
					SetIndexState(0, letter.LetterX).
					SetIndexState(8, letter.LetterX),
			},
			wantCanWinX:  true,
			wantResultsX: []int{4},
			wantCanWinO:  false,
			wantResultsO: []int{},
		},
		{
			name: "3x3 board: X can win (two ways)",
			args: args{
				baseBoard: board.Create(3, 3, 3).
					SetIndexState(0, letter.LetterX).
					SetIndexState(2, letter.LetterX).
					SetIndexState(6, letter.LetterX).
					SetIndexState(4, letter.LetterO),
			},
			wantCanWinX:  true,
			wantResultsX: []int{1, 3},
			wantCanWinO:  false,
			wantResultsO: []int{},
		},
		{
			name: "3x3 board: X can win (3 ways)",
			args: args{
				baseBoard: board.Create(3, 3, 3).
					SetIndexState(2, letter.LetterX).
					SetIndexState(6, letter.LetterX).
					SetIndexState(8, letter.LetterX),
			},
			wantCanWinX:  true,
			wantResultsX: []int{4, 5, 7},
			wantCanWinO:  false,
			wantResultsO: []int{},
		},
		{
			name: "3x3 board: X and O can win",
			args: args{
				baseBoard: board.Create(3, 3, 3).
					SetIndexState(0, letter.LetterX).
					SetIndexState(6, letter.LetterX).
					SetIndexState(2, letter.LetterO).
					SetIndexState(8, letter.LetterO),
			},
			wantCanWinX:  true,
			wantResultsX: []int{3},
			wantCanWinO:  true,
			wantResultsO: []int{5},
		},

		{
			name: "Empty board 4x4",
			args: args{
				baseBoard: board.Create(4, 4, 4),
			},
			wantCanWinX:  false,
			wantResultsX: []int{},
			wantCanWinO:  false,
			wantResultsO: []int{},
		},
		{
			name: "4x4 board: X can win",
			args: args{
				baseBoard: board.Create(4, 4, 4).
					SetIndexState(5, letter.LetterX).
					SetIndexState(10, letter.LetterX).
					SetIndexState(15, letter.LetterX),
			},
			wantCanWinX:  true,
			wantResultsX: []int{0},
			wantCanWinO:  false,
			wantResultsO: []int{},
		},
		{
			name: "4x4 board: X can win (two ways)",
			args: args{
				baseBoard: board.Create(4, 4, 4).
					SetIndexState(0, letter.LetterX).
					SetIndexState(4, letter.LetterX).
					SetIndexState(8, letter.LetterX).
					SetIndexState(1, letter.LetterX).
					SetIndexState(9, letter.LetterX).
					SetIndexState(13, letter.LetterX),
			},
			wantCanWinX:  true,
			wantResultsX: []int{5, 12},
			wantCanWinO:  false,
			wantResultsO: []int{},
		},
		{
			name: "4x4 board: X can win (3 ways)",
			args: args{
				baseBoard: board.Create(4, 4, 4).
					SetIndexState(0, letter.LetterX).
					SetIndexState(1, letter.LetterX).
					SetIndexState(4, letter.LetterX).
					SetIndexState(5, letter.LetterX).
					SetIndexState(7, letter.LetterX).
					SetIndexState(9, letter.LetterX).
					SetIndexState(15, letter.LetterX),
			},
			wantCanWinX:  true,
			wantResultsX: []int{6, 10, 13},
			wantCanWinO:  false,
			wantResultsO: []int{},
		},
		{
			name: "4x4 board: X and O can win",
			args: args{
				baseBoard: board.Create(4, 4, 4).
					SetIndexState(0, letter.LetterX).
					SetIndexState(4, letter.LetterX).
					SetIndexState(8, letter.LetterX).
					SetIndexState(3, letter.LetterO).
					SetIndexState(6, letter.LetterO).
					SetIndexState(9, letter.LetterO),
			},
			wantCanWinX:  true,
			wantResultsX: []int{12},
			wantCanWinO:  true,
			wantResultsO: []int{12},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCanWinX, gotResultsX := NewPCPlayer(tt.args.baseBoard, letter.LetterX).canWin(tt.args.baseBoard, letter.LetterX)
			gotCanWinO, gotResultsO := NewPCPlayer(tt.args.baseBoard, letter.LetterO).canWin(tt.args.baseBoard, letter.LetterO)
			assert.Equal(t, gotCanWinX, tt.wantCanWinX, "canWin returned unexpected value")
			assert.Equal(t, gotResultsX, tt.wantResultsX, "canWin returned unexpected value (list of winning combos)")
			assert.Equal(t, gotCanWinO, tt.wantCanWinO, "canWin returned unexpected value")
			assert.Equal(t, gotResultsO, tt.wantResultsO, "canWin returned unexpected value (list of winning combos)")
		})
	}
}

func Test_canWinTwoMoves(t *testing.T) {
	type args struct {
		gameBoard *board.Board
		player    letter.Letter
	}

	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{
			name: "5x5 board, chain len 4",
			args: args{
				gameBoard: board.Create(5, 5, 4).
					SetIndexState(6, letter.LetterX).
					SetIndexState(7, letter.LetterX),
				player: letter.LetterX,
			},
			wantResult: []int{8},
		},
		{
			name: "5x5 board; chain len 4 (v2)",
			args: args{
				gameBoard: board.Create(5, 5, 4).
					SetIndexState(6, letter.LetterX).
					SetIndexState(12, letter.LetterX),
				player: letter.LetterX,
			},
			wantResult: []int{18},
		},
		{
			name: "5x5 board; chain len 3",
			args: args{
				gameBoard: board.Create(5, 5, 3).
					SetIndexState(12, letter.LetterX),
				player: letter.LetterX,
			},
			wantResult: []int{11, 13, 7, 17, 6, 18, 8, 16},
		},
		{
			name: "too small chain length",
			args: args{
				gameBoard: board.Create(3, 3, 2),
				player:    letter.LetterX,
			},
			wantResult: []int{},
		},

		{
			name: "unable to win in two moves (nil expected)",
			args: args{
				gameBoard: board.Create(5, 5, 4).
					SetIndexState(20, letter.LetterO).
					SetIndexState(21, letter.LetterX).
					SetIndexState(22, letter.LetterX),
				player: letter.LetterX,
			},
			wantResult: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult := NewPCPlayer(tt.args.gameBoard, tt.args.player).canWinTwoMoves(tt.args.gameBoard, tt.args.player)
			assert.Equal(t, tt.wantResult, gotResult, "canWinTwoMoves() = %v, want %v", gotResult, tt.wantResult)
		})
	}
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}

func TestPCPlayer_String(t *testing.T) {
	type fields struct {
		b        *board.Board
		pcLetter letter.Letter
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "for player X",
			fields: fields{
				b:        nil,
				pcLetter: letter.LetterX,
			},
			want: "PC X",
		},
		{
			name: "for player O",
			fields: fields{
				b:        nil,
				pcLetter: letter.LetterO,
			},
			want: "PC O",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PCPlayer{
				b:        tt.fields.b,
				pcLetter: tt.fields.pcLetter,
			}

			assert.Equalf(t, tt.want, p.String(), "String()")
		})
	}
}
