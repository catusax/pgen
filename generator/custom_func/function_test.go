package custom_func

import (
	"fmt"
	"testing"
)

func TestWASM(t *testing.T) {
	funcs := LoadWasmFunctions([]Libs{
		{
			Path:  "/Users/coolrc/sourcecode/pgen/.template/wasm/target/wasm32-unknown-unknown/debug/wasm.wasm",
			Funcs: []string{"greet", "upper"},
		},
		{
			Path:  "/Users/coolrc/sourcecode/pgen/.template/wasm/target/wasm32-unknown-unknown/release/wasm.wasm",
			Funcs: []string{"lower"},
		},
	}, "/")

	for _, function := range funcs {
		fmt.Println(function.FuncP.(func(string) string)("abicfuwberh83hqgf8p90awjdocvfp03ihq5[t24ufew8h9svpybicdo;nc"))
	}
}
