package ttgmenu

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

const (
	menuQuestion = "What'd you like to do?"
)

// Menu represent's game's menu
type Menu struct {
	state          State
	reader         *bufio.Reader
	menus          map[State]*menuIndex
	boardW, boardH int
	chainLen       int
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
		Help: readMarkdown("README.md"),
		Settings: {
			"\n\tSettings:",
			"\t\t1) change board size",
			"\t\t0) back to main menu",
		},
	}

	cb := map[State]map[int]func(){
		MainMenu: {
			0: func() { os.Exit(0) },
			1: func() {
				var g *game.TTT

				rand.Seed(time.Now().UnixNano())
				// nolint:gomnd // two players in game
				r := rand.Intn(2) // nolint:gosec // it is ok

				switch r {
				case 0:
					g = game.NewTTT(ttgcommon.BaseBoardW, ttgcommon.BaseBoardH, m.chainLen, game.PlayerPerson, game.PlayerPC)
				case 1:
					g = game.NewTTT(ttgcommon.BaseBoardW, ttgcommon.BaseBoardH, m.chainLen, game.PlayerPC, game.PlayerPerson)
				}

				g.Run()
			},
			2: func() {
				game := game.NewTTT(m.boardW, m.boardH, m.chainLen, game.PlayerPerson, game.PlayerPerson)
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
			1: func() {
				w, err := m.getUserAction("Enter new board width")
				if err != nil {
					log.Fatal(err)
				}

				h, err := m.getUserAction("Enter new board height")
				if err != nil {
					log.Fatal(err)
				}

				l, err := m.getUserAction("Enter new chain len")
				if err != nil {
					log.Fatal(err)
				}

				m.boardW, m.boardH = w, h
				m.chainLen = l
			},
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
		state:    MainMenu,
		reader:   bufio.NewReader(os.Stdin),
		boardW:   ttgcommon.BaseBoardW,
		boardH:   ttgcommon.BaseBoardH,
		chainLen: ttgcommon.BaseChainLen,
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

func (m *Menu) getUserAction(question string) (int, error) {
	fmt.Print("\n" + question + ": ")

	text, err := m.reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("error reading user action: %w", err)
	}

	text = strings.ReplaceAll(text, "\n", "")

	num, err := strconv.Atoi(text)
	if err != nil {
		if m.menus[m.state].multiAction {
			return 0, fmt.Errorf("error reading user action: %w", err)
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
func (m *Menu) Run() {
	for {
		ttgcommon.Clear()
		fmt.Println("Welcome in tic-tac-go")
		m.printMenu()

		action, err := m.getUserAction(menuQuestion)
		if err != nil {
			log.Print(err)

			continue
		}

		m.processUserAction(action)
	}
}
