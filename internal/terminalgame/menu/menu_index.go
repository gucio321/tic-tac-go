package menu

import "fmt"

type menuIndex struct {
	number int16
	label  string
	cb     func()
}

func (i *menuIndex) String() string {
	return fmt.Sprintf("%d) %s", i.number, i.label)
}
