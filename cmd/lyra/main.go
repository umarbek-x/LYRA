package main

import (
	"fmt"
	"log"

	tea "charm.land/bubbletea/v2"
	"github.com/umarbek-x/LYRA/internal/tui"
	"github.com/umarbek-x/LYRA/pkg/config"
)

func main() {
	// check if the directory exists
	config.CheckMusicDir()

	fmt.Print("\033c") // Reset terminal
	// run the tui
	p := tea.NewProgram(tui.InitialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

}
