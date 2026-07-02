package lyrics

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"net/url"
	"time"
)

type Lyrics struct {
	Id           int64   `json:"id"`
	Name         string  `json:"name"`
	Artist_Name  string  `json:"artistName"`
	AlbomName    string  `json:"albumName"`
	Duration     float64 `json:"duration"`
	SyncedLyrics string  `json:"syncedLyrics"`
}

func Parse(trackName, artist_name string, duration time.Duration) Lyrics {
	res, err := http.Get(fmt.Sprintf("https://lrclib.net/api/search?track_name=%s&artist_name=%s", url.QueryEscape(trackName), url.QueryEscape(artist_name)))
	if err != nil {
		log.Println(err)
		return Lyrics{}
	}
	defer res.Body.Close()

	var lyricses []Lyrics
	if err := json.NewDecoder(res.Body).Decode(&lyricses); err != nil {
		log.Println(err)
		return Lyrics{}
	}

	if len(lyricses) == 0 {
		return Lyrics{}
	}

	target := duration.Seconds()

	best := lyricses[0]
	bestDiff := math.Abs(best.Duration - target)

	for _, lyric := range lyricses[1:] {
		diff := math.Abs(lyric.Duration - target)

		if diff < bestDiff {
			best = lyric
			bestDiff = diff
		}
	}

	return best
}

