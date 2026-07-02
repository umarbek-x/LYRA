package lyrics

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var re = regexp.MustCompile(`^\[(\d{2}):(\d{2})\.(\d{2})\](.*)$`)

// parse the lyric of the music from api (lrclib.net)
func Parse(trackName, artist_name string, duration time.Duration) (Lyrics, error) {
	res, err := http.Get(fmt.Sprintf("https://lrclib.net/api/search?track_name=%s&artist_name=%s", url.QueryEscape(trackName), url.QueryEscape(artist_name)))
	if err != nil {
		return Lyrics{}, err
	}
	defer res.Body.Close()

	var lyricses []Lyrics
	if err := json.NewDecoder(res.Body).Decode(&lyricses); err != nil {
		return Lyrics{}, err
	}

	if len(lyricses) == 0 {
		return Lyrics{}, err
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

	if best.SyncedLyrics == "" {
		return Lyrics{}, errors.New("Unable to parse the syncedLyrics")
	}

	best.Name = trackName
	best.AlbomName = artist_name

	return best, nil
}

func ParseLyricsLineWithTime(lyrics Lyrics) (JsonLyrics, error) {

	syncedLyrics := lyrics.SyncedLyrics
	lines := strings.Split(syncedLyrics, "\n")
	var lyricsLine []Line
	for _, line := range lines {
		line := strings.TrimSpace(line)
		if line == "" {
			continue
		}

		match := re.FindStringSubmatch(line)
		if match == nil {
			continue
		}

		minute, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatalln("Unable to extract minute - ", err)
		}
		second, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatalln("Unable to extract second - ", err)
		}
		msec, err := strconv.Atoi(match[3])
		if err != nil {
			log.Fatalln("Unable to extract millisecond - ", err)
		}

		time := time.Duration(minute)*time.Minute + time.Duration(second)*time.Second + time.Duration(msec)*time.Millisecond
		lyricsLine = append(lyricsLine, Line{Time: time, Text: match[4]})
	}

	return JsonLyrics{
		Name:        lyrics.Name,
		Artist_Name: lyrics.Artist_Name,
		Duration:    lyrics.Duration,
		Lines:       lyricsLine,
	}, nil
}

func SaveTOFile(content JsonLyrics) error {

	var cache []JsonLyrics
	// Read existing cache if it exists.
	file, err := os.ReadFile("cache.json")
	if err == nil {
		json.Unmarshal(file, &cache)
	}

	cache = append(cache, content)
	// write everything back
	data, err := json.MarshalIndent(cache, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile("cache.json", data, 0644)
}

