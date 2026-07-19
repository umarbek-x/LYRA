package tui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/umarbek-x/LYRA/internals/metadata"
)

type Model struct {
	Music     []string
	Cursor    int
	Directory string
	Current   string
}

func InitialModel() Model {
	dir := "/home/udev/Music/lyra"
	return Model{
		Directory: dir,
		Cursor:    0,
		Music:     metadata.GetMusicNames(dir),
	}
}

func (s Model) Init() tea.Cmd {
	// no io for now
	return nil
}
