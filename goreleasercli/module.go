package goreleasercli

import (
	_ "embed"

	"github.com/getsyncer/syncer-core/drift/syncers/templatefiles/templatemutator"

	"github.com/getsyncer/syncer-core/fxregistry"

	"github.com/getsyncer/syncer-core/config"

	"github.com/getsyncer/public-sync-modules/gitignore"

	"github.com/getsyncer/syncer-core/drift/syncers/templatefiles"
)

func init() {
	fxregistry.Register(Module)
}

//go:embed .goreleaser.yaml.template
var templateStrGoReleaser string

//go:embed goreleaser.yaml.template
var templateStrActionReleaser string

const Name = config.Name("goreleasercli")

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name: Name,
	Files: map[string]string{
		".goreleaser.yaml":                  templateStrGoReleaser,
		".github/workflows/goreleaser.yaml": templateStrActionReleaser,
	},
	Setup: &templatemutator.SetupMutator[gitignore.Config]{
		Name: gitignore.Name,
		Mutator: templatemutator.SimpleMutator[gitignore.Config](func(cfg gitignore.Config) (gitignore.Config, error) {
			cfg.Ignores = append(cfg.Ignores, "/dist/")
			return cfg, nil
		}),
	},
})

type Config struct {
	LinkerVarPath string `yaml:"linkerVarPath"`
	MainDir       string `yaml:"mainDir"`
}
