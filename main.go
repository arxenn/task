package main

import (
	"github.com/arxenn/task/cmd"
)

// TODO LIST:
// - Complete windows shell integration

var version = "dev"

func main() {
	cmd.Execute(version)
}
