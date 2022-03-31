package player

import (
	"fmt"
)

// Player is an interface implemented by any further Player implementations.
type Player interface {
	GetMove() int
	fmt.Stringer
}
