package buildgolib

import (
	_ "embed"

	"github.com/getsyncer/syncer/sharedapi/drift/templatefiles"
	"github.com/getsyncer/syncer/sharedapi/syncer"
)

func init() {
	syncer.FxRegister(Module)
}

const Name = syncer.Name("buildgolib")

type Config struct {
	RunsOn   string   `yaml:"runs_on"`
	PostTest []string `yaml:"post_test"`
	Jobs     []string `yaml:"jobs"`
}

//go:embed buildgolib.yaml.template
var templateStr string

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name: Name,
	Files: map[string]string{
		".github/workflows/buildgolib.yaml": templateStr,
	},
})
