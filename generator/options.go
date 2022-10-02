package generator

import (
	"os"
	"strings"
	"sync"
)

var envs []string

func init() {
	envs = os.Environ()
}

type Bindings map[string]any

// LoadFromFile load Makefile variables from Makefile and replace existing variables
func (o Bindings) LoadFromFile() Bindings {
	envs := ReadMakefileENV("Makefile")
	for k, v := range envs {
		o[k] = v
	}
	return o
}

func (o Bindings) LoadFromENV() Bindings {
	envOnce.Do(copyenv)
	for name, i := range env {
		// code from syscall.Getenv
		if strings.HasPrefix(name, "PGEN_") {
			s := envs[i]
			for i := 0; i < len(s); i++ {
				if s[i] == '=' {
					o[name[5:]] = s[i+1:]
				}
			}
		}
	}

	return o
}

func (o Bindings) Set(key string, value any) Bindings {
	o[key] = value

	return o
}

var (
	env     = make(map[string]int)
	envOnce sync.Once
)

func copyenv() {
	env = make(map[string]int)
	for i, s := range envs {
		for j := 0; j < len(s); j++ {
			if s[j] == '=' {
				key := s[:j]
				if _, ok := env[key]; !ok {
					env[key] = i // first mention of key
				} else {
					// Clear duplicate keys. This permits Unsetenv to
					// safely delete only the first item without
					// worrying about unshadowing a later one,
					// which might be a security problem.
					envs[i] = ""
				}

				break
			}
		}
	}
}
