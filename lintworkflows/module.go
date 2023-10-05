package lintworkflows

import (
	_ "embed"

	// To make sure we get defaults of the latest versions of actions
	_ "github.com/getsyncer/public-sync-modules/latestversions"
	"github.com/getsyncer/syncer-core/config"
	"github.com/getsyncer/syncer-core/drift/syncers/templatefiles"
	"github.com/getsyncer/syncer-core/fxregistry"
)

func init() {
	fxregistry.Register(Module)
}

type Config struct {
	PrimaryBranch                    string `yaml:"primary_branch"`
	ActionsCheckoutVersion           string `yaml:"actions_checkout_version"`
	ReviewdogActionActionlintVersion string `yaml:"reviewdog_action_actionlint_version"`
	GithubRunner                     string `yaml:"github_runner"`
}

//go:embed lintworkflows.yaml.template
var templateStr string

const Name = config.Name("lintworkflows")

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name: Name,
	Files: map[string]string{
		".github/workflows/lintworkflows.yaml": templateStr,
	},
})
