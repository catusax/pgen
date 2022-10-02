package generator

import (
	"errors"
	"os"
)

// ReadFile recursively reads a file for given path till user home directory
// and returns nearest file to given path
// given path must be relative to working directory
func ReadFile(filename string) (filePath string, fileBytes []byte, err error) {
	currentdir, _ := os.Getwd()
	home, _ := os.UserHomeDir()

	for {
		dat, err := os.ReadFile(filename)
		if err == nil {
			return filePath, dat, nil
		}

		if dir, _ := os.Getwd(); dir == home {
			break
		}

		os.Chdir("..")
	}

	os.Chdir(currentdir)

	return "", nil, errors.New("file not found " + filename)
}
