package buildgo

import (
	_ "embed"

	"github.com/getsyncer/syncer-core/drift/templatefiles"
	"github.com/getsyncer/syncer-core/syncer"
)

func init() {
	syncer.FxRegister(Module)
}

const Name = syncer.Name("buildgo")
const RunPriority = syncer.PriorityNormal

type Config struct {
	RunsOn   string   `yaml:"runs_on"`
	PostTest []string `yaml:"post_test"`
	Jobs     []string `yaml:"jobs"`
}

//go:embed buildgo.yaml.template
var templateStr string

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name:     Name,
	Priority: RunPriority,
	Files: map[string]string{
		".github/workflows/buildgo.yaml": templateStr,
	},
})
