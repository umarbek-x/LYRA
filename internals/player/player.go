package player

import (
	"log"
	"os"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

var speakerInitialized bool

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

	// prepare the speakers for the mp3
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speakerInitialized = true

	// make a chanal of bool to inform when the music is finished
	done := make(chan bool)

	// clear the current playing audio
	speaker.Clear()

	// play the music
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
