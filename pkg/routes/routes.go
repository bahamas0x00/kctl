package routes

type ServiceID struct {
	ID string `json:"id"`
}

type Route struct {
	Hosts   []string  `json:"hosts"`
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Paths   []string  `json:"paths"`
	Service ServiceID `json:"service"`
}

type Data struct {
	Routes []Route `json:"routes"`
}

type RoutesResponse struct {
	Data   []Data `json:"data"`
	Offset string `json:"offset"`
}

// list all routes associated with a service in a workspace
func ListAllRoutesAssociatedWithServiceInworkspace(service, workspace string) (RoutesResponse, error) {
	//TODO
	return nil, nil
}

// create a new route associated with a service in a workspace
func (r *Route) CreateNewRouteAssociatedWithServiceInWorkspace(service, workspace string) error {
	return nil
}
