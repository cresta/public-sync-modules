package opensourcegolib

import (
	_ "embed"

	"github.com/getsyncer/syncer/sharedapi/syncer"
)

func init() {
	syncer.FxRegister(Module)
}

//go:embed config.yaml
var configYaml []byte

const Name = syncer.Name("github.com/getsyncer/public-sync-modules/opensourcegolib")

var Module = syncer.NewChildModule(Name, configYaml)
