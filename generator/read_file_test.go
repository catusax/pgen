package generator

import (
	"log"
	"testing"
)

func TestReadFile(t *testing.T) {
	path, _, err := ReadFile("generator/options.go")
	if err != nil {
		t.Error(err)
	}
	log.Println(path)
}
