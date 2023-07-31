package golangcilint

import (
	"testing"

	"github.com/getsyncer/syncer-core/files"

	"github.com/getsyncer/syncer-core/drifttest"
)

func TestModule(t *testing.T) {
	config := `
version: 1
logic:
  - source: github.com/getsyncer/syncer-core/drift/syncers/golangcilint
  - source: github.com/getsyncer/syncer-core/drift/syncers/buildgo
syncs:
  - logic: golangcilint
  - logic: buildgo
`
	t.Run("update-fresh-file", drifttest.WithRun(config, &files.System[*files.State]{}, func(t *testing.T, items *drifttest.Items) {
		items.TestRun.MustExitCode(t, 0)
		drifttest.FileContains(t, ".golangci.yml", "goimports")
		drifttest.FileContains(t, ".github/workflows/buildgo.yaml", "golangci/golangci-lint-action")
		drifttest.FileIsYAML(t, ".golangci.yml")
		drifttest.FileIsYAML(t, ".github/workflows/buildgo.yaml")
		drifttest.OnlyGitChanges(t, ".golangci.yml", ".github/workflows/buildgo.yaml")
	}))
}
