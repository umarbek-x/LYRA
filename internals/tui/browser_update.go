package tui

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/umarbek-x/LYRA/internals/lyrics"
	"github.com/umarbek-x/LYRA/internals/metadata"
	"github.com/umarbek-x/LYRA/internals/player"
)

func updateBrowser(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.browser.cursor > 0 {
				m.browser.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.browser.cursor < len(m.browser.music)-1 {
				m.browser.cursor++
			}

		// The "enter" key and the space bar toggle the selected state
		// for the item that the cursor is pointing at.
		case "enter", "space":
			m.player.Song = m.browser.music[m.browser.cursor]
			// get the file information
			file := player.GetFile(m.browser.directory, m.browser.music[m.browser.cursor])
			// get the file from cache .json
			var flrc lyrics.JsonLyrics
			cache, err := metadata.ReadCache()
			if err != nil {
				log.Fatal(err)
			}
			// find the expected music
			for _, value := range cache {
				if value.Name == file.Title && value.Artist_Name == file.Artist {
					flrc = value
				}
			}
			if flrc.Name == "" {

				// parce the lyrics if there is no lyrics in cache
				plrc, err := lyrics.Parse(file.Title, file.Artist, file.Duration)
				if err != nil {
					log.Fatal(err)
				}
				// format the lyrics
				flrc, err = lyrics.ParseLyricsLineWithTime(plrc)
				if err != nil {
					log.Fatal(err)
				}
				// save to file
				err = lyrics.SaveTOFile(flrc)
				if err != nil {
					log.Fatal(err)
				}
			}
			m.player.Lyrics = flrc
			go func() {

				// play the music
				player.Play(m.browser.directory + m.browser.music[m.browser.cursor])
			}()
			m.screen = PlayerScreen
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}
