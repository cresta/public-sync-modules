package golangcilint

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/cresta/public-sync-modules/buildgolib"
	"github.com/cresta/syncer/sharedapi/drift/templatefiles"
	"github.com/cresta/syncer/sharedapi/syncer"
)

func init() {
	syncer.FxRegister(Module)
}

//go:embed .golangci.yaml.template
var templateStrGolangCi string

//go:embed updatedbuildgolib.yaml.template
var updatedBuildGoLibTemplate string

type UpdateGoBuild struct{}

func (t *UpdateGoBuild) Mutate(ctx context.Context, runData *syncer.SyncRun, cfg buildgolib.Config) (buildgolib.Config, error) {
	updatedBuildGoLib, err := templatefiles.NewTemplate("updatedbuildgolib", updatedBuildGoLibTemplate)
	if err != nil {
		return cfg, fmt.Errorf("unable to parse updatedbuildgolib template: %w", err)
	}
	res, err := templatefiles.ExecuteTemplateOnConfig(ctx, runData, cfg, updatedBuildGoLib)
	if err != nil {
		return cfg, fmt.Errorf("unable to execute template: %w", err)
	}
	cfg.PostTest = append(cfg.PostTest, res)
	return cfg, nil
}

type MutatorSetup struct {
}

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name: "golangcilint",
	Files: map[string]string{
		".golangci.yml": templateStrGolangCi,
	},
	Priority: syncer.PriorityNormal,
	Decoder:  templatefiles.DefaultDecoder[Config](),
	Setup: syncer.SetupSyncerFunc(func(ctx context.Context, runData *syncer.SyncRun) error {
		if err := syncer.AddMutator[buildgolib.Config](runData.Registry, "buildgolib", &UpdateGoBuild{}); err != nil {
			return fmt.Errorf("unable to add mutator: %w", err)
		}
		return nil
	}),
})

type Config struct{}
