package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/umarbek-x/LYRA/internals/player"
)

type model struct {
	music_list  []string
	cursor      int
	directory   string
	currentsong string
	playing     bool
}

func InitModel() model {
	return model{
		music_list: player.ListFiles("/home/udev/Music/lyra"),
		directory:  "/home/udev/Music/lyra/",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
