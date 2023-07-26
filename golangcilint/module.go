package golangcilint

import (
	"context"
	_ "embed"

	"github.com/getsyncer/public-sync-modules/gosemanticrelease"

	"github.com/getsyncer/public-sync-modules/buildgo"

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
	Setup: syncer.MultiSetupSyncer([]syncer.SetupSyncer{
		&syncer.SetupMutator[buildgo.Config]{
			Name: buildgo.Name,
			Mutator: &templatefiles.GenericConfigMutator[buildgo.Config]{
				TemplateStr: updatedBuildGoLibTemplate,
				MutateFunc: func(_ context.Context, renderedTemplate string, cfg buildgo.Config) (buildgo.Config, error) {
					cfg.Jobs = append(cfg.Jobs, renderedTemplate)
					return cfg, nil
				},
			},
		},
		&syncer.SetupMutator[gosemanticrelease.Config]{
			Name: gosemanticrelease.Name,
			Mutator: &templatefiles.GenericConfigMutator[gosemanticrelease.Config]{
				TemplateStr: "",
				MutateFunc: func(_ context.Context, renderedTemplate string, cfg gosemanticrelease.Config) (gosemanticrelease.Config, error) {
					cfg.RequiredSteps = append(cfg.RequiredSteps, "lint")
					return cfg, nil
				},
			},
		},
	}),
})

type Config struct{}
