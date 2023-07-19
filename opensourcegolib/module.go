package opensourcegolib

import (
	_ "embed"

	"github.com/cresta/syncer/sharedapi/syncer"
)

func init() {
	syncer.FxRegister(Module)
}

//go:embed config.yaml
var configYaml []byte

var Module = syncer.NewChildModule("github.com/cresta/public-sync-modules/opensourcegolib", configYaml)
