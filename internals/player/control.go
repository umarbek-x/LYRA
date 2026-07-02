package player

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/hajimehoshi/go-mp3"
)

type File struct {
	Title    string
	Artist   string
	Duration time.Duration
}

// lists the mp3 files
func ListFiles(directory string) []string {
	// read the directory
	enteries, err := os.ReadDir(directory)
	if err != nil {
		log.Fatalln("Unable  to read directory - ", err)
	}

	var fileNames []string
	// list and read the file in the directory
	for _, entery := range enteries {
		// validates
		if !entery.IsDir() && strings.HasSuffix(entery.Name(), ".mp3") {
			fileNames = append(fileNames, entery.Name())
		}
	}
	return fileNames
}

// extract the needed information from the file
func GetFile(directory, music string) File {
	enteries, err := os.ReadDir(directory)
	if err != nil {
		log.Fatalln("Unable to read directory of the music - ", err)
	}

	for _, entery := range enteries {
		if entery.Name() == music {

			// extract the title and the artist name from the file name, by rule the name should be Artist - Title
			parts := strings.Split(entery.Name(), " - ")
			Artist := parts[0]
			Title := parts[1]

			// extract the duration of the file
			duration, err := mp3Duration(directory + music)
			if err != nil {
				log.Fatal("unable to extract mp3 duration - ", err)
			}

			return File{
				Title:    strings.TrimSuffix(Title, ".mp3"),
				Artist:   Artist,
				Duration: duration,
			}
		}
	}
	return File{}
}

// extract the duration of the music
func mp3Duration(path string) (time.Duration, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, nil
	}
	defer f.Close()

	decode, err := mp3.NewDecoder(f)
	if err != nil {
		return 0, nil
	}

	length := decode.Length() // bytes of the  decoded pcm

	const (
		bytesPerSample = 2 // 16-bit PCM
		channels       = 2 // stereo
	)

	samples := length / (bytesPerSample * channels)
	seconds := float64(samples) / float64(decode.SampleRate())

	return time.Duration(seconds * float64(time.Second)), nil
}
