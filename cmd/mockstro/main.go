package main

import (
	"os"

	"github.com/Tapadoo/Mockstro/internal/cli"
	"github.com/Tapadoo/Mockstro/internal/tui"
)

var version = "dev"

func main() {
	cli.TUIFunc = tui.Run
	os.Exit(cli.Execute(version))
}
