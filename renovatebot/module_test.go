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
		drifttest.FileContains(t, "renovate.json", "docs.renovatebot.com")
	}))
	config2 := `
version: 1
config:
  extends:
    - hello-world2.json
logic:
  - source: github.com/getsyncer/public-sync-modules/renovatebot
syncs:
  - logic: renovatebot
    config:
      extends:
        - local>org/repo:hello-world.json
`
	t.Run("with-config", drifttest.WithRun(config2, files.SimpleState(map[string]string{}), func(t *testing.T, items *drifttest.Items) {
		items.TestRun.MustExitCode(t, 0)
		drifttest.FileContains(t, "renovate.json", "docs.renovatebot.com")
		drifttest.FileContains(t, "renovate.json", "hello-world.json")
		drifttest.FileContains(t, "renovate.json", "hello-world2.json")
	}))
}
