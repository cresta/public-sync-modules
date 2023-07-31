package opensourcegocli

import (
	"testing"

	"github.com/getsyncer/syncer-core/drifttest"
	"github.com/getsyncer/syncer-core/files"
)

func TestModule(t *testing.T) {
	config := `
version: 1
children:
  - source: github.com/getsyncer/public-sync-modules/opensourcegocli
`
	t.Run("update-fresh-file", drifttest.WithRun(config, &files.System[*files.State]{}, func(t *testing.T, items *drifttest.Items) {
		items.TestRun.MustExitCode(t, 0)
		drifttest.FileContains(t, ".github/workflows/buildgo.yaml", "go test -v -race")
		drifttest.FileContains(t, ".github/workflows/lintworkflows.yaml", "reviewdog/action-actionlint")
		drifttest.FileContains(t, "LICENSE", "Apache")
	}))
}
