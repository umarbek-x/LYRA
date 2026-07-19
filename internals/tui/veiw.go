package tui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m Model) View() tea.View {
	// The header
	s := "Lyra Music Streaming CLI Platform\n\n"

	// Iterate over our music
	for i, music := range m.Music {
		var cursor = " "
		if m.Cursor == i {
			cursor = ">"
		}
		
		s += fmt.Sprintf("[%s] %s\n", cursor, strings.TrimSuffix(music, ".mp3"))

	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return tea.NewView(s)
}
