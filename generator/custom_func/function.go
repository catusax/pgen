package custom_func

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/bytecodealliance/wasmtime-go"
)

var engine *wasmtime.Engine

func init() {
	engine = wasmtime.NewEngine()
}

type Libs struct {
	Path  string   `yaml:"path"`
	Funcs []string `yaml:"funcs"`
}

type Func struct {
	Name  string
	FuncP any
}

func LoadWasmFunctions(libs []Libs, basePath string) []Func {
	var functions []Func
	for _, lib := range libs {
		store := wasmtime.NewStore(engine)
		store.SetWasi(wasmtime.NewWasiConfig())

		module, err := wasmtime.NewModuleFromFile(store.Engine, filepath.Join(basePath, lib.Path))
		if err != nil {
			fmt.Println("failed to load library ", err.Error())

			continue
		}

		linker := wasmtime.NewLinker(store.Engine)

		linker.DefineWasi()

		instance, _ := linker.Instantiate(store, module)

		// instance, err := wasmtime.NewInstance(store, module, nil)
		// if err != nil {
		// 	panic(err)
		// }

		for _, FuncName := range lib.Funcs {
			FuncName := FuncName

			funcP := func(input string) string {
				memory := instance.GetExport(store, "memory").Memory()

				inputLen := int32(len(input))
				inputP, err := instance.GetExport(store, "allocate").Func().Call(store, inputLen)
				if err != nil {
					panic(err)
				}

				inputMem := memory.UnsafeData(store)[inputP.(int32):]

				for i := 0; i < len([]byte(input)); i++ {
					inputMem[i] = []byte(input)[i]
				}

				res, err := instance.GetExport(store, FuncName).Func().Call(store, inputP)
				if err != nil {
					panic(err)
				}

				outputP := res.(int32)

				outmem := memory.UnsafeData(store)[outputP:]
				outputLen := 0
				var output strings.Builder
				for {
					if outmem[outputLen] == 0 {
						break
					}

					output.WriteByte(outmem[outputLen])
					outputLen++
				}

				dealloc := instance.GetExport(store, "deallocate").Func()

				dealloc.Call(store, inputP, inputLen)
				dealloc.Call(store, outputP, outputLen)

				return output.String()
			}

			functions = append(functions, Func{
				Name:  FuncName,
				FuncP: funcP,
			},
			)
		}
	}

	return functions
}
