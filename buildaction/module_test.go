package buildaction

import (
	"testing"

	_ "github.com/getsyncer/public-sync-modules/renovatebot"
	"github.com/getsyncer/syncer-core/drifttest"
)

func TestModule(t *testing.T) {
	const config = `
version: 1
logic:
  - source: github.com/getsyncer/public-sync-modules/buildaction
syncs:
  - logic: buildaction
`
	t.Run("make-new-file", drifttest.WithRun(config, drifttest.ReasonableSampleFilesystem(), func(t *testing.T, items *drifttest.Items) {
		items.TestRun.MustExitCode(t, 0)
		drifttest.FileContains(t, ".github/workflows/buildgithubaction.yaml", "Build and test github action")
		drifttest.FileDoesNotContain(t, ".github/workflows/buildgithubaction.yaml", "actions/checkout@\n")
		drifttest.OnlyGitChanges(t, ".github/workflows/buildgithubaction.yaml")
	}))
	const config2 = `
version: 1
config:
  actions_checkout_version: v1234
logic:
  - source: github.com/getsyncer/public-sync-modules/buildaction
syncs:
  - logic: buildaction
`
	t.Run("make-new-file", drifttest.WithRun(config2, drifttest.ReasonableSampleFilesystem(), func(t *testing.T, items *drifttest.Items) {
		items.TestRun.MustExitCode(t, 0)
		drifttest.FileContains(t, ".github/workflows/buildgithubaction.yaml", "actions/checkout@v1234")
	}))
	const config3 = `
version: 1
logic:
  - source: github.com/getsyncer/public-sync-modules/buildaction
  - source: github.com/getsyncer/public-sync-modules/renovatebot
syncs:
  - logic: renovatebot
  - logic: buildaction
    config:
      actions_checkout_version: v1235
`
	t.Run("make-new-file", drifttest.WithRun(config3, drifttest.ReasonableSampleFilesystem(), func(t *testing.T, items *drifttest.Items) {
		items.TestRun.MustExitCode(t, 0)
		drifttest.FileContains(t, ".renovate-autogen.json", ".github/workflows/buildgithubaction.yaml")
		drifttest.FileContains(t, ".github/workflows/buildgithubaction.yaml", "actions/checkout@v1235")
	}))
}
