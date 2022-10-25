package generator

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sergi/go-diff/diffmatchpatch"
)

var dmp = diffmatchpatch.New()

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

func ReadTemplate(tmplDir, tmpl string) (fileBytes []byte, filePath string, err error) {
	filePath = "."
	fileBytes, err = os.ReadFile(tmpl + ".tmpl") // read file from current directory
	if err != nil {
		// search for parent directory's .template directory
		fileBytes, filePath, err = ReadFile(filepath.Join(tmplDir, tmpl+".tmpl"))
		if err != nil {
			return nil, "", fmt.Errorf("read template file: %w", err)
		}
	}

	patch, err := ReadPatch(tmplDir, tmpl)
	if err != nil {
		return nil, "", fmt.Errorf("failed to read patch: %w ", err)
	}

	res, _ := dmp.PatchApply(patch, string(fileBytes))
	return []byte(res), filePath, nil
}

func ReadPatch(tmplDir, tmpl string) ([]diffmatchpatch.Patch, error) {
	fileBytes, err := os.ReadFile(tmpl + ".tmpl.patch")
	if err != nil {
		fileBytes, _, err = ReadFile(filepath.Join(tmplDir, tmpl+".tmpl.patch"))
		if err != nil {
			if !errors.Is(err, ErrFileNotFound) {
				return nil, fmt.Errorf("read template file: %w", err)
			} else {
				return nil, nil
			}
		}
	}

	patch, err := dmp.PatchFromText(string(fileBytes))
	if err != nil {
		return nil, fmt.Errorf("parse patch file %w", err)
	}

	return patch, nil
}
