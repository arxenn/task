package main

import (
	"runtime/debug"

	"github.com/arxenn/task/cmd"
)

var version = "dev"

func main() {
	cmd.Execute(resolveVersion())
}

func resolveVersion() string {
	if version != "dev" {
		return version
	}
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "" && info.Main.Version != "(devel)" {
		return info.Main.Version
	}
	return "dev"
}
