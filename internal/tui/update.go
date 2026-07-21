package tui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/gopxl/beep/speaker"
	"github.com/umarbek-x/LYRA/internal/player"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// is it a key press
	case tea.KeyPressMsg:
		switch msg.String() {
		// if the user presses Ctrl+c or q the program will quit
		case "Ctrl+c", "q":
			return m, tea.Quit
		case "up":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down":
			if m.Cursor < len(m.Music.Songs)-1 {
				m.Cursor++
			}
		case "p":
			if m.Music.IsPaused {
				speaker.Resume()
				m.Music.IsPaused = false
			} else {
				speaker.Suspend()
				m.Music.IsPaused = true
			}
		case "enter", "space":

			go func() {
				m.Music.Current = m.Music.Songs[m.Cursor]
				player.Play(m.Music.Directory + "/" + m.Music.Current)
				m.Music.Current = ""
			}()
		}

	}
	return m, nil
}
