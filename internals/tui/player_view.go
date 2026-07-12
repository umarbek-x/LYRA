package tui

import (
	"fmt"
	"log"

	"github.com/umarbek-x/LYRA/internals/lyrics"
	"github.com/umarbek-x/LYRA/internals/metadata"
	"github.com/umarbek-x/LYRA/internals/player"
)

func PlayerView(m Model) (s string) {
	// The header
	s = "🎵 " + m.player.Song + "\n\n"
	cache, err := metadata.ReadCache()
	if err != nil {
		log.Fatal(err)
	}

	if lrc, ok := cache[m.player.Song]; ok {
		for i := 0; i < len(lrc.Lines); i++ {
			s += fmt.Sprintf("%d. %s - %s\n", i+1, lrc.Lines[i].Time, lrc.Lines[i].Text)
		}
	} else {
		file := player.GetFile(m.browser.directory, m.browser.music[m.browser.cursor])

		formatedLyrics, err := lyrics.Parse(file.Title, file.Artist, file.Duration)
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < len(formatedLyrics.Lines); i++ {
			s += fmt.Sprintf("%d. %s - %s\n", i+1, formatedLyrics.Lines[i].Time, formatedLyrics.Lines[i].Text)
		}
		lyrics.SaveTOFile(formatedLyrics)
	}

	s += "\nPress q to quit || p to pause or resume.\n"

	// Send the UI for rendering
	return s
}
