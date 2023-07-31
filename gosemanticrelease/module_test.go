package gosemanticrelease

import (
	"testing"

	"github.com/getsyncer/syncer-core/drifttest"
	"github.com/getsyncer/syncer-core/files"
)

func TestModule(t *testing.T) {
	config := `
version: 1
logic:
  - source: github.com/getsyncer/syncer-core/drift/syncers/gosemanticrelease
syncs:
  - logic: gosemanticrelease
  - logic: buildgo
`
	t.Run("update-fresh-file", drifttest.WithRun(config, &files.System[*files.State]{}, func(t *testing.T, items *drifttest.Items) {
		items.TestRun.MustExitCode(t, 0)
		drifttest.FileContains(t, ".github/workflows/buildgo.yaml", "go-semantic-release/action")
		drifttest.OnlyGitChanges(t, ".github/workflows/buildgo.yaml")
	}))
}
