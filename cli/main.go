package main

import (
	"github.com/amirhossein2831/httpServerGo/cli/cmd"
	"github.com/amirhossein2831/httpServerGo/src/App"
)

func main() {
	// configure the app
	App.Configure()

	// run cobra command
	cmd.Execute()
}
