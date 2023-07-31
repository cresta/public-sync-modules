package setlicense

import (
	"testing"

	"github.com/getsyncer/syncer-core/drifttest"
	"github.com/getsyncer/syncer-core/files"
)

func TestModule(t *testing.T) {
	config := `
version: 1
logic:
  - source: github.com/getsyncer/public-sync-modules/setlicense
syncs:
  - logic: setlicense
    config:
      license: Apache-2.0
`
	t.Run("update-fresh-file", drifttest.WithRun(config, &files.System[*files.State]{}, func(t *testing.T, items *drifttest.Items) {
		items.TestRun.MustExitCode(t, 0)
		drifttest.FileContains(t, "LICENSE", "Apache")
		drifttest.OnlyGitChanges(t, "LICENSE")
	}))
}
