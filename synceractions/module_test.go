package synceractions

import (
	"testing"

	"github.com/getsyncer/syncer-core/drifttest"
)

func TestModule(t *testing.T) {
	config := `
version: 1
logic:
  - source: github.com/getsyncer/syncer-core/drift/syncers/synceractions
syncs:
  - logic: synceractions
`
	t.Run("make-new-file", drifttest.WithRun(config, drifttest.ReasonableSampleFilesystem(), func(t *testing.T, items *drifttest.Items) {
		items.TestRun.MustExitCode(t, 0)
		drifttest.FileContains(t, ".github/workflows/watchsynccomment.yaml", "github.event.comment.body")
		drifttest.OnlyGitChanges(t, ".github/workflows/watchsynccomment.yaml", ".github/workflows/checksync.yaml")
	}))
	config2 := `
version: 1
config:
  repush_on_actor: bob
logic:
  - source: github.com/getsyncer/syncer-core/drift/syncers/synceractions
syncs:
  - logic: synceractions
`
	t.Run("make-new-file", drifttest.WithRun(config2, drifttest.ReasonableSampleFilesystem(), func(t *testing.T, items *drifttest.Items) {
		items.TestRun.MustExitCode(t, 0)
		drifttest.FileContains(t, ".github/workflows/watchsynccomment.yaml", "github.event.comment.body")
		drifttest.OnlyGitChanges(t, ".github/workflows/watchsynccomment.yaml", ".github/workflows/checksync.yaml")
		drifttest.FileIsYAML(t, ".github/workflows/watchsynccomment.yaml")
	}))
}
