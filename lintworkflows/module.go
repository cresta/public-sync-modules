package lintworkflows

import (
	_ "embed"

	"github.com/getsyncer/syncer-core/fxregistry"

	"github.com/getsyncer/syncer-core/config"

	"github.com/getsyncer/syncer-core/drift/syncers/templatefiles"
)

func init() {
	fxregistry.Register(Module)
}

type Config struct {
}

//go:embed lintworkflows.yaml.template
var templateStr string

const Name = config.Name("lintworkflows")

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name: Name,
	Files: map[string]string{
		".github/workflows/lintworkflows.yaml": templateStr,
	},
})
