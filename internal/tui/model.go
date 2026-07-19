package tui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/umarbek-x/LYRA/internal/metadata"
)

type Model struct {
	Cursor int
	Music  Music
}

type Music struct {
	Songs     []string
	Directory string
	Current   string
	IsPaused  bool
}

func InitialModel() Model {
	dir := "/home/udev/Music/lyra"
	return Model{
		Cursor: 0,
		Music: Music{
			Songs:     metadata.GetMusicNames(dir),
			Directory: dir,
		},
	}
}

func (s Model) Init() tea.Cmd {
	// no io for now
	return nil
}
