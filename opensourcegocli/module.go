package opensourcegocli

import (
	_ "embed"

	"github.com/getsyncer/syncer/sharedapi/syncer"
)

func init() {
	syncer.FxRegister(Module)
}

//go:embed config.yaml
var configYaml []byte

var Module = syncer.NewChildModule("github.com/getsyncer/public-sync-modules/opensourcegocli", configYaml)
