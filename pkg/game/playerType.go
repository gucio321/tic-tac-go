package game

// PlayerType represents type of player (human or Computer).
type PlayerType byte

// player types.
const (
	PlayerTypeHuman PlayerType = iota
	PlayerTypePCOriginal
	PlayerTypePCMinMax
)
