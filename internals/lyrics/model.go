package lyrics

import "time"

type Lyrics struct {
	Id           int64   `json:"id"`
	Name         string  `json:"name"`
	Artist_Name  string  `json:"artistName"`
	AlbomName    string  `json:"albumName"`
	Duration     float64 `json:"duration"`
	SyncedLyrics string  `json:"syncedLyrics"`
}

type JsonLyrics struct {
	Name        string  `json:"name"`
	Artist_Name string  `json:"artistName"`
	Duration    float64 `json:"duration"`
	Lines       []Line  `json:"line_w_time"`
}

type Line struct {
	Time time.Duration `json:"duration"`
	Text string        `json:"text"`
}
