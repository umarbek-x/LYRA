package metadata

import (
	"os"
	"strings"

	errh "github.com/umarbek-x/LYRA/pkg/errorHanler"
)

func GetMusicNames(directory string) []string {

	// read the directory
	enteries, err := os.ReadDir(directory)
	errh.CheckError(err)

	// initialize a songs variable of lise of strings
	var names []string

	// loop over evry entery
	for _, v := range enteries {
		// check if it is type mp3 but with file name not its type, it might be not accurat but for this level it is fine
		if strings.Contains(v.Name(), ".mp3") {
			// trim suffix .mp3 so it is more orgonized
			songName := strings.TrimSuffix(v.Name(), ".mp3")
			// append to the list above
			names = append(names, songName)
		}
	}
	return names
}
