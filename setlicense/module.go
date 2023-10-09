package setlicense

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/getsyncer/syncer-core/syncer/planner"

	"github.com/getsyncer/syncer-core/drift"

	"github.com/getsyncer/syncer-core/fxregistry"

	"github.com/getsyncer/syncer-core/config"

	"github.com/cresta/zapctx"
	"github.com/getsyncer/syncer-core/files"
	"go.uber.org/fx"
)

func init() {
	fxregistry.Register(Module)
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

const Name = config.Name("setlicense")

func (l *Syncer) DetectDrift(ctx context.Context, runData *drift.RunData) (*files.System[*files.StateWithChangeReason], error) {
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

func (l *Syncer) Name() config.Name {
	return Name
}

func (l *Syncer) Priority() drift.Priority {
	return drift.PriorityNormal
}

var _ drift.Detector = &Syncer{}

var Module = fx.Module("setlicense",
	fx.Provide(
		fx.Annotate(
			New,
			fx.As(new(drift.Detector)),
			fx.ResultTags(`group:"syncers"`),
		),
	),
	planner.FxOption(planner.WithFilesAllowedNoMagicString("LICENSE")),
)
