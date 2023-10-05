package buildgo

import (
	_ "embed"

	_ "github.com/getsyncer/public-sync-modules/latestversions"
	"github.com/getsyncer/syncer-core/config"
	"github.com/getsyncer/syncer-core/drift"
	"github.com/getsyncer/syncer-core/drift/syncers/templatefiles"
	"github.com/getsyncer/syncer-core/fxregistry"
)

func init() {
	fxregistry.Register(Module)
}

const Name = config.Name("buildgo")

const RunPriority = drift.PriorityLow

type Config struct {
	RunsOn                 string   `yaml:"runs_on"`
	PostTest               []string `yaml:"post_test"`
	Jobs                   []string `yaml:"jobs"`
	ActionsCheckoutVersion string   `yaml:"actions_checkout_version"`
	PrimaryBranch          string   `yaml:"primary_branch"`
	GithubRunner           string   `yaml:"github_runner"`
	SetupGoVersion         string   `yaml:"setup_go_version"`
}

//go:embed buildgo.yaml.template
var templateStr string

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name:     Name,
	Priority: RunPriority,
	Files: map[string]string{
		".github/workflows/buildgo.yaml": templateStr,
	},
})
