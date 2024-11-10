package routes

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/bahamas0x00/kctl/pkg/common"
)

// Route represents a Route in Kong API
type Route struct {
	Paths        []string    `json:"paths"`
	StripPath    bool        `json:"strip_path"`
	Sources      interface{} `json:"sources"`
	PreserveHost bool        `json:"preserve_host"`
	// Service      *struct {
	// 	ID string `json:"id"`
	// } `json:"service"`
	Destinations            interface{} `json:"destinations"`
	Methods                 []string    `json:"methods"`
	ID                      string      `json:"id"`
	PathHandling            string      `json:"path_handling"`
	Protocols               []string    `json:"protocols"`
	Hosts                   *[]string   `json:"hosts"`
	Snis                    interface{} `json:"snis"`
	Headers                 interface{} `json:"headers"`
	RegexPriority           int         `json:"regex_priority"`
	Tags                    *[]string   `json:"tags"`
	RequestBuffering        bool        `json:"request_buffering"`
	ResponseBuffering       bool        `json:"response_buffering"`
	HTTPSRedirectStatusCode int         `json:"https_redirect_status_code"`
	Name                    string      `json:"name"`
}

// Routes represents the response for a list of routes in Kong API
type Routes struct {
	Data []Route `json:"data"` // List of routes
	Next *string `json:"next"` // Pagination offset
}

// list(GET)
// 1. all routes associated with a service in a workspace  		/{workspace}/services/{serviceName}/routes
// 2. all routes in a workspace 								/{workspace}/routes
// 3. all routes associated with a service 						/services/{serviceName}/routes
// 4. all routes 												/routes
func ListAllRoutes(apiEndpoint, workspace string, serviceName string) (*http.Response, error) {
	var pathComponents []string
	if common.IsStringSet(workspace) && common.IsStringSet(serviceName) {
		pathComponents = append(pathComponents, workspace, "services", serviceName, "routes")
	} else if common.IsStringSet(workspace) {
		pathComponents = append(pathComponents, workspace, "routes")
	} else if common.IsStringSet(serviceName) {
		pathComponents = append(pathComponents, "services", serviceName, "routes")
	} else {
		pathComponents = append(pathComponents, "routes")
	}

	return common.SendRequest("GET", apiEndpoint, pathComponents, nil)
}

// create(POST)
// 1. a route associated with a service in a workspace  		/{workspace}/services/{serviceName}/routes
// 2. a route in a workspace 									/{workspace}/routes
// 3. a route associated with a service 						/services/{serviceName}/routes
// 4. a route 													/routes
func (r *Route) CreateRoute(apiEndpoint, workspace string, serviceName string) (*http.Response, error) {
	var pathComponents []string
	if common.IsStringSet(workspace) && common.IsStringSet(serviceName) {
		pathComponents = append(pathComponents, workspace, "services", serviceName, "routes")
	} else if common.IsStringSet(workspace) {
		pathComponents = append(pathComponents, workspace, "routes")
	} else if common.IsStringSet(serviceName) {
		pathComponents = append(pathComponents, "services", serviceName, "routes")
	} else {
		pathComponents = append(pathComponents, "routes")
	}
	return common.SendRequest("POST", apiEndpoint, pathComponents, r)
}

// delete(DELETE)
// 1. a route 													/routes/{routeName}
// 2. a route associated with a service							/services/{serviceName}/routes/{routeName}
// 3. a route in a workspace									/{workspace}/routes/{routeName}
// 4. a route associated with a service in a workspace			/{workspace}/services/{serviceName}/routes/{routeName}
func (r *Route) DeleteRoute(apiEndpoint, workspace string, serviceName string) (*http.Response, error) {
	var pathComponents []string
	if common.IsStringSet(workspace) && common.IsStringSet(serviceName) {
		pathComponents = append(pathComponents, workspace, "services", serviceName, "routes", r.Name)
	} else if common.IsStringSet(workspace) {
		pathComponents = append(pathComponents, workspace, "routes", r.Name)
	} else if common.IsStringSet(serviceName) {
		pathComponents = append(pathComponents, "services", serviceName, "routes", r.Name)
	} else {
		pathComponents = append(pathComponents, "routes", r.Name)
	}
	return common.SendRequest("DELETE", apiEndpoint, pathComponents, nil)
}

// update (PATCH)
// 1. a route 													/routes/{routeName}
// 2. a route associated with a service							/services/{serviceName}/routes/{routeName}
// 3. a route in a workspace									/{workspace}/routes/{routeName}
// 4. a route associated with a service in a workspace			/{workspace}/services/{serviceName}/routes/{routeName}
func (r *Route) UpdateRoute(apiEndpoint, workspace string, serviceName string) (*http.Response, error) {
	var pathComponents []string
	if common.IsStringSet(workspace) && common.IsStringSet(serviceName) {
		pathComponents = append(pathComponents, workspace, "services", serviceName, "routes", r.Name)
	} else if common.IsStringSet(workspace) {
		pathComponents = append(pathComponents, workspace, "routes", r.Name)
	} else if common.IsStringSet(serviceName) {
		pathComponents = append(pathComponents, "services", serviceName, "routes", r.Name)
	} else {
		pathComponents = append(pathComponents, "routes", r.Name)
	}
	return common.SendRequest("PATCH", apiEndpoint, pathComponents, r)
}

// batch create
func (r *Routes) BatchCreateRoutes(apiEndpoint, workspace string, serviceName string) ([]*http.Response, []error) {
	return batchExecuteRoutes(apiEndpoint, workspace, serviceName, *r, "create")
}

// batch delete
func (r *Routes) BatchDeleteServices(apiEndpoint, workspace string, serviceName string) ([]*http.Response, []error) {
	return batchExecuteRoutes(apiEndpoint, workspace, serviceName, *r, "delete")
}

// batch update
func (r *Routes) BatchUpdateServices(apiEndpoint, workspace string, serviceName string) ([]*http.Response, []error) {
	return batchExecuteRoutes(apiEndpoint, workspace, serviceName, *r, "update")
}

func batchExecuteRoutes(apiEndpoint string, workspace string, serviceName string, routes Routes, operation string) ([]*http.Response, []error) {
	var wg sync.WaitGroup
	var responses []*http.Response
	var errs []error

	ch := make(chan struct {
		response *http.Response
		err      error
	}, len(routes.Data))

	for _, route := range routes.Data {
		wg.Add(1)
		go func(route Route) {
			defer wg.Done()
			var resp *http.Response
			var err error

			switch operation {
			case "create":
				resp, err = route.CreateRoute(apiEndpoint, workspace, serviceName)
			case "update":
				resp, err = route.UpdateRoute(apiEndpoint, workspace, serviceName)
			case "delete":
				resp, err = route.DeleteRoute(apiEndpoint, workspace, serviceName)
			default:
				err = fmt.Errorf("invalid operation type: %s", operation)
			}

			ch <- struct {
				response *http.Response
				err      error
			}{response: resp, err: err}
		}(route)
	}

	// wait all goroutine complete
	wg.Wait()
	close(ch)

	for result := range ch {
		if result.err != nil {
			errs = append(errs, result.err)
		} else {
			responses = append(responses, result.response)
		}
	}

	return responses, errs
}
