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

type Routes struct {
	Data   []Route `json:"data"`
	Offset string  `json:"offset"`
}

// list all routes associated with a service in a workspace
func ListAllRoutesAssociatedWithServiceInworkspace(service, workspace string) (Routes, error)

// create a new route associated with a service in a workspace
func (r *Route) CreateNewRouteAssociatedWithServiceInWorkspace(service, workspace string) error {
	return nil
}
