package generator

import (
	"errors"
	"os"
)

// ReadFile recursively reads a file for given path till user home directory
// and returns nearest file to given path
// given path must be relative to working directory
func ReadFile(filename string) (fileBytes []byte, filePath string, err error) {
	currentdir, _ := os.Getwd()
	home, _ := os.UserHomeDir()
	defer os.Chdir(currentdir)

	for {
		dat, err := os.ReadFile(filename)

		wd, _ := os.Getwd()

		if err == nil {
			return dat, wd, nil
		}

		if dir, _ := os.Getwd(); dir == home {
			break
		}

		os.Chdir("..")
	}

	return nil, "", ErrFileNotFound
}

var ErrFileNotFound = errors.New("file not found")
