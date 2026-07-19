package tui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/umarbek-x/LYRA/internals/player"
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
			if m.Cursor < len(m.Music)-1 {
				m.Cursor++
			}
		case "enter", "space":
			go func() {
				m.Current = m.Music[m.Cursor]
				player.Play(m.Directory + "/" + m.Current)
				m.Current = ""
			}()
		}

	}
	return m, nil
}
