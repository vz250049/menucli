package cursor

import (
	"fmt"
	"github.com/buger/goterm"
	"github.com/pkg/term"
	"log"
	"os"
)

var j byte = 106
var k byte = 107
var q byte = 113
var up byte = 65
var down byte = 66
var escape byte = 27
var enter byte = 13

var keys = map[byte]bool{
	up:   true,
	down: true,
	j:    true,
	k:    true,
	q:    true,
}

type Menu struct {
	Prompt    string
	CursorPos int
	MenuItems []*MenuItem
}

type MenuItem struct {
	Text    string
	ID      string
	SubMenu *Menu
}

func NewMenu(prompt string) *Menu {
	return &Menu{
		Prompt:    prompt,
		MenuItems: make([]*MenuItem, 0),
	}
}

func (m *Menu) AddItem(option string, id string) *Menu {
	menuItem := &MenuItem{
		Text: option,
		ID:   id,
	}

	m.MenuItems = append(m.MenuItems, menuItem)
	return m
}

func (m *Menu) renderMenuItems(redraw bool) {
	if redraw {
		fmt.Printf("\033[%dA", len(m.MenuItems)-1)
	}

	for index, menuItem := range m.MenuItems {
		var newline = "\n"
		if index == len(m.MenuItems)-1 {
			newline = ""
		}

		menuItemText := menuItem.Text
		cursor := "  "
		if index == m.CursorPos {
			cursor = goterm.Color("> ", goterm.YELLOW)
			menuItemText = goterm.Color(menuItemText, goterm.YELLOW)
		}

		fmt.Printf("\r%s %s%s", cursor, menuItemText, newline)
	}
}

func (m *Menu) Display() string {
	defer func() {
		// Show cursor again.
		fmt.Printf("\033[?25h")
	}()

	totalRows := len(m.MenuItems) + 2
	fmt.Printf("\033[8;%d;180t", totalRows)

	fmt.Print("\033[H\033[2J")
	fmt.Printf("%s\n", goterm.Color(goterm.Bold(m.Prompt)+":", goterm.CYAN))

	m.renderMenuItems(false)

	// Turn the terminal cursor off
	fmt.Printf("\033[?25l")

	for {
		keyCode := getInput()
		if keyCode == escape { //todo: find different function for escape key
			fmt.Printf("\033[?25h")
			os.Exit(0)
		} else if keyCode == enter {
			menuItem := m.MenuItems[m.CursorPos]
			fmt.Println("\r")
			return menuItem.ID
		} else if keyCode == up {
			m.CursorPos = (m.CursorPos + len(m.MenuItems) - 1) % len(m.MenuItems)
			m.renderMenuItems(true)
		} else if keyCode == j {
			m.CursorPos = (m.CursorPos + 1) % len(m.MenuItems)
			m.renderMenuItems(true)
		} else if keyCode == k {
			m.CursorPos = (m.CursorPos + len(m.MenuItems) - 1) % len(m.MenuItems)
			m.renderMenuItems(true)
		} else if keyCode == down {
			m.CursorPos = (m.CursorPos + 1) % len(m.MenuItems)
			m.renderMenuItems(true)		
		}else if keyCode == q {
			fmt.Printf("\033[?25h")
			os.Exit(0)
		}
	}
}


func getInput() byte {
	t, _ := term.Open("/dev/tty")

	err := term.RawMode(t)
	if err != nil {
		log.Fatal(err)
	}

	var read int
	readBytes := make([]byte, 3)
	read, err = t.Read(readBytes)

	err = t.Restore()
	if err != nil {
		return 0
	}
	err = t.Close()
	if err != nil {
		return 0
	}

	if read == 3 {
		if _, ok := keys[readBytes[2]]; ok {
			return readBytes[2]
		}
	} else {
		return readBytes[0]
	}

	return 0
}
