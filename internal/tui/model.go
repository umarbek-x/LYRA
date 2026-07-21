package tui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/umarbek-x/LYRA/internal/metadata"
	"github.com/umarbek-x/LYRA/pkg/config"
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
	cfg := config.LoadConfig()
	return Model{
		Cursor: 0,
		Music: Music{
			Songs:     metadata.GetMusicNames(cfg.Dir),
			Directory: cfg.Dir,
		},
	}
}

func (s Model) Init() tea.Cmd {
	// no io for now
	return nil
}
