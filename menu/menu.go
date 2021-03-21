package menu

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

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

func (m *Menu) getMenuData(state State) (lines []string, actions map[int]func()) {
	text := map[State][]string{
		MainMenu: {
			"\nMainMenu",
			"\t1) start Player VS PC game",
			"\t2) start Player VS Player game",
			"\t3) settings",
			"\t4) Help",
			"\t0) exit",
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
			"\n\tSettings:",
			"\t\t0) back to main menu",
		},
	}

	cb := map[State]map[int]func(){
		MainMenu: {
			0: func() { os.Exit(0) },
			1: func() {
				var g *game.TTT

				rand.Seed(time.Now().UnixNano())
				r := rand.Intn(2)

				switch r {
				case 0:
					g = game.NewTTT(game.PlayerPerson, game.PlayerPC)
				case 1:
					g = game.NewTTT(game.PlayerPC, game.PlayerPerson)
				}

				g.Run()
			},
			2: func() {
				game := game.NewTTT(game.PlayerPerson, game.PlayerPerson)
				game.Run()
			},
			3: func() {
				m.state = Settings
			},
			4: func() {
				m.state = Help
			},
		},
		Help: {
			0: func() {
				m.state = MainMenu
			},
		},
		Settings: {
			0: func() {
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
	userActions map[int]func()
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

func (m *Menu) printMenu() {
	lines := m.menus[m.state].lines
	if lines != nil {
		fmt.Println(strings.Join(lines, "\n"))
	}
}

func (m *Menu) getUserAction() (int, error) {
	fmt.Print("\nWhat'd you like to do?: ")

	text, err := m.reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	text = strings.ReplaceAll(text, "\n", "")

	num, err := strconv.Atoi(text)
	if err != nil {
		if m.menus[m.state].multiAction {
			return 0, err
		} else {
			return 0, nil
		}
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
func (m *Menu) Run() {
	for {
		ttgcommon.Clear()
		fmt.Println("Welcome in tic-tac-go")
		m.printMenu()
		action, err := m.getUserAction()
		if err != nil {
			log.Print(err)

			continue
		}

		m.processUserAction(action)
	}
}
