package tui

func playerView(m model) string {
	// The header
	s := "🎵 " + m.player.Song + "\n\n"

	current := m.player.CurrentLine

	if current > 0 {
		s += m.player.Lyrics.Lines[current-1].Text + "\n"
	}

	s += "> " + m.player.Lyrics.Lines[current].Text + "\n"
	// The footer

	if current+1 < len(m.player.Lyrics.Lines) {
		s += m.player.Lyrics.Lines[current+1].Text + "\n"
	}
	
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
