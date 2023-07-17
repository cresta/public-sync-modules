package buildgolib

import (
	_ "embed"

	"github.com/cresta/syncer/sharedapi/drift/templatefiles"
	"github.com/cresta/syncer/sharedapi/syncer"
)

func init() {
	syncer.FxRegister(Module)
}

type Config struct {
	RunsOn   string   `yaml:"runs_on"`
	PostTest []string `yaml:"post_test"`
}

//go:embed buildgolib.yaml.template
var templateStr string

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name: "buildgolib",
	Files: map[string]string{
		".github/workflows/buildgolib.yaml": templateStr,
	},
	Priority: syncer.PriorityNormal,
	Decoder: func(runConfig syncer.RunConfig) (Config, error) {
		var cfg Config
		if err := runConfig.Decode(&cfg); err != nil {
			return cfg, err
		}
		return cfg, nil
	},
})
