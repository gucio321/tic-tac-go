package ttgmenu

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/AllenDang/giu"

	game "github.com/gucio321/tic-tac-go/ttgboard"
	"github.com/gucio321/tic-tac-go/ttgcommon"
)

// Menu represent's game's menu
type Menu struct {
	state  State
	reader *bufio.Reader
	menus  map[State]*menuIndex
}

// NewMenu creates a new game menu
func (m *Menu) getMenuData(state State) (lines []string, actions []func()) {
	text := map[State][]string{
		MainMenu: {
			"1) start Player VS PC game",
			"2) start Player VS Player game",
			"3) settings",
			"4) Help",
			"0) exit",
		},
		Help: {
			"TicTacToe Version 1",
			"Copytight (C) 2021 by M. Sz.",
			"",
			"To go around main menu use number buttons",
			"In game use 1-9 buttons to select index",
			"+---+---+---+",
			"| 1 | 2 | 3 |",
			"+---+---+---+",
			"| 4 | 5 | 6 |",
			"+---+---+---+",
			"| 7 | 8 | 9 |",
			"+---+---+---+",
			"",
			"Press enter to back to main menu",
		},
		Settings: {
			"0) back to main menu",
		},
	}

	cb := map[State][]func(){
		MainMenu: {
			func() {
				var g *game.TTT

				rand.Seed(time.Now().UnixNano())
				// nolint:gomnd // number of players in game
				r := rand.Intn(2) // nolint:gosec // it is ok

				switch r {
				case 0:
					g = game.NewTTT(ttgcommon.BaseBoardW, ttgcommon.BaseBoardH, game.PlayerPerson, game.PlayerPC)
				case 1:
					g = game.NewTTT(ttgcommon.BaseBoardW, ttgcommon.BaseBoardH, game.PlayerPC, game.PlayerPerson)
				}

				g.Run()
			},
			func() {
				game := game.NewTTT(ttgcommon.BaseBoardW, ttgcommon.BaseBoardH, game.PlayerPerson, game.PlayerPerson)
				game.Run()
			},
			func() {
				m.state = Settings
			},
			func() {
				m.state = Help
			},
			func() { os.Exit(0) },
		},
		Help: {
			func() {
				m.state = MainMenu
			},
		},
		Settings: {
			func() {
				m.state = MainMenu
			},
		},
	}

	return text[state], cb[state]
}

// State represents menu's state
type State int

// menu states
const (
	MainMenu State = iota
	Help
	Settings
)

type menuIndex struct {
	lines       []string
	userActions []func()
	multiAction bool
}

func (m *Menu) newMenuIndex(state State) *menuIndex {
	result := &menuIndex{}

	result.lines, result.userActions = m.getMenuData(state)

	if len(result.userActions) > 1 && result.userActions != nil {
		result.multiAction = true
	} else {
		result.multiAction = false
	}

	return result
}

// NewMenu creates a new menu
func NewMenu() *Menu {
	result := &Menu{
		state:  MainMenu,
		reader: bufio.NewReader(os.Stdin),
	}

	result.menus = map[State]*menuIndex{
		MainMenu: result.newMenuIndex(MainMenu),
		Help:     result.newMenuIndex(Help),
		Settings: result.newMenuIndex(Settings),
	}

	return result
}

func (m *Menu) printMenu() giu.Layout {
	var l giu.Layout
	lines := m.menus[m.state].lines
	if m.menus[m.state].multiAction {
		for n, line := range lines {
			n := n
			l = append(l,
				giu.Button(line).OnClick(func() {
					m.menus[m.state].userActions[n]()
				}),
			)
		}
	} else {
		l = append(l,
			giu.Label(strings.Join(lines, "\n")),
			giu.Button("Back").OnClick(m.menus[m.state].userActions[0]),
		)
	}

	return l
}

func (m *Menu) getUserAction() (int, error) {
	fmt.Print("\nWhat'd you like to do?: ")

	text, err := m.reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("error reading answer given by user: %w", err)
	}

	text = strings.ReplaceAll(text, "\n", "")
	text = strings.ReplaceAll(text, "\r", "")

	num, err := strconv.Atoi(text)
	if err != nil {
		if m.menus[m.state].multiAction {
			return 0, fmt.Errorf("error converting user's answer to intager: %w", err)
		}

		return 0, nil
	}

	return num, nil
}

func (m *Menu) processUserAction(action int) {
	var act func()
	if m.menus[m.state].multiAction {
		act = m.menus[m.state].userActions[action]
	} else {
		act = m.menus[m.state].userActions[0]
	}

	if act != nil {
		act()
	}
}

// Run start's main menu
func (m *Menu) Build() giu.Layout {
	return giu.Layout{
		giu.Label("Welcome in tic-tac-go"),

		m.printMenu(),
	}
}
