package goreleasercli

import (
	"testing"

	"github.com/getsyncer/syncer-core/drifttest"
	"github.com/getsyncer/syncer-core/files"
)

func TestModule(t *testing.T) {
	config := `
version: 1
logic:
  - source: github.com/getsyncer/syncer-core/drift/syncers/goreleasercli
syncs:
  - logic: goreleasercli
`
	t.Run("update-fresh-file", drifttest.WithRun(config, &files.System[*files.State]{}, func(t *testing.T, items *drifttest.Items) {
		items.TestRun.MustExitCode(t, 0)
		drifttest.FileContains(t, ".goreleaser.yaml", "CGO_ENABLED=0")
		drifttest.FileContains(t, ".github/workflows/goreleaser.yaml", "Release binaries using GoReleaser")
		drifttest.FileIsYAML(t, ".goreleaser.yaml")
		drifttest.FileIsYAML(t, ".github/workflows/goreleaser.yaml")
		drifttest.OnlyGitChanges(t, ".goreleaser.yaml", ".github/workflows/goreleaser.yaml")
	}))
}
