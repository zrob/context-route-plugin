package commands

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cloudfoundry/cli/plugin"
	. "github.com/zrob/context-route-plugin/models"
	. "github.com/zrob/context-route-plugin/util"
)

func ListContextRoutes(cliConnection plugin.CliConnection, args []string) {
	mySpace, _ := cliConnection.GetCurrentSpace()

	output, err := cliConnection.CliCommandWithoutTerminalOutput("curl", fmt.Sprintf("v2/spaces/%s/routes", mySpace.Guid))
	FreakOut(err)
	routes := RoutesModel{}
	err = json.Unmarshal([]byte(output[0]), &routes)
	FreakOut(err)

	table := NewTable([]string{"space", "host", "domain", "path", "apps"})
	for _, route := range routes.Resources {
		output, err := cliConnection.CliCommandWithoutTerminalOutput("curl", route.Entity.DomainUrl)
		FreakOut(err)
		domain := DomainModel{}
		err = json.Unmarshal([]byte(output[0]), &domain)
		FreakOut(err)

		output, err = cliConnection.CliCommandWithoutTerminalOutput("curl", route.Entity.AppsUrl)
		FreakOut(err)
		apps := AppsModel{}
		err = json.Unmarshal([]byte(output[0]), &apps)
		FreakOut(err)

		appNames := []string{}
		for _, app := range apps.Resources {
			appNames = append(appNames, app.Entity.Name)
		}

		table.Add(mySpace.Name, route.Entity.Host, domain.Entity.Name, route.Entity.Path, strings.Join(appNames, ","))
	}
	table.Print()
}
