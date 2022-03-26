package pcplayer

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gucio321/tic-tac-go/pkg/core/board"
	"github.com/gucio321/tic-tac-go/pkg/core/board/letter"
)

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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetPCMove(tt.args.gameBoard, tt.args.pcLetter)
			assert.Truef(t, Contains(tt.want, got), "GetPCMove() returned unexpected result: expected %v, got %v", tt.want, got)
		})
	}
}

func Test_canWin(t *testing.T) {
	type args struct {
		baseBoard *board.Board
		player    letter.Letter
	}
	tests := []struct {
		name        string
		args        args
		wantCanWin  bool
		wantResults []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCanWin, gotResults := canWin(tt.args.baseBoard, tt.args.player)
			if gotCanWin != tt.wantCanWin {
				t.Errorf("canWin() gotCanWin = %v, want %v", gotCanWin, tt.wantCanWin)
			}
			if !reflect.DeepEqual(gotResults, tt.wantResults) {
				t.Errorf("canWin() gotResults = %v, want %v", gotResults, tt.wantResults)
			}
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := canWinTwoMoves(tt.args.gameBoard, tt.args.player); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("canWinTwoMoves() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_getRandomNumber(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRandomNumber(tt.args.numbers); got != tt.want {
				t.Errorf("getRandomNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
