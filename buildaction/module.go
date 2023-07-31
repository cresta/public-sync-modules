package buildaction

import (
	_ "embed"

	"github.com/getsyncer/syncer-core/drift"

	"github.com/getsyncer/syncer-core/fxregistry"

	"github.com/getsyncer/syncer-core/config"

	"github.com/getsyncer/syncer-core/drift/syncers/templatefiles"
)

func init() {
	fxregistry.Register(Module)
}

const Name = config.Name("buildaction")
const RunPriority = drift.PriorityNormal

type Config struct {
	RunsOn   string   `yaml:"runs_on"`
	PostTest []string `yaml:"post_test"`
	Jobs     []string `yaml:"jobs"`
}

//go:embed buildgithubaction.yaml.template
var templateStr string

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name:     Name,
	Priority: RunPriority,
	Files: map[string]string{
		".github/workflows/buildgithubaction.yaml": templateStr,
	},
})
