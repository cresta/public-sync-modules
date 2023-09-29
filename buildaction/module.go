package buildaction

import (
	_ "embed"

	"github.com/getsyncer/syncer-core/drift/syncers/templatefiles/templatemutator"

	"github.com/getsyncer/syncer-core/drift"

	"github.com/getsyncer/syncer-core/fxregistry"

	"github.com/getsyncer/syncer-core/config"

	"github.com/getsyncer/syncer-core/drift/syncers/templatefiles"
)

func init() {
	fxregistry.Register(Module)
}

type UpdatedVersion string

// renovate: datasource=github-tags depName=actions/checkout versioning=loose
const actionsCheckout = "v3"

const Name = config.Name("buildaction")
const RunPriority = drift.PriorityNormal

type Config struct {
	RunsOn                 string   `yaml:"runs_on"`
	PostTest               []string `yaml:"post_test"`
	Jobs                   []string `yaml:"jobs"`
	ActionsCheckoutVersion string
}

//go:embed buildgithubaction.yaml.template
var templateStr string

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name:     Name,
	Priority: RunPriority,
	Setup:    templatemutator.SimpleTemplateSetupMutator[Config](Name, setDefaults),
	Files: map[string]string{
		".github/workflows/buildgithubaction.yaml": templateStr,
	},
})

func setDefaults(cfg Config) Config {
	if cfg.ActionsCheckoutVersion == "" {
		cfg.ActionsCheckoutVersion = actionsCheckout
	}
	return cfg
}
