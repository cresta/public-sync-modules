package setlicense

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/cresta/zapctx"
	"github.com/getsyncer/syncer-core/files"
	"github.com/getsyncer/syncer-core/syncer"
	"go.uber.org/fx"
)

func init() {
	syncer.FxRegister(Module)
}

//go:embed Apache-2.0.LICENSE
var apacheLicense string

type Config struct {
	License string `yaml:"license"`
}

type Syncer struct {
	logger *zapctx.Logger
}

func New(logger *zapctx.Logger) *Syncer {
	return &Syncer{
		logger: logger,
	}
}

const Name = syncer.Name("setlicense")

func (l *Syncer) Run(ctx context.Context, runData *syncer.SyncRun) (*files.System[*files.StateWithChangeReason], error) {
	var ret files.System[*files.StateWithChangeReason]
	var cfg Config
	if err := runData.RunConfig.Decode(&cfg); err != nil {
		return nil, err
	}
	if cfg.License == "" {
		l.logger.Debug(ctx, "running license logic with no license set in config")
		if err := ret.Add("LICENSE", &files.StateWithChangeReason{
			State: files.State{
				FileExistence: files.FileExistenceAbsent,
			},
			ChangeReason: &files.ChangeReason{
				Reason: "Missing license",
			},
		}); err != nil {
			return nil, fmt.Errorf("unable to add file for deletion %s: %w", "LICENSE", err)
		}
		return &ret, nil
	}
	licenseToText := map[string]string{
		"Apache-2.0": apacheLicense,
	}
	licenseText, ok := licenseToText[cfg.License]
	if !ok {
		return nil, fmt.Errorf("unknown license %s", cfg.License)
	}
	if err := ret.Add("LICENSE", &files.StateWithChangeReason{
		State: files.State{
			Mode:          0644,
			Contents:      []byte(licenseText),
			FileExistence: files.FileExistencePresent,
		},
		ChangeReason: &files.ChangeReason{
			Reason: "license text",
		},
	}); err != nil {
		return nil, fmt.Errorf("unable to add file %s: %w", "LICENSE", err)
	}
	return &ret, nil
}

func (l *Syncer) Name() syncer.Name {
	return Name
}

func (l *Syncer) Priority() syncer.Priority {
	return syncer.PriorityNormal
}

var _ syncer.DriftSyncer = &Syncer{}

var Module = fx.Module("setlicense",
	fx.Provide(
		fx.Annotate(
			New,
			fx.As(new(syncer.DriftSyncer)),
			fx.ResultTags(`group:"syncers"`),
		),
	),
)
