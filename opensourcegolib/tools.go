//go:build syncer
// +build syncer

package opensourcegolib

// TODO: Auto generate this from the config.yaml file
import (
	_ "github.com/getsyncer/public-sync-modules/autoapprove"
	_ "github.com/getsyncer/public-sync-modules/buildgolib"
	_ "github.com/getsyncer/public-sync-modules/golangcilint"
	_ "github.com/getsyncer/public-sync-modules/setlicense"
	_ "github.com/getsyncer/public-sync-modules/synceractions"
)
