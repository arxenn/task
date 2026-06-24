package main

import (
	"os/exec"
	"runtime/debug"
	"strings"

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

	if v, err := getGitVersion(); err == nil && v != "" {
		return v
	}

	return "dev"
}

func getGitVersion() (string, error) {
	cmd := exec.Command("git", "describe", "--tags")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}
