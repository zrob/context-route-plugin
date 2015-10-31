package main

import (
	"github.com/cloudfoundry/cli/plugin"
	"github.com/zrob/context-route-plugin/commands"
)

type ContextRoutePlugin struct {
}

func main() {
	plugin.Start(new(ContextRoutePlugin))
}

func (crPlugin *ContextRoutePlugin) Run(cliConnection plugin.CliConnection, args []string) {

	if args[0] == "create-context-route" {
		//		if len(args) == 4 {
		commands.CreateContextRoute(cliConnection, args)
		//		} else {
		//			print help
		//		}
	}
}

func (crPlugin *ContextRoutePlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "context-route-plugin",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 1,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "create-context-route",
				Alias:    "ccr",
				HelpText: "creates a route with a path",
				UsageDetails: plugin.Usage{
					Usage: "create-context-route SPACE DOMAIN HOSTNAME PATH",
				},
			},
		},
	}
}
