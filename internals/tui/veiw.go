package tui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

func (m Model) View() tea.View {
	// The header
	s := "Lyra Music Streaming CLI Platform\n\n"

	// Iterate over our music
	for i, music := range m.Music {

		if m.Cursor != i {
			s += fmt.Sprintf("%d. [%s] %s\n", i+1, " ", music)
		}
		s += fmt.Sprintf("%d. [%s] %s\n", i+1, ">", music)

	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return tea.NewView(s)
}
