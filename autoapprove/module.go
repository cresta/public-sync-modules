package autoapprove

import (
	_ "embed"

	"github.com/cresta/syncer/sharedapi/drift/templatefiles"
	"github.com/cresta/syncer/sharedapi/syncer"
)

func init() {
	syncer.FxRegister(Module)
}

type Config struct {
	RunsOn string `yaml:"runs_on"`
}

//go:embed autoapprove.yaml.template
var templateStr string

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name: "autoapprove",
	Files: map[string]string{
		".github/workflows/autoapprove.yaml": templateStr,
	},
	Priority: syncer.PriorityNormal,
	Decoder:  templatefiles.DefaultDecoder[Config](),
})
