package buildgolib

import (
	_ "embed"

	"github.com/cresta/syncer/sharedapi/syncer"
	"github.com/cresta/syncer/sharedapi/templatefiles"
)

func init() {
	syncer.FxRegister(Module)
}

type Config struct {
	RunsOn string `yaml:"runs_on"`
}

//go:embed buildgolib.yaml.template
var templateStr string

var Module = templatefiles.NewModule("buildgolib", map[string]string{
	".github/workflows/buildgolib.yaml": templateStr,
}, syncer.PriorityNormal, func(runConfig syncer.RunConfig) (interface{}, error) {
	var cfg Config
	if err := runConfig.Decode(&cfg); err != nil {
		return nil, err
	}
	return cfg, nil
})
