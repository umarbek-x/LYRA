package tui

func (m model) View() string {
	switch m.screen {
	case BrowserScreen:
		return browserView(m)
	case PlayerScreen:
		return playerView(m)
	}
	return ""
}
