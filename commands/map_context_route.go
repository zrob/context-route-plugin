package commands

import (
	"encoding/json"
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
	. "github.com/zrob/context-route-plugin/models"
	. "github.com/zrob/context-route-plugin/util"
)

func MapContextRoute(cliConnection plugin.CliConnection, args []string) {
	app := args[1]
	domain := args[2]
	host := args[3]
	path := args[4]

	mySpace, _ := cliConnection.GetCurrentSpace()

	output, err := cliConnection.CliCommandWithoutTerminalOutput("curl", fmt.Sprintf("v2/apps?q=name:%s&q=space_guid:%s", app, mySpace.Guid))
	FreakOut(err)
	apps := AppsModel{}
	err = json.Unmarshal([]byte(output[0]), &apps)
	FreakOut(err)
	if len(apps.Resources) == 0 {
		fmt.Printf("App %s not found", app)
		return
	}
	appGuid := apps.Resources[0].Metadata.Guid

	output, err = cliConnection.CliCommandWithoutTerminalOutput("curl", fmt.Sprintf("v2/domains?q=name:%s", domain))
	FreakOut(err)
	domains := DomainsModel{}
	err = json.Unmarshal([]byte(output[0]), &domains)
	FreakOut(err)
	if len(domains.Resources) == 0 {
		fmt.Printf("Domain %s not found", domain)
		return
	}
	domainGuid := domains.Resources[0].Metadata.Guid

	output, err = cliConnection.CliCommandWithoutTerminalOutput("curl",
		fmt.Sprintf("v2/routes?q=domain_guid:%s&q=host:%s&q=path:%s", domainGuid, host, path))
	FreakOut(err)
	routes := RoutesModel{}
	err = json.Unmarshal([]byte(output[0]), &routes)
	FreakOut(err)
	if len(routes.Resources) == 0 {
		fmt.Printf("Route not found host: %s, path: %s", host, path)
		return
	}
	routeGuid := routes.Resources[0].Metadata.Guid

	output, err = cliConnection.CliCommandWithoutTerminalOutput("curl",
		fmt.Sprintf("v2/apps/%s/routes/%s", appGuid, routeGuid), "-X", "PUT")
	FreakOut(err)
	mappedApp := AppModel{}
	err = json.Unmarshal([]byte(output[0]), &mappedApp)
	FreakOut(err)
	if mappedApp.Metadata.Guid == "" {
		error := ErrorModel{}
		err = json.Unmarshal([]byte(output[0]), &error)
		FreakOut(err)
		fmt.Printf("Failed to map route: %s", error.Description)
		return
	}

	fmt.Printf("Route successfully mapped.")
}
