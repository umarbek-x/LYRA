package tui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

var (
	AppNameStyle    = lipgloss.NewStyle().Background(lipgloss.Color("99")).Padding(0, 1)
	faintStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("105")).Faint(true)
	enumeratorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("105")).MarginRight(1)
)

func (m Model) View() tea.View {

	// The header
	s := AppNameStyle.Render("Lyra Music Streaming CLI Platform") + "\n\n"
	// s := "Lyra Music Streaming CLI Platform \n\n"

	// Iterate over our music
	for i, music := range m.Music.Songs {
		var cursor = " "
		if m.Cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("[%s] %s\n", enumeratorStyle.Render(cursor), strings.TrimSuffix(music, ".mp3"))

	}

	// The footer
	s += "\n" + faintStyle.Render("Press q to quit || Press p to pause or resume.") + "\n"

	// Send the UI for rendering
	return tea.NewView(s)
}
