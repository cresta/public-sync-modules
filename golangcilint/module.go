package golangcilint

import (
	"context"
	_ "embed"
	"github.com/cresta/public-sync-modules/buildgolib"
	"github.com/cresta/syncer/sharedapi/drift/templatefiles"
	"github.com/cresta/syncer/sharedapi/syncer"
)

func init() {
	syncer.FxRegister(Module)
}

//go:embed .golangci.yaml.template
var templateStrGolangCi string

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name: "golangcilint",
	Files: map[string]string{
		".golangci.yml": templateStrGolangCi,
	},
	Priority: syncer.PriorityNormal,
	Decoder: func(runConfig syncer.RunConfig) (Config, error) {
		var cfg Config
		if err := runConfig.Decode(&cfg); err != nil {
			return cfg, err
		}
		return cfg, nil
	},
	Setup: syncer.SetupSyncerFunc(func(ctx context.Context, runData *syncer.SyncRun) error {
		goSyncerIface, exists := runData.Registry.Get("buildgolib")
		if !exists {
			return nil
		}
		goSyncer := goSyncerIface.(*templatefiles.Generator[buildgolib.Config])
		goSyncer.AddMutator(func(cfg buildgolib.Config) buildgolib.Config {
			cfg.PostTest = append(cfg.PostTest, "golangci-lint run")
			return cfg
		})
		return nil
	}),
})

type Config struct{}
