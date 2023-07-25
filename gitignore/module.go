package gitignore

import (
	_ "embed"
	"sort"
	"strings"

	"github.com/getsyncer/syncer/sharedapi/files/existingfileparser"

	"github.com/getsyncer/syncer/sharedapi/drift/templatefiles"
	"github.com/getsyncer/syncer/sharedapi/syncer"
)

func init() {
	syncer.FxRegister(Module)
}

type Config struct {
	Ignores        []string `yaml:"line"`
	postAutogenMsg string
	preAutogenMsg  string
}

func (c Config) SectionStart() string {
	return existingfileparser.RecommendedSectionStart
}

func (c Config) SectionEnd() string {
	return existingfileparser.RecommendedSectionEnd
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

func (c Config) ApplyParse(parse *existingfileparser.ParseResult) (Config, error) {
	c.preAutogenMsg = parse.PreAutogenMsg
	c.postAutogenMsg = parse.PostAutogenMsg
	return c, nil
}

//go:embed .gitignore.template
var templateStr string

const Name = syncer.Name("gitignore")

var Module = templatefiles.NewModule(templatefiles.NewModuleConfig[Config]{
	Name: Name,
	Files: map[string]string{
		".gitignore": templateStr,
	},
	Setup: &syncer.SetupMutator[Config]{
		Name:    Name,
		Mutator: syncer.DefaultParseMutator[Config](".gitignore"),
	},
})
