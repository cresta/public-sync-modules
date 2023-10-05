package buildaction

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

const Name = config.Name("buildaction")

type Config struct {
	RunsOn                 string   `yaml:"runs_on"`
	PostTest               []string `yaml:"post_test"`
	Jobs                   []string `yaml:"jobs"`
	ActionsCheckoutVersion string   `yaml:"actions_checkout_version"`
	PrimaryBranch          string   `yaml:"primary_branch"`
	GithubRunner           string   `yaml:"github_runner"`
}

//go:embed buildgithubaction.yaml.template
var templateStr string

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name: Name,
	Files: map[string]string{
		".github/workflows/buildgithubaction.yaml": templateStr,
	},
})
