package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/umarbek-x/LYRA/internals/player"
)

func updateBrowser(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.browser.cursor > 0 {
				m.browser.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.browser.cursor < len(m.browser.music)-1 {
				m.browser.cursor++
			}

		// The "enter" key and the space bar toggle the selected state
		// for the item that the cursor is pointing at.
		case "enter", "space":
			music := player.GetFile(m.browser.directory, m.browser.music[m.browser.cursor])
			m.player.Song = music.Title

			go func() {
				// play the music
				for i := 0 + m.browser.cursor; i < len(m.browser.music); i++ {
					player.Play(m.browser.directory + m.browser.music[i])
					if i == len(m.browser.music)-1 {
						i = 0
						m.browser.cursor = 0
					}
				}
			}()
			// change screen to player screen
			m.screen = PlayerScreen
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}
