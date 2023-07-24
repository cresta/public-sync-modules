package synceractions

import (
	_ "embed"

	"github.com/getsyncer/syncer/sharedapi/drift/templatefiles"
	"github.com/getsyncer/syncer/sharedapi/syncer"
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

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name: "synceractions",
	Files: map[string]string{
		".github/workflows/watchsynccomment.yaml": watchsynccommentTemplateStr,
		".github/workflows/checksync.yaml":        checksyncerTemplateStr,
	},
	Priority: syncer.PriorityNormal,
	Decoder:  templatefiles.DefaultDecoder[Config](),
})
