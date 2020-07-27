package main

import (
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"fmt"

	"github.com/amsokol/protoc-gen-gotag/module"
)

func main() {
	fmt.Println("start protoc-gen-gotag...")
	pgs.Init(pgs.DebugEnv("DEBUG")).
		RegisterModule(module.New()).
		RegisterPostProcessor(pgsgo.GoFmt()).
		Render()
}
