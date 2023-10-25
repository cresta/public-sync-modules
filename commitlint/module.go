package commitlint

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

//go:embed commitlint.yaml.template
var commitLintTemplate string

const Name = config.Name("commitlint")

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name: Name,
	Files: map[string]string{
		".github/workflows/commitlint.yaml": commitLintTemplate,
	},
})

type Config struct {
	ActionsCheckoutVersion string `yaml:"actions_checkout_version"`
	GithubRunner           string `yaml:"github_runner"`
	CommitLintVersion      string `yaml:"commit_lint_version"`
}
