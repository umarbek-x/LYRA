package tui

func (m Model) View() string {
	switch m.screen {
	case BrowserScreen:
		return browserView(m)
	case PlayerScreen:
		return PlayerView(m)
	}
	return ""
}
