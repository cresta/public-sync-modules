package autoapprove

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

//go:embed autoapprove.yaml.template
var templateStr string

var Module = templatefiles.NewModule("autoapprove", map[string]string{
	".github/workflows/autoapprove.yaml": templateStr,
}, syncer.PriorityNormal, func(runConfig syncer.RunConfig) (interface{}, error) {
	var cfg Config
	if err := runConfig.Decode(&cfg); err != nil {
		return nil, err
	}
	return cfg, nil
})
