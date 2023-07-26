package goreleasercli

import (
	_ "embed"

	"github.com/getsyncer/public-sync-modules/gitignore"

	"github.com/getsyncer/syncer-core/drift/templatefiles"
	"github.com/getsyncer/syncer-core/syncer"
)

func init() {
	syncer.FxRegister(Module)
}

//go:embed .goreleaser.yaml.template
var templateStrGoReleaser string

//go:embed goreleaser.yaml.template
var templateStrActionReleaser string

const Name = syncer.Name("goreleasercli")

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name: Name,
	Files: map[string]string{
		".goreleaser.yaml":                  templateStrGoReleaser,
		".github/workflows/goreleaser.yaml": templateStrActionReleaser,
	},
	Setup: &syncer.SetupMutator[gitignore.Config]{
		Name: gitignore.Name,
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
