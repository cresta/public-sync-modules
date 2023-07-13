package sync_autoapprove

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"github.com/cresta/syncer/sharedapi/syncer"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed autoapprove.yaml.template
var autoapproveTemplateStr string
var autoapproveTemplate = template.Must(template.New("autoapprove").Parse(autoapproveTemplateStr))

type Config struct {
	RunsOn string `yaml:"runs_on"`
}

func New() *Syncer {
	return &Syncer{}
}

type Syncer struct {
}

func (f *Syncer) Run(_ context.Context, runData *syncer.SyncRun) error {
	var cfg Config
	if err := runData.RunConfig.Decode(&cfg); err != nil {
		return fmt.Errorf("failed to unmarshal staticfile config: %w", err)
	}
	newPath := ".github/workflows/autoapprove.yaml"
	pathDir := filepath.Dir(newPath)
	if err := os.MkdirAll(pathDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", pathDir, err)
	}
	var into bytes.Buffer
	if err := autoapproveTemplate.Execute(&into, cfg); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}
	if err := os.WriteFile(newPath, into.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", newPath, err)
	}
	return nil
}

func (f *Syncer) Name() string {
	return "autoapprove"
}

func (f *Syncer) Priority() int {
	return syncer.PriorityNormal
}

var _ syncer.DriftSyncer = &Syncer{}
