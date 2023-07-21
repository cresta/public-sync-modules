package goreleasercli

import (
	_ "embed"

	"github.com/cresta/public-sync-modules/gitignore"

	"github.com/cresta/syncer/sharedapi/drift/templatefiles"
	"github.com/cresta/syncer/sharedapi/syncer"
)

func init() {
	syncer.FxRegister(Module)
}

//go:embed .goreleaser.yaml.template
var templateStrGoReleaser string

//go:embed goreleaser.yaml.template
var templateStrActionReleaser string

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name: "goreleasercli",
	Files: map[string]string{
		".goreleaser.yaml":                  templateStrGoReleaser,
		".github/workflows/goreleaser.yaml": templateStrActionReleaser,
	},
	Priority: syncer.PriorityNormal,
	Decoder:  templatefiles.DefaultDecoder[Config](),
	Setup: &syncer.SetupMutator[gitignore.Config]{
		Name: "gitignore",
		Mutator: syncer.SimpleConfigMutator[gitignore.Config](func(cfg gitignore.Config) (gitignore.Config, error) {
			cfg.Ignores = append(cfg.Ignores, "/dist/")
			return cfg, nil
		}),
	},
})

type Config struct {
	LinkerVarPath string `yaml:"linkerVarPath"`
	MainDir       string `yaml:"mainDir"`
}
