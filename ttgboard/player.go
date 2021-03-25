package ttgboard

type player struct {
	name       string
	playerType PlayerType
	letter     IdxState
	moveCb     func() (x, y int)
}

func newPlayer(t PlayerType, letter IdxState, cb func() (x, y int)) *player {
	result := &player{
		playerType: t,
		letter:     letter,
		moveCb:     cb,
		name:       t.String() + " " + letter.String(),
	}

	return result
}
