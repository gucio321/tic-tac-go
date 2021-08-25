package menu

const (
	header = "Welcome in TicTacGo\n\n"
)

type menuPage struct {
	title   string
	options []*menuIndex
}

func (p *menuPage) String() string {
	page := header
	page += "\t" + p.title + "\n"

	for _, i := range p.options {
		page += "\t\t" + i.String() + "\n"
	}

	return page
}

func (p *menuPage) Exec(idx int16) {
	for _, option := range p.options {
		if option.number == idx {
			option.cb()

			return
		}
	}
}

func (p *menuPage) Max() int16 {
	var max int16
	for _, option := range p.options {
		if option.number > max {
			max = option.number
		}
	}

	return max
}
