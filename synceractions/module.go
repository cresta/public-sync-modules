package synceractions

import (
	_ "embed"

	"github.com/getsyncer/syncer-core/drift/templatefiles"
	"github.com/getsyncer/syncer-core/syncer"
)

func init() {
	syncer.FxRegister(Module)
}

type Config struct {
}

//go:embed watchsynccomment.yaml.template
var watchsynccommentTemplateStr string

//go:embed checksyncer.yaml.template
var checksyncerTemplateStr string

const Name = syncer.Name("synceractions")

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name: Name,
	Files: map[string]string{
		".github/workflows/watchsynccomment.yaml": watchsynccommentTemplateStr,
		".github/workflows/checksync.yaml":        checksyncerTemplateStr,
	},
})
