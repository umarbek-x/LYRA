package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/umarbek-x/LYRA/internals/lyrics"
	"github.com/umarbek-x/LYRA/internals/player"
)

type Screen int

const (
	BrowserScreen Screen = iota
	PlayerScreen
)

type model struct {
	screen  Screen
	browser Browser_Model
	player  Player_Model
}

type Browser_Model struct {
	music     []string
	cursor    int
	directory string
}

type Player_Model struct {
	Song string

	Position time.Duration
	Duration time.Duration

	Lyrics      lyrics.JsonLyrics
	CurrentLine int

	Paused bool
}

func InitModel() model {
	return model{
		screen: BrowserScreen,
		browser: Browser_Model{
			music:     player.ListFiles("/home/udev/Music/lyra"),
			directory: "/home/udev/Music/lyra/",
		},
		player: Player_Model{
			Song:     "",
			Position: 0,
			Duration: 0,
			Lyrics:   lyrics.JsonLyrics{},
			CurrentLine: 0,
			Paused: false,
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
