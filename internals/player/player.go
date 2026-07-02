package player

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

var (
	currentStreamer    beep.StreamSeekCloser
	currentFormat      beep.Format
	speakerInitialized bool

	mu sync.RWMutex
)

func Play(music string) {
	// open the file
	file, err := os.Open(music)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read the mp3 file and make a streamer
	streamer, format, err := mp3.Decode(file)
	if err != nil {
		log.Fatalln(err)
	}
	defer streamer.Close()

	mu.Lock()
	currentFormat = format
	currentStreamer = streamer
	mu.Unlock()

	// prepare the speakers for the mp3
	if !speakerInitialized {
		speaker.Init(currentFormat.SampleRate, currentFormat.SampleRate.N(time.Second/10))
		speakerInitialized = true
	}

	// make a chanal of bool to inform when the music is finished
	done := make(chan bool)

	// clear the current playing audio
	speaker.Clear()

	// play the music
	speaker.Play(beep.Seq(currentStreamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}

func Position() time.Duration {
	mu.RLock()
	defer mu.RUnlock()

	if currentStreamer == nil {
		return 0
	}

	samples := currentStreamer.Position()

	return time.Duration(samples) * time.Second / time.Duration(currentFormat.SampleRate)
}
