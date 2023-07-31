package gitignore

import (
	_ "embed"
	"sort"
	"strings"

	"github.com/getsyncer/syncer-core/drift/syncers/templatefiles/templatemutator"

	"github.com/getsyncer/syncer-core/fxregistry"

	"github.com/getsyncer/syncer-core/config"

	"github.com/getsyncer/syncer-core/files/stateloader"

	"github.com/getsyncer/syncer-core/drift/syncers/templatefiles"
)

func init() {
	fxregistry.Register(Module)
}

type Config struct {
	Ignores        []string `yaml:"line"`
	postAutogenMsg string
	preAutogenMsg  string
}

func (c Config) SectionStart() string {
	return stateloader.RecommendedSectionStart
}

func (c Config) SectionEnd() string {
	return stateloader.RecommendedSectionEnd
}

func (c Config) PostAutogenMsg() string {
	return c.postAutogenMsg
}
func (c Config) PreAutogenMsg() string {
	return c.preAutogenMsg
}

func (c Config) UniqueLines() []string {
	seen := map[string]struct{}{}
	ret := make([]string, 0, len(c.Ignores))
	for _, ignoreLine := range c.Ignores {
		ignoreLine = strings.TrimSpace(ignoreLine)
		if ignoreLine == "" {
			continue
		}
		if _, ok := seen[ignoreLine]; ok {
			continue
		}
		seen[ignoreLine] = struct{}{}
		ret = append(ret, ignoreLine)
	}
	sort.Strings(ret)
	return ret
}

func (c Config) ApplyParse(parse *stateloader.ParseResult) (Config, error) {
	c.preAutogenMsg = parse.PreAutogenMsg
	c.postAutogenMsg = parse.PostAutogenMsg
	return c, nil
}

//go:embed .gitignore.template
var templateStr string

const Name = config.Name("gitignore")

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name: Name,
	Files: map[string]string{
		".gitignore": templateStr,
	},
	Setup: &templatemutator.SetupMutator[Config]{
		Name:    Name,
		Mutator: templatemutator.DefaultParseMutator[Config](".gitignore"),
	},
})
