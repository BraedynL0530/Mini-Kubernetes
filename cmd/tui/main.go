package main

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/BraedynL0530/Mini-Kubernetes/pkg/types"
)

//ToDO: add screens, finish sub menus and options, this is skeleton untill i have actual functions

type Screen int

const (
	mainScreen Screen = iota
	GetSubMenu
	ApplyResultScreen
)

// Bubble Tea Docs i stole and edited
type model struct {
	header        string
	textIcon      []string
	choices       []types.Cat
	cursor        int              // which to-do list item our cursor is pointing at
	selected      map[int]struct{} // which to-do items are selected
	getSubChoices []string         // items on the get sub-menu
	currentScreen Screen           // current screen being displayed
	outputMessage string           // message to display on the apply result screen
}

// add a method for like -r path/to/ymal file

//re adding comments: again mainly copied from bubble tui

func initialModel() model {
	return model{
		header: "                  Meowbernetes", //adjust whitespace
		textIcon: []string{
			"       |\\      _,,,---,,_",
			"ZZZzz  /,`.-'`'    -.  ;-;;,_",
			"      |,4-  ) )-,_. ,\\ (  `'-'",
			"       '---''(_/--'  `- '\\_)",
		},
		choices:       []types.Cat{}, // Todo: make sure to only have name and status. possibly
		getSubChoices: []string{"fetch logs", "fetch metrics", "back to main"},

		// A map which indicates which choices are selected. We're using
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyPressMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the space bar toggle the selected state
		// for the item that the cursor is pointing at.
		case "enter", "space":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() tea.View {
	// The header
	s := strings.Join(m.textIcon, "\n") + "\n" + m.header

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	return tea.NewView(s)
}
