package gosemanticrelease

import (
	"context"
	_ "embed"
	"sort"

	"github.com/getsyncer/public-sync-modules/buildaction"
	"github.com/getsyncer/public-sync-modules/buildgo"

	// To make sure we get defaults of the latest versions of actions
	_ "github.com/getsyncer/public-sync-modules/latestversions"
	"github.com/getsyncer/syncer-core/config"
	"github.com/getsyncer/syncer-core/drift/syncers/templatefiles"
	"github.com/getsyncer/syncer-core/drift/syncers/templatefiles/templatemutator"
	"github.com/getsyncer/syncer-core/fxregistry"
)

func init() {
	fxregistry.Register(Module)
}

//go:embed bump_tag_step.yaml.template
var templateStr string

const Name = config.Name("gosemanticrelease")

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name: Name,
	Files: map[string]string{
		// Note: Empty string filename is removed by PostGenProcessor
		"_go_": templateStr,
		"_gh_": templateStr,
	},
	Priority: buildgo.RunPriority + 1, // Force it to run before buildgo so our mutation is rendered.
	PostGenProcessor: templatefiles.PostGenProcessorList{
		&templatemutator.PostGenConfigMutator[buildgo.Config]{
			ToMutate:     buildgo.Name,
			TemplateName: "_go_",
			PostGenMutatorFunc: func(_ context.Context, renderedTemplate string, cfg buildgo.Config) (buildgo.Config, error) {
				cfg.Jobs = append(cfg.Jobs, renderedTemplate)
				return cfg, nil
			},
		},
		&templatemutator.PostGenConfigMutator[buildaction.Config]{
			ToMutate:     buildaction.Name,
			TemplateName: "_gh_",
			PostGenMutatorFunc: func(_ context.Context, renderedTemplate string, cfg buildaction.Config) (buildaction.Config, error) {
				cfg.Jobs = append(cfg.Jobs, renderedTemplate)
				return cfg, nil
			},
		},
	},
})

type Config struct {
	RequiredSteps                  []string `yaml:"required_steps"`
	GithubRunner                   string   `yaml:"github_runner"`
	ActionsCheckoutVersion         string   `yaml:"actions_checkout_version"`
	GithubAppTokenAction           string   `yaml:"github_app_token_action"`
	GoSemanticReleaseActionVersion string   `yaml:"go_semantic_release_action_version"`
}

func (c Config) AllRequiredSteps() []string {
	ret := make([]string, 0, len(c.RequiredSteps))
	ret = append(ret, "build", "test")
	ret = append(ret, c.RequiredSteps...)
	ret = removeDuplicate(ret)
	sort.Strings(ret)
	return ret
}

func removeDuplicate[T comparable](items []T) []T {
	ret := make([]T, 0, len(items))
	seen := make(map[T]struct{})
	for _, item := range items {
		if _, ok := seen[item]; !ok {
			ret = append(ret, item)
			seen[item] = struct{}{}
		}
	}
	return ret
}
