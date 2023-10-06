package renovatebot

import (
	"testing"

	"github.com/getsyncer/syncer-core/files"

	"github.com/getsyncer/syncer-core/drifttest"
)

func TestModule(t *testing.T) {
	config := `
version: 1
logic:
  - source: github.com/getsyncer/public-sync-modules/renovatebot
syncs:
  - logic: renovatebot
`
	t.Run("update-fresh-file", drifttest.WithRun(config, files.SimpleState(map[string]string{}), func(t *testing.T, items *drifttest.Items) {
		items.TestRun.MustExitCode(t, 0)
		drifttest.FileContains(t, ".renovate-autogen.json", "docs.renovatebot.com")
	}))
}
