package main

import (
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"

	"github.com/amsokol/protoc-gen-gotag/module"
)

func main() {
	pgs.Init(pgs.DebugEnv("DEBUG")).
		RegisterModule(module.New()).
		RegisterPostProcessor(pgsgo.GoFmt()).
		Render()
}
