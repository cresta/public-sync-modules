package gosemanticrelease

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

//go:embed bump_tag_step.yaml.template
var templateStr string

const Name = syncer.Name("gosemanticrelease")

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name: Name,
	Setup: &syncer.SetupMutator[buildgolib.Config]{
		Name: buildgolib.Name,
		Mutator: &templatefiles.GenericConfigMutator[buildgolib.Config]{
			TemplateStr: templateStr,
			MutateFunc: func(_ context.Context, renderedTemplate string, cfg buildgolib.Config) (buildgolib.Config, error) {
				cfg.Jobs = append(cfg.Jobs, renderedTemplate)
				return cfg, nil
			},
		},
	},
})

type Config struct {
	LinkerVarPath string `yaml:"linkerVarPath"`
	MainDir       string `yaml:"mainDir"`
}
