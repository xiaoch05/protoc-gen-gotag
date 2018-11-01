package main

import (
	"github.com/lyft/protoc-gen-star"
	"github.com/amsokol/protoc-gen-gotag/module"
)

func main() {
	pgs.Init().RegisterModule(module.New()).RegisterPostProcessor(pgs.GoFmt()).Render()
}
