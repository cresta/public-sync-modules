package buildgo

import (
	"testing"

	"github.com/getsyncer/syncer-core/drifttest"
)

func TestModule(t *testing.T) {
	config := `
version: 1
logic:
  - source: github.com/getsyncer/syncer-core/drift/syncers/buildgo
syncs:
  - logic: buildgo
`
	t.Run("make-new-file", drifttest.WithRun(config, drifttest.ReasonableSampleFilesystem(), func(t *testing.T, items *drifttest.Items) {
		items.TestRun.MustExitCode(t, 0)
		drifttest.FileContains(t, ".github/workflows/buildgo.yaml", "Build and test go code")
		drifttest.OnlyGitChanges(t, ".github/workflows/buildgo.yaml")
		drifttest.FileIsYAML(t, ".github/workflows/buildgo.yaml")
	}))
	config2 := `
version: 1
config:
  setup_go_mods:
    - blarg
logic:
  - source: github.com/getsyncer/syncer-core/drift/syncers/buildgo
syncs:
  - logic: buildgo
`
	t.Run("middle-config", drifttest.WithRun(config2, drifttest.ReasonableSampleFilesystem(), func(t *testing.T, items *drifttest.Items) {
		items.TestRun.MustExitCode(t, 0)
		drifttest.FileContains(t, ".github/workflows/buildgo.yaml", "Build and test go code")
		drifttest.FileContains(t, ".github/workflows/buildgo.yaml", "blarg")
		drifttest.OnlyGitChanges(t, ".github/workflows/buildgo.yaml")
	}))
}
