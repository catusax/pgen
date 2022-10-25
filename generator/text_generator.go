package generator

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

type TextGeneraotr struct {
	bindings  map[string]any
	templates []TextTemplate
	funcs     map[string]any
}

type TextTemplate struct {
	path     string
	template *template.Template
}

func NewTextGenerator() Generator {
	return &TextGeneraotr{
		bindings: make(map[string]any),
		funcs:    make(map[string]any),
	}
}

func (l *TextGeneraotr) SetOptions(bindings map[string]any) {
	l.bindings = bindings
}

func (l *TextGeneraotr) Register(tmplDir, tmpl string) error {
	fileBytes, _, err := ReadTemplate(tmplDir, tmpl)
	if err != nil {
		return fmt.Errorf("read template file: %w", err)
	}

	_template, err := template.New("").Funcs(l.funcs).Parse(string(fileBytes))
	if err != nil {
		return fmt.Errorf("parse template %w", err)
	}

	l.templates = append(l.templates, TextTemplate{
		path:     tmpl,
		template: _template,
	})

	return nil
}

func (l *TextGeneraotr) RegisterFunc(funcName string, function any) {
	l.funcs[funcName] = function
}

func (l *TextGeneraotr) Generate() error {
	for _, _template := range l.templates {
		var filename bytes.Buffer

		err := template.Must(template.New("").Funcs(l.funcs).Parse(_template.path)).Execute(&filename, l.bindings)
		if err != nil {
			return fmt.Errorf("error rendering template for filename %s : %w",
				_template.path, err)
		}

		err = os.MkdirAll(filepath.Dir(filename.String()), 0o777)
		if err != nil {
			return fmt.Errorf("error mkdir %w", err)
		}

		fmt.Println(filename.String())

		var out bytes.Buffer

		err = _template.template.Funcs(l.funcs).Execute(&out, l.bindings)
		if err != nil {
			return fmt.Errorf("error rendering template %s : %w",
				_template.path, err)
		}

		ioutil.WriteFile(filename.String(), out.Bytes(), 0o644)
	}

	return nil
}
