package synceractions

import (
	_ "embed"

	"github.com/getsyncer/syncer-core/fxregistry"

	"github.com/getsyncer/syncer-core/config"

	"github.com/getsyncer/syncer-core/drift/syncers/templatefiles"
)

func init() {
	fxregistry.Register(Module)
}

type Config struct {
}

//go:embed watchsynccomment.yaml.template
var watchsynccommentTemplateStr string

//go:embed checksyncer.yaml.template
var checksyncerTemplateStr string

const Name = config.Name("synceractions")

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name: Name,
	Files: map[string]string{
		".github/workflows/watchsynccomment.yaml": watchsynccommentTemplateStr,
		".github/workflows/checksync.yaml":        checksyncerTemplateStr,
	},
})
