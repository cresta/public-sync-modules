package lintworkflows

import (
	"testing"

	"github.com/getsyncer/syncer-core/drifttest"
	"github.com/getsyncer/syncer-core/files"
)

func TestModule(t *testing.T) {
	config := `
version: 1
logic:
  - source: github.com/getsyncer/syncer-core/drift/syncers/lintworkflows
syncs:
  - logic: lintworkflows
`
	t.Run("update-fresh-file", drifttest.WithRun(config, &files.System[*files.State]{}, func(t *testing.T, items *drifttest.Items) {
		items.TestRun.MustExitCode(t, 0)
		drifttest.FileContains(t, ".github/workflows/lintworkflows.yaml", "reviewdog/action-actionlint")
		drifttest.OnlyGitChanges(t, ".github/workflows/lintworkflows.yaml")
	}))
}
