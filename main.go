package main

import (
	"github.com/hanymamdouh82/operatree/cmd"
	"github.com/hanymamdouh82/operatree/internal/activitylog"
)

// Injected at build time by ldflags
var (
	Version   = "dev"
	Commit    = "none"
	BuildDate = "unknown"
)

func main() {
	cmd.SetVersion(Version, Commit, BuildDate)
	activitylog.AppVersion = Version
	cmd.Execute()
}
