package golangcilint

import (
	"context"
	_ "embed"

	"github.com/getsyncer/public-sync-modules/buildgolib"
	"github.com/getsyncer/syncer/sharedapi/drift/templatefiles"
	"github.com/getsyncer/syncer/sharedapi/syncer"
)

func init() {
	syncer.FxRegister(Module)
}

//go:embed .golangci.yaml.template
var templateStrGolangCi string

//go:embed updatedbuildgolib.yaml.template
var updatedBuildGoLibTemplate string

const Name = syncer.Name("golangcilint")

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name: Name,
	Files: map[string]string{
		".golangci.yml": templateStrGolangCi,
	},
	Setup: &syncer.SetupMutator[buildgolib.Config]{
		Name: buildgolib.Name,
		Mutator: &templatefiles.GenericConfigMutator[buildgolib.Config]{
			TemplateStr: updatedBuildGoLibTemplate,
			MutateFunc: func(_ context.Context, renderedTemplate string, cfg buildgolib.Config) (buildgolib.Config, error) {
				cfg.Jobs = append(cfg.Jobs, renderedTemplate)
				return cfg, nil
			},
		},
	},
})

type Config struct{}
