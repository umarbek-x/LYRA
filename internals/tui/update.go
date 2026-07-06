package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.screen {

	case BrowserScreen:
		return updateBrowser(m, msg)

	case PlayerScreen:

		return updatePlayer(m, msg)
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}
