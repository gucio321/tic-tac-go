package ttgboard

// PlayerType represents players' types
type PlayerType int

// player types
const (
	PlayerPC PlayerType = iota
	PlayerPerson
)

func (p PlayerType) String() string {
	switch p {
	case PlayerPC:
		return "PC"
	case PlayerPerson:
		return "Player"
	}

	return "?"
}

type player struct {
	name       string
	playerType PlayerType
	letter     Letter
	moveCb     func() (i int)
}

func newPlayer(t PlayerType, letter Letter, cb func() (i int)) *player {
	result := &player{
		playerType: t,
		letter:     letter,
		moveCb:     cb,
		name:       t.String() + " " + letter.String(),
	}

	return result
}
