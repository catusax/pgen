package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/osteele/liquid"
)

type LiquidGenerator struct {
	bindings  map[string]any
	engine    *liquid.Engine
	templates []liquidTemplate
	patch     *Patch
}

type liquidTemplate struct {
	path     string
	template *liquid.Template
}

func NewLiquidGenerator() Generator {
	return &LiquidGenerator{
		bindings: make(map[string]any),
		engine:   liquid.NewEngine(),
		patch:    NewPatch(),
	}
}

func (l *LiquidGenerator) SetOptions(bindings map[string]any) {
	l.bindings = bindings
}

func (l *LiquidGenerator) RegisterFunc(funcName string, function any) {
	l.engine.RegisterFilter(funcName, function)
}

func (l *LiquidGenerator) Register(tmplDir, tmpl string) error {
	fileBytes, err := os.ReadFile(tmpl + ".tmpl") // read file from current directory
	if err != nil {
		// search for parent directory's .template directory
		fileBytes, _, err = ReadFile(filepath.Join(tmplDir, tmpl+".tmpl"))
		if err != nil {
			return fmt.Errorf("read template file: %w", err)
		}
	}

	template, err := l.engine.ParseTemplate(fileBytes)
	if err != nil {
		return fmt.Errorf("failed to parse template: %s,%w ", tmpl, err)
	}

	err = l.patch.Register(tmplDir, tmpl)
	if err != nil {
		return fmt.Errorf("failed to Register patch: %w ", err)
	}

	l.templates = append(l.templates, liquidTemplate{
		path:     tmpl,
		template: template,
	})

	return nil
}

func (l *LiquidGenerator) Generate() error {
	for _, template := range l.templates {
		out, sourceErr := template.template.Render(l.bindings)
		if sourceErr != nil {
			return fmt.Errorf("error rendering template %s at line %d: %w",
				template.path, sourceErr.LineNumber(), sourceErr.Cause())
		}

		out = l.patch.Patch(template.path, out)

		outPath, sourceErr := l.engine.ParseAndRenderString(template.path, l.bindings)
		if sourceErr != nil {
			return fmt.Errorf("error rendering template path %s at line %d: %w",
				template.path, sourceErr.LineNumber(), sourceErr.Cause())
		}

		err := os.MkdirAll(filepath.Dir(outPath), os.ModeDir)
		if err != nil {
			return fmt.Errorf("error mkdir %w", err)
		}

		err = os.WriteFile(outPath, out, 0o644)
		if err != nil {
			return fmt.Errorf("error write file %w", err)
		}
	}

	return nil
}
