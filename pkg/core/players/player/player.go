package player

import (
	"fmt"
)

type Player interface {
	GetMove() int
	fmt.Stringer
}
