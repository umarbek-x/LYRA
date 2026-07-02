package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/umarbek-x/LYRA/internals/tui"
)

func main() {
	// tui.ListFiles("/home/udev/Music/lyra")
	m := tui.InitModel()
	program := tea.NewProgram(m)
	if _, err := program.Run(); err != nil {
		log.Fatalln("Unable to run tui\nERROR: ", err)
	}

}
