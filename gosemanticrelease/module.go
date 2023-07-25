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

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name:     "gosemanticrelease",
	Priority: syncer.PriorityNormal,
	Decoder:  templatefiles.DefaultDecoder[Config](),
	Setup: &syncer.SetupMutator[buildgolib.Config]{
		Name: "buildgolib",
		Mutator: &templatefiles.GenericConfigMutator[buildgolib.Config]{
			Name:        "add semantic release bump tag step",
			TemplateStr: templateStr,
			MutateFunc: func(_ context.Context, renderedTemplate string, cfg buildgolib.Config) (buildgolib.Config, error) {
				cfg.PostTest = append(cfg.PostTest, renderedTemplate)
				return cfg, nil
			},
		},
	},
})

type Config struct {
	LinkerVarPath string `yaml:"linkerVarPath"`
	MainDir       string `yaml:"mainDir"`
}
