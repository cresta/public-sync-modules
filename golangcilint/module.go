package golangcilint

import (
	_ "embed"
	"github.com/cresta/syncer/sharedapi/drift/templatefiles"
	"github.com/cresta/syncer/sharedapi/syncer"
)

func init() {
	syncer.FxRegister(Module)
}

//go:embed .golangci.yaml.template
var templateStrGolangCi string

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name: "golangcilint",
	Files: map[string]string{
		".golangci.yml": templateStrGolangCi,
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

type Config struct{}
