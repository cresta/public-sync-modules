package gosemanticrelease

import (
	"context"
	_ "embed"

	"github.com/getsyncer/public-sync-modules/buildgo"

	"github.com/getsyncer/syncer-core/drift/templatefiles"
	"github.com/getsyncer/syncer-core/syncer"
)

func init() {
	syncer.FxRegister(Module)
}

//go:embed bump_tag_step.yaml.template
var templateStr string

const Name = syncer.Name("gosemanticrelease")

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name: Name,
	Files: map[string]string{
		// Note: Empty string filename is removed by PostGenProcessor
		"": templateStr,
	},
	Priority: buildgo.RunPriority + 1, // Force it to run before buildgo so our mutation is rendered.
	PostGenProcessor: &templatefiles.PostGenConfigMutator[buildgo.Config]{
		ToMutate: buildgo.Name,
		PostGenMutatorFunc: func(_ context.Context, renderedTemplate string, cfg buildgo.Config) (buildgo.Config, error) {
			cfg.Jobs = append(cfg.Jobs, renderedTemplate)
			return cfg, nil
		},
	},
})

type Config struct {
	RequiredSteps []string `yaml:"required_steps"`
}

func (c Config) AllRequiredSteps() []string {
	ret := make([]string, 0, len(c.RequiredSteps))
	for _, step := range c.RequiredSteps {
		ret = append(ret, step)
	}
	ret = append(ret, "build", "test")
	return ret
}
