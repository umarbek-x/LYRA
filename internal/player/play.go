package player

import (
	"os"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
	errorhanler "github.com/umarbek-x/LYRA/pkg/errorHanler"
)

func Play(song string) {
	data, err := os.Open(song)
	errorhanler.CheckError(err)

	streamer, format, err := mp3.Decode(data)
	errorhanler.CheckError(err)
	defer streamer.Close()

	speaker.Clear()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
