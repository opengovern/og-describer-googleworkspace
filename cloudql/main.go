package main

import (
	"github.com/opengovern/og-describer-googleworkspace/cloudql/googleworkspace"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: googleworkspace.Plugin})
}
