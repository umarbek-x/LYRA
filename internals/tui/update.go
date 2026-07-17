package tui

import tea "charm.land/bubbletea/v2"

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// is it a key press
	case tea.KeyPressMsg:
		switch msg.String() {
		// if the user presses Ctrl+c or q the program will quit
		case "Ctrl+c", "q":
			return m, tea.Quit

		}
	}
	return m, nil
}
