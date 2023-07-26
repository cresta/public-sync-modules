//go:build syncer
// +build syncer

package opensourcegolib

// TODO: Auto generate this from the config.yaml file
import (
	_ "github.com/getsyncer/public-sync-modules/buildgo"
	_ "github.com/getsyncer/public-sync-modules/gitignore"
	_ "github.com/getsyncer/public-sync-modules/golangcilint"
	_ "github.com/getsyncer/public-sync-modules/gosemanticrelease"
	_ "github.com/getsyncer/public-sync-modules/lintworkflows"
	_ "github.com/getsyncer/public-sync-modules/setlicense"
	_ "github.com/getsyncer/public-sync-modules/synceractions"
)
