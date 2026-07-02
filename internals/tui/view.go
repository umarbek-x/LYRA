package tui

import (
	"fmt"
)

func (m model) View() string {
	// The header
	s := "Wellcome to LYRA\n\nSelect Music you want to listen to\n\n"

	// Iterate over our musics
	for i, music := range m.music_list {

		// Is the cursor pointing at this music?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = "*" // cursor!
		}

		// Render the row
		s += fmt.Sprintf(" [%s] %s\n", cursor, music)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
