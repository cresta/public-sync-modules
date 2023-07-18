//go:build syncer
// +build syncer

package main

import (
	_ "github.com/cresta/public-sync-modules/autoapprove"
	_ "github.com/cresta/public-sync-modules/buildgolib"
	_ "github.com/cresta/public-sync-modules/golangcilint"
	_ "github.com/cresta/public-sync-modules/setlicense"
	_ "github.com/cresta/syncer/sharedapi/drift/staticfile"
	"github.com/cresta/syncer/sharedapi/syncer"
)

func main() {
	syncer.Apply(syncer.DefaultFxOptions())
}