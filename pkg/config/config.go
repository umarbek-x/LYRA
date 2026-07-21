package config

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	errorhanler "github.com/umarbek-x/LYRA/pkg/errorHanler"
)

type Config struct {
	Dir string
}

var (
	dirrequest     = "The CLI platform requires Music Directory from which it will read files and play \nEnter the music Directory path 'home/user/music/lyra': "
	configfilepath = "config.json"
)

// check if the file exists
func IsFileExists(path string) (bool, error) {
	// get file information
	_, err := os.Lstat(path)
	if err != nil {
		// check if the file not exists
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		// another error
		return false, err
	}
	// if exists
	return true, nil
}

// check if directory exists
func ISDirExists(path string) (bool, error) {
	_, err := os.ReadDir(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// check if the music dir exists or save
func CheckMusicDir() {

	// prepare a variable to store json data
	var cfg Config

	// check if the file exists or no
	exists, err := IsFileExists(configfilepath)
	errorhanler.CheckError(err)

	if !exists {
		reader := bufio.NewReader(os.Stdin)
		var input = ""
		for {
			// first make a request so it is Enter the music file path:
			fmt.Print(dirrequest)
			input, err = reader.ReadString('\n') // and then read what is written after the request
			errorhanler.CheckError(err)

			// check if the input(music directory path) exists or valid
			input = strings.TrimSpace(input)
			exists, err = ISDirExists(input)
			errorhanler.CheckError(err)
			if exists {
				break
			} else {
				fmt.Printf("\nInvalid directory! Please insert valid Directory from which the platform can read audio\n\n")
			}
		}
		// assign the music path to the struct and save in config file
		cfg.Dir = input

		// create config file
		file, err := os.OpenFile(configfilepath, os.O_CREATE|os.O_WRONLY, 0666)
		errorhanler.CheckError(err)
		defer file.Close()

		err = json.NewEncoder(file).Encode(cfg)
		errorhanler.CheckError(err)
	}
}

func LoadConfig() *Config {
	var cfg Config
	file, err := os.OpenFile(configfilepath, os.O_RDONLY, 0666)
	errorhanler.CheckError(err)
	defer file.Close()

	// decode the file information to struct
	err = json.NewDecoder(file).Decode(&cfg)
	errorhanler.CheckError(err)

	// return the config
	return &Config{
		Dir: cfg.Dir,
	}
}
