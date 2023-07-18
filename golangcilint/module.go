package golangcilint

import (
	"context"
	_ "embed"
	"github.com/cresta/public-sync-modules/buildgolib"
	"github.com/cresta/syncer/sharedapi/drift/templatefiles"
	"github.com/cresta/syncer/sharedapi/syncer"
)

func init() {
	syncer.FxRegister(Module)
}

//go:embed .golangci.yaml.template
var templateStrGolangCi string

//go:embed updatedbuildgolib.yaml.template
var updatedBuildGoLibTemplate string

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name: "golangcilint",
	Files: map[string]string{
		".golangci.yml": templateStrGolangCi,
	},
	Priority: syncer.PriorityNormal,
	Decoder:  templatefiles.DefaultDecoder[Config](),
	Setup: &syncer.SetupMutator[buildgolib.Config]{
		Name: "buildgolib",
		Mutator: &templatefiles.GenericConfigMutator[buildgolib.Config]{
			Name:        "buildgolib",
			TemplateStr: updatedBuildGoLibTemplate,
			MutateFunc: func(_ context.Context, renderedTemplate string, cfg buildgolib.Config) (buildgolib.Config, error) {
				cfg.PostTest = append(cfg.PostTest, renderedTemplate)
				return cfg, nil
			},
		},
	},
})

type Config struct{}
