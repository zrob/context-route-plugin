package commands

import (
	"encoding/json"
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
	. "github.com/zrob/context-route-plugin/models"
	. "github.com/zrob/context-route-plugin/util"
)

func CreateContextRoute(cliConnection plugin.CliConnection, args []string) {
	space := args[1]
	domain := args[2]
	host := args[3]
	path := args[4]

	myOrg, _ := cliConnection.GetCurrentOrg()

	output, err := cliConnection.CliCommandWithoutTerminalOutput("curl", fmt.Sprintf("v2/spaces?q=name:%s&q=organization_guid:%s", space, myOrg.Guid))
	FreakOut(err)
	spaces := SpacesModel{}
	err = json.Unmarshal([]byte(output[0]), &spaces)
	FreakOut(err)
	if len(spaces.Spaces) == 0 {
		fmt.Printf("Space %s not found", space)
		return
	}
	spaceGuid := spaces.Spaces[0].Metadata.Guid

	output, err = cliConnection.CliCommandWithoutTerminalOutput("curl", fmt.Sprintf("v2/domains?q=name:%s", domain))
	FreakOut(err)
	domains := DomainsModel{}
	err = json.Unmarshal([]byte(output[0]), &domains)
	FreakOut(err)
	if len(spaces.Spaces) == 0 {
		fmt.Printf("Domain %s not found", domain)
		return
	}
	domainGuid := domains.Resources[0].Metadata.Guid

	output, err = cliConnection.CliCommandWithoutTerminalOutput("curl", "v2/routes", "-X", "POST", "-d",
		fmt.Sprintf(`{"host":"%s","domain_guid":"%s","space_guid":"%s","path":"%s"}`, host, domainGuid, spaceGuid, path))
	FreakOut(err)
	route := RouteModel{}
	err = json.Unmarshal([]byte(output[0]), &route)
	FreakOut(err)
	routeGuid := route.Metadata.Guid

	if routeGuid == "" {
		error := ErrorModel{}
		err = json.Unmarshal([]byte(output[0]), &error)
		FreakOut(err)
		fmt.Printf("Failed to create route: %s", error.Description)
		return
	}

	fmt.Printf("Route successfully created.")
}
