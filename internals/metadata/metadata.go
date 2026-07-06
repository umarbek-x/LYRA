package metadata

import (
	"encoding/json"
	"os"

	"github.com/umarbek-x/LYRA/internals/lyrics"
)

func ReadCache() (cache map[string]lyrics.JsonLyrics, err error) {
	cache = make(map[string]lyrics.JsonLyrics)
	err = nil
	var lyrics []lyrics.JsonLyrics
	// Read existing cache if it exists.
	file, err := os.ReadFile("cache.json")
	if err == nil {
		err = json.Unmarshal(file, &lyrics)
		if err != nil {
			return nil, err
		}
		for i, lrc := range lyrics {
			cache[lrc.Name] = lyrics[i]
		}
	}

	return cache, nil
}
