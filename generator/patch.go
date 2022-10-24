package generator

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sergi/go-diff/diffmatchpatch"
)

type Patch struct {
	dmp   *diffmatchpatch.DiffMatchPatch
	patch map[string][]diffmatchpatch.Patch
}

func NewPatch() *Patch {
	return &Patch{
		dmp:   diffmatchpatch.New(),
		patch: make(map[string][]diffmatchpatch.Patch),
	}
}

func (p *Patch) Register(tmplDir, tmpl string) error {
	fileBytes, err := os.ReadFile(tmpl + ".patch") // read file from current directory
	if err != nil {
		// search for parent directory's .template directory
		fileBytes, _, err = ReadFile(filepath.Join(tmplDir, tmpl+".patch"))
		if err != nil {
			if !errors.Is(err, ErrFileNotFound) {
				return fmt.Errorf("read template file: %w", err)
			} else {
				return nil
			}
		}
	}

	patch, err := p.dmp.PatchFromText(string(fileBytes))
	if err != nil {
		return fmt.Errorf("parse patch file %w", err)
	}

	p.patch[tmpl] = patch

	return nil
}

func (p *Patch) Patch(tmpl string, text []byte) []byte {
	res, _ := p.dmp.PatchApply(p.patch[tmpl], string(text))
	return []byte(res)
}
