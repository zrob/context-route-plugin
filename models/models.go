package models

type MetadataModel struct {
	Guid string `json:"guid"`
}

type EntityModel struct {
	Name string `json:"name"`
}
type RouteEntityModel struct {
	Host string `json:"host"`
	Path string `json:"path"`
	DomainUrl string `json:"domain_url"`
	AppsUrl string `json:"apps_url"`
}

type ErrorModel struct {
	Description string `json:"description"`
}

type DomainsModel struct {
	Resources []DomainModel `json:"resources"`
}
type DomainModel struct {
	Metadata MetadataModel `json:"metadata"`
	Entity   EntityModel   `json:"entity"`
}
type RouteModel struct {
	Metadata MetadataModel    `json:"metadata"`
	Entity   RouteEntityModel `json:"entity"`
}
type RoutesModel struct {
	Resources []RouteModel `json:"resources"`
}
type SpaceModel struct {
	Metadata MetadataModel `json:"metadata"`
	Entity   EntityModel   `json:"entity"`
}
type SpacesModel struct {
	Resources []SpaceModel `json:"resources"`
}
type AppModel struct {
	Metadata MetadataModel `json:"metadata"`
	Entity   EntityModel   `json:"entity"`
}
type AppsModel struct {
	Resources []AppModel `json:"resources"`
}
