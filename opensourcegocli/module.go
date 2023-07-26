package opensourcegocli

import (
	_ "embed"

	"github.com/getsyncer/syncer-core/syncer"
)

func init() {
	syncer.FxRegister(Module)
}

//go:embed config.yaml
var configYaml []byte

var Module = syncer.NewChildModule(Name, configYaml)

const Name = syncer.Name("github.com/getsyncer/public-sync-modules/opensourcegocli")
