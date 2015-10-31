package commands

import (
	"encoding/json"
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
	. "github.com/zrob/context-route-plugin/models"
	. "github.com/zrob/context-route-plugin/util"
)

func DeleteContextRoute(cliConnection plugin.CliConnection, args []string) {
	domain := args[1]
	host := args[2]
	path := args[3]

	output, err := cliConnection.CliCommandWithoutTerminalOutput("curl", fmt.Sprintf("v2/domains?q=name:%s", domain))
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
		fmt.Sprintf("v2/routes/%s", routeGuid), "-X", "DELETE")
	FreakOut(err)

	fmt.Printf("Route successfully deleted.")
}
