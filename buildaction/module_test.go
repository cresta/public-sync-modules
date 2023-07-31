package buildaction

import (
	"testing"

	"github.com/getsyncer/syncer-core/drifttest"
)

func TestModule(t *testing.T) {
	config := `
version: 1
logic:
  - source: github.com/getsyncer/syncer-core/drift/syncers/buildaction
syncs:
  - logic: buildaction
`
	t.Run("make-new-file", drifttest.WithRun(config, drifttest.ReasonableSampleFilesystem(), func(t *testing.T, items *drifttest.Items) {
		items.TestRun.MustExitCode(t, 0)
		drifttest.FileContains(t, ".github/workflows/buildgithubaction.yaml", "Build and test github action")
		drifttest.OnlyGitChanges(t, ".github/workflows/buildgithubaction.yaml")
	}))
}
