package main

import (
	"log"

	tea "charm.land/bubbletea/v2"
	"github.com/umarbek-x/LYRA/internals/tui"
)

func main() {
	p := tea.NewProgram(tui.InitialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
