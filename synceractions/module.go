package synceractions

import (
	_ "embed"

	// To make sure we get defaults of the latest versions of actions
	_ "github.com/getsyncer/public-sync-modules/latestversions"
	"github.com/getsyncer/syncer-core/config"
	"github.com/getsyncer/syncer-core/drift/syncers/templatefiles"
	"github.com/getsyncer/syncer-core/fxregistry"
)

func init() {
	fxregistry.Register(Module)
}

type Config struct {
	PrimaryBranch                                    string   `yaml:"primary_branch"`
	GithubRunner                                     string   `yaml:"github_runner"`
	ActionsCheckoutVersion                           string   `yaml:"actions_checkout_version"`
	SetupGoMods                                      []string `yaml:"setup_go_mods"`
	CreateGithubAppTokenVersion                      string   `yaml:"create_github_app_token_version"`
	RepushOnActor                                    string   `yaml:"repush_on_actor"`
	PeterMurrayWorkflowApplicationTokenActionVersion string   `yaml:"peter_murray_workflow_application_token_action_version"`
	RepushApp                                        string   `yaml:"repush_app"`
	RepushAppPem                                     string   `yaml:"repush_app_pem"`
}

//go:embed watchsynccomment.yaml.template
var watchsynccommentTemplateStr string

//go:embed checksyncer.yaml.template
var checksyncerTemplateStr string

const Name = config.Name("synceractions")

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name: Name,
	Files: map[string]string{
		".github/workflows/watchsynccomment.yaml": watchsynccommentTemplateStr,
		".github/workflows/checksync.yaml":        checksyncerTemplateStr,
	},
})
