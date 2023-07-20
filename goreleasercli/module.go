package goreleasercli

import (
	_ "embed"

	"github.com/cresta/syncer/sharedapi/drift/templatefiles"
	"github.com/cresta/syncer/sharedapi/syncer"
)

func init() {
	syncer.FxRegister(Module)
}

//go:embed .goreleaser.yaml.template
var templateStrGoReleaser string

//go:embed goreleaser.yaml.template
var templateStrActionReleaser string

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name: "goreleasercli",
	Files: map[string]string{
		".goreleaser.yaml":                  templateStrGoReleaser,
		".github/workflows/goreleaser.yaml": templateStrActionReleaser,
	},
	Priority: syncer.PriorityNormal,
	Decoder:  templatefiles.DefaultDecoder[Config](),
})

type Config struct {
	LinkerVarPath string
	MainDir       string
}
