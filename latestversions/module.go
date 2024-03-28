package latestversions

import (
	"github.com/getsyncer/syncer-core/fxregistry"
	"github.com/getsyncer/syncer-core/syncer/planner/plannerhook"
)

func init() {
	fxregistry.Register(Module)
}

type Config struct {
	ActionsCheckoutVersion           string `yaml:"actions_checkout_version"`
	SetupGoVersion                   string `yaml:"setup_go_version"`
	PrimaryBranch                    string `yaml:"primary_branch"`
	CreateGithubAppTokenVersion      string `yaml:"create_github_app_token_version"`
	GithubRunner                     string `yaml:"github_runner"`
	GolangciLintVersion              string `yaml:"golangci_lint_version"`
	GithubAppTokenAction             string `yaml:"github_app_token_action"`
	GoSemanticReleaseActionVersion   string `yaml:"go_semantic_release_action_version"`
	ReviewdogActionActionlintVersion string `yaml:"reviewdog_action_actionlint_version"`
	GoreleaserActionVersion          string `yaml:"goreleaser_action_version"`
	CommitLintVersion                string `yaml:"commit_lint_version"`
}

// renovate: datasource=github-tags depName=actions/checkout versioning=docker
const actionsCheckout = "v4"

// renovate: datasource=github-tags depName=actions/setup-go versioning=docker
const setupGo = "v4"

// renovate: datasource=github-tags depName=golangci/golangci-lint versioning=docker
const golangCiLintVersion = "v3"

// renovate: datasource=github-tags depName=tibdex/github-app-token versioning=docker
const githubAppTokenActionVersion = "v2"

// renovate: datasource=github-tags depName=go-semantic-release/action versioning=docker
const goSemanticReleaseActionVersion = "v1"

// renovate: datasource=github-tags depName=reviewdog/action-actionlint versioning=docker
const reviewdogActionActionlintVersion = "v1"

// renovate: datasource=github-tags depName=goreleaser/goreleaser-action versioning=docker
const goReleaserActionVersion = "v5"

// renovate: datasource=github-tags depName=actions/create-github-app-token versioning=docker
const createGithubAppTokenVersion = "v1"

// renovate: datasource=github-tags depName=wagoid/commitlint-github-action versioning=docker
const commitLintVersion = "v6"

var Module = plannerhook.DefaultConfigModule("latest-defaults", Config{
	ActionsCheckoutVersion:           actionsCheckout,
	SetupGoVersion:                   setupGo,
	GolangciLintVersion:              golangCiLintVersion,
	GithubAppTokenAction:             githubAppTokenActionVersion,
	GoSemanticReleaseActionVersion:   goSemanticReleaseActionVersion,
	ReviewdogActionActionlintVersion: reviewdogActionActionlintVersion,
	CreateGithubAppTokenVersion:      createGithubAppTokenVersion,
	CommitLintVersion:                commitLintVersion,
	GoreleaserActionVersion:          goReleaserActionVersion,
	PrimaryBranch:                    "main",
	GithubRunner:                     "ubuntu-latest",
})
