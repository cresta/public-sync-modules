package renovatebot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/getsyncer/syncer-core/syncer/planner"

	"github.com/getsyncer/syncer-core/files"

	"github.com/getsyncer/syncer-core/drift"

	"github.com/getsyncer/syncer-core/drift/syncers/staticfile"

	// To make sure we get defaults of the latest versions of actions
	_ "github.com/getsyncer/public-sync-modules/latestversions"
	"github.com/getsyncer/syncer-core/config"
	"github.com/getsyncer/syncer-core/fxregistry"
)

func init() {
	fxregistry.Register(Module)
}

const Name = config.Name("renovatebot")

type Config struct {
	Schema      string   `yaml:"$schema" json:"$schema"`
	IgnorePaths []string `yaml:"ignorePaths,omitempty" json:"ignorePaths,omitempty"`
	GeneratedBy string   `yaml:"generatedBy,omitempty" json:"generatedBy,omitempty"`
}

func (c Config) Changes(ctx context.Context) (files.System[*files.StateWithChangeReason], error) {
	var ret files.System[*files.StateWithChangeReason]
	if c.Schema == "" {
		c.Schema = "https://docs.renovatebot.com/renovate-schema.json"
	}
	c.GeneratedBy = drift.MagicTrackedString
	seenChanges := planner.GetCurrentChanges(ctx)
	for _, change := range seenChanges {
		for _, path := range change.Paths() {
			toChange := change.Get(path)
			if toChange.State.FileExistence == files.FileExistenceAbsent {
				continue
			}
			c.IgnorePaths = append(c.IgnorePaths, path.String())
		}
	}
	sort.Strings(c.IgnorePaths)
	var content bytes.Buffer
	enc := json.NewEncoder(&content)
	enc.SetIndent("", "\t")
	if err := enc.Encode(c); err != nil {
		return ret, fmt.Errorf("failed to encode json: %w", err)
	}
	if err := ret.Add(".renovate-autogen.json", &files.StateWithChangeReason{
		ChangeReason: &files.ChangeReason{
			Reason: "renovatebot",
		},
		State: files.State{
			Mode:          0644,
			Contents:      content.Bytes(),
			FileExistence: files.FileExistencePresent,
		},
	}); err != nil {
		return ret, fmt.Errorf("failed to add file: %w", err)
	}
	return ret, nil
}

var Module = staticfile.NewCustomModule[Config](Name, drift.Priority(0))
