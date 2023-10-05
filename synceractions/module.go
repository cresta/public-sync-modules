package synceractions

import (
	_ "embed"

	_ "github.com/getsyncer/public-sync-modules/latestversions"
	"github.com/getsyncer/syncer-core/config"
	"github.com/getsyncer/syncer-core/drift/syncers/templatefiles"
	"github.com/getsyncer/syncer-core/fxregistry"
)

func init() {
	fxregistry.Register(Module)
}

type Config struct {
	PrimaryBranch          string `yaml:"primaryBranch"`
	GithubRunner           string `yaml:"github_runner"`
	ActionsCheckoutVersion string `yaml:"actions_checkout_version"`
}

//go:embed watchsynccomment.yaml.template
var watchsynccommentTemplateStr string

//go:embed checksyncer.yaml.template
var checksyncerTemplateStr string

const Name = config.Name("synceractions")

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name: Name,
	Files: map[string]string{
		".github/workflows/watchsynccomment.yaml": watchsynccommentTemplateStr,
		".github/workflows/checksync.yaml":        checksyncerTemplateStr,
	},
})
