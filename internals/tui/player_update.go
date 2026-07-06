package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gopxl/beep/speaker"
)

var isLucked bool

func updatePlayer(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c":
			return m, tea.Quit
		// These keys should return to browser screen.
		case "q":
			m.screen = BrowserScreen
			speaker.Clear()
		// The "down" and "j" keys move the cursor down
		case "p":
			if isLucked {
				isLucked = false
				speaker.Unlock()
			} else {
				isLucked = true
				speaker.Lock()
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}
