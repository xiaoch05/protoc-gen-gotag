package module

import (
	"go/parser"
	"go/printer"
	"go/token"
	"strings"
	"fmt"

	"github.com/fatih/structtag"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

type mod struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
}

func New() pgs.Module {
	return &mod{ModuleBase: &pgs.ModuleBase{}}
}

func (m *mod) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.ctx = pgsgo.InitContext(c.Parameters())
}

func (m *mod) Name() string {
	return "gotag"
}

func (m *mod) Execute(targets map[string]pgs.File, packages map[string]pgs.Package) []pgs.Artifact {
	fmt.Println("Execute xxx process")
	xtv := m.Parameters().Str("xxx")
	xtv = strings.Replace(xtv, "+", ":", -1)
	xt, err := structtag.Parse(xtv)
	m.CheckErr(err)

	extractor := newTagExtractor(m)
	for _, f := range targets {
		tags := extractor.Extract(f)

		tags.AddTagsToXXXFields(xt)

		gfname := m.BuildContext.JoinPath(f.InputPath().SetExt(".pb.go").Base())

		fs := token.NewFileSet()
		fn, err := parser.ParseFile(fs, gfname, nil, parser.ParseComments)
		m.CheckErr(err)

		m.CheckErr(Retag(fn, tags))

		var buf strings.Builder
		m.CheckErr(printer.Fprint(&buf, fs, fn))

		m.OverwriteGeneratorFile(gfname, buf.String())
	}

	return m.Artifacts()
}
