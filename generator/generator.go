package generator

type Generator interface {
	// SetOptions sets template engine binding
	SetOptions(bindings map[string]any)
	// Register a new template file to template engine
	Register(tmplDir, tmplFilePath string) error

	RegisterFunc(funcName string, function any)
	// Generate all templates
	Generate() error
}

type Template struct {
	// Path is the path of output file name
	Path string
	// Template is the raw template to be parsed by template engine
	Template []byte
}
