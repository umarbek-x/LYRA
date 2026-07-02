package metadata

import (
	"encoding/json"
	"os"

	"github.com/umarbek-x/LYRA/internals/lyrics"
)

func ReadCache() ([]lyrics.JsonLyrics, error) {
	var cache []lyrics.JsonLyrics
	// Read existing cache if it exists.
	file, err := os.ReadFile("cache.json")
	if err == nil {
		err := json.Unmarshal(file, &cache)
		if err != nil {
			return nil, err
		}
	}

	return cache, nil
}
