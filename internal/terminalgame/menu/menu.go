package menu

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	osinfo "gist.github.com/2335e953b45f46591839a21c502588ea.git"
	"github.com/jaytaylor/html2text"
	"github.com/pkg/browser"
	"github.com/russross/blackfriday"

	"github.com/gucio321/tic-tac-go/internal/terminalgame/game"
	"github.com/gucio321/tic-tac-go/internal/terminalgame/utils"
	"github.com/gucio321/tic-tac-go/pkg/core/board"
	"github.com/gucio321/tic-tac-go/pkg/core/players/ttgplayer"
)

const githubURL = "https://github.com/gucio321/tic-tac-go"

type menuPosition byte

const (
	mainMenu menuPosition = iota
	settingsMenu
)

type settings struct {
	chainLen,
	width,
	height byte
}

// Menu represents a game menu.
type Menu struct {
	*settings
	pages  []*menuPage
	done   bool
	pos    menuPosition
	reader *bufio.Reader
	readme *[]byte
}

// New creates a new menu
// readme is a README.md file (pass nil if no readme).
func New(readme []byte) *Menu {
	result := &Menu{
		settings: &settings{
			board.BaseChainLen,
			board.BaseBoardW,
			board.BaseBoardH,
		},
		done:   false,
		pos:    mainMenu,
		reader: bufio.NewReader(os.Stdin),
		readme: &readme,
	}

	result.loadMenu()

	return result
}

// Run runs the menu.
func (m *Menu) Run() {
	for !m.done {
		utils.Clear()
		fmt.Println(m.currentPage())

		num, err := m.getUserAction("What'd you like to do?")
		if err != nil {
			continue
		}

		if num < 0 || int16(num) > m.currentPage().Max() {
			fmt.Printf("You must enter number from 0 to %d", m.currentPage().Max())

			continue
		}

		m.currentPage().Exec(int16(num))
	}
}

func (m *Menu) currentPage() *menuPage {
	return m.pages[m.pos]
}

func (m *Menu) loadMenu() {
	m.pages = []*menuPage{
		{
			"main menu",
			[]*menuIndex{
				{1, "PvC game", m.runPVC},
				{2, "PvP game", m.runPVP},
				{3, "Demo", m.runDemo},
				{4, "Settings", func() { m.pos = settingsMenu }},
				{5, "Help", m.printHelp},
				{6, "README", m.printReadme},
				{7, "website", m.openWebsite},
				{8, "Report Bug on GitHub", m.reportBug},
				{0, "Exit", func() { m.done = true }},
			},
		},
		{
			"Settings",
			[]*menuIndex{
				{1, "Change board size", m.changeBoardSize},
				{2, "Reset board size", m.resetBoardSize},
				{0, "Back", func() { m.pos = mainMenu }},
			},
		},
	}
}

func (m *Menu) runPVP() {
	pvp := game.NewTTG(m.width, m.height, m.chainLen, ttgplayer.PlayerPerson, ttgplayer.PlayerPerson)
	pvp.Run()
}

func (m *Menu) runPVC() {
	var g *game.TTG

	rand.Seed(time.Now().UnixNano())
	// nolint:gomnd // two players in game
	r := rand.Intn(2) // nolint:gosec // it is ok

	switch r {
	case 0:
		g = game.NewTTG(m.width, m.height, m.chainLen, ttgplayer.PlayerPerson, ttgplayer.PlayerPC)
	case 1:
		g = game.NewTTG(m.width, m.height, m.chainLen, ttgplayer.PlayerPC, ttgplayer.PlayerPerson)
	}

	g.Run()
}

func (m *Menu) runDemo() {
	demo := game.NewTTG(m.width, m.height, m.chainLen, ttgplayer.PlayerPC, ttgplayer.PlayerPC)
	demo.Run()
}

func (m *Menu) changeBoardSize() {
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

	m.width, m.height = byte(w), byte(h)
	m.chainLen = byte(l)
}

func (m *Menu) resetBoardSize() {
	m.width, m.height = board.BaseBoardW, board.BaseBoardH
	m.chainLen = board.BaseChainLen
	_, _ = m.getUserAction("Board width and height was set to default\nPress ENTER to continue")
}

func (m *Menu) printHelp() {
	utils.Clear()
	fmt.Println(strings.Join([]string{
		"TicTacToe Version 1",
		"Copyright (C) 2021 by gucio321 (https://github.com/gucio321)",
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
	}, "\n"),
	)

	_, _ = m.getUserAction("Press ENTER to continue")
}

func (m *Menu) printReadme() {
	var err error

	html := blackfriday.MarkdownBasic(*m.readme)

	text, err := html2text.FromString(string(html), html2text.Options{PrettyTables: true})
	if err != nil {
		fmt.Printf(
			"Unable to convert readme's html to text: %v\n%s", err,
			"Please raport it on https://github.com/gucio321/tic-tac-go",
		)
	}

	fmt.Println(text)

	_, _ = m.getUserAction("Press ENTER to continue")
}

func (m *Menu) openWebsite() {
	err := browser.OpenURL(githubURL)
	if err != nil {
		log.Println(err)
	}
}

func (m *Menu) reportBug() {
	var err error

	osInfo := osinfo.NewOS()

	body := []string{
		"%23%23 Describe the bug",
		"A clear and concise description of what the bug is.",
		"",
		"%23%23 To Reproduce",
		"Steps to reproduce the behavior:",
		"1. Go to '...'",
		"2. Click on '....'",
		"3. Scroll down to '....'",
		"4. See error",
		"",
		"%23%23 Expected behavior",
		"A clear and concise description of what you expected to happen.",
		"",
		"%23%23 Screenshots",
		"If applicable, add screenshots to help explain your problem.",
		"",
		"%23%23 Desktop:",
		"- OS: " + osInfo.Name,
		"- Version: " + osInfo.Version,
		"- Arch: " + osInfo.Arch,
		"- Go version: " + runtime.Version(),
		"",
		"%23%23 Additional context",
		"Add any other context about the problem here.",
	}

	err = browser.OpenURL("https://github.com/gucio321/tic-tac-go/issues/new?body=" + strings.Join(body, "%0D"))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *Menu) getUserAction(question string) (int, error) {
	fmt.Print("\n" + question + ": ")

	text, err := m.reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("error reading user action: %w", err)
	}

	text = strings.ReplaceAll(text, "\n", "")
	text = strings.ReplaceAll(text, "\r", "")

	num, err := strconv.Atoi(text)
	if err != nil {
		return num, fmt.Errorf("error converting user answer: %w", err)
	}

	return num, nil
}
