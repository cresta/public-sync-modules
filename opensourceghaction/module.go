package opensourceghaction

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

var Module = childrenregistry.NewModule(Name, configYaml)

const Name = config.Name("github.com/getsyncer/public-sync-modules/opensourceghaction")
