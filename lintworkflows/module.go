package lintworkflows

import (
	_ "embed"
	"github.com/getsyncer/syncer/sharedapi/drift/templatefiles"
	"github.com/getsyncer/syncer/sharedapi/syncer"
)

func init() {
	syncer.FxRegister(Module)
}

type Config struct {
}

//go:embed lintworkflows.yaml.template
var templateStr string

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name: "lintworkflows",
	Files: map[string]string{
		".github/workflows/lintworkflows.yaml": templateStr,
	},
	Priority: syncer.PriorityNormal,
	Decoder:  templatefiles.DefaultDecoder[Config](),
})
