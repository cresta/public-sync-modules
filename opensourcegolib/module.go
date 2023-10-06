package opensourcegolib

import (
	_ "embed"

	"github.com/getsyncer/syncer-core/config"
	"github.com/getsyncer/syncer-core/fxregistry"
	"github.com/getsyncer/syncer-core/syncer/childrenregistry"
)

func init() {
	fxregistry.Register(Module)
}

//go:embed config.yaml
var configYaml []byte

const Name = config.Name("github.com/getsyncer/public-sync-modules/opensourcegolib")

var Module = childrenregistry.NewModule(Name, configYaml)
