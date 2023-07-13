package sync_autoapprove

import (
	_ "embed"
	"github.com/cresta/syncer/sharedapi/syncer"
	"github.com/cresta/syncer/sharedapi/templatefiles"
	"text/template"
)

func init() {
	syncer.FxRegister(Module)
}

type Config struct {
	RunsOn string `yaml:"runs_on"`
}

//go:embed autoapprove.yaml.template
var autoapproveTemplateStr string
var autoapproveTemplate = template.Must(template.New("autoapprove").Parse(autoapproveTemplateStr))

var Module = templatefiles.NewModule("autoapprove", map[string]*template.Template{
	".github/workflows/autoapprove.yaml": autoapproveTemplate,
}, syncer.PriorityNormal, func(runConfig syncer.RunConfig) (interface{}, error) {
	var cfg Config
	if err := runConfig.Decode(&cfg); err != nil {
		return nil, err
	}
	return cfg, nil
})
