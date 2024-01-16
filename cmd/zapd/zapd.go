package main

import (
	"runtime/debug"

	"github.com/zapj/zap/cmd/zapd/cmd"
)

func init() {
	debug.SetMemoryLimit(1_000_000 * 50)
}

func main() {
	cmd.Execute()
}
