package routes

import (
	"fmt"
	"net/http"
	"sync"
)

// ServiceID represents the Service ID structure in Kong API
type ServiceID struct {
	ID string `json:"id"`
}

// Route represents a Route in Kong API
type Route struct {
	ID                      string              `json:"id"`                         // Unique identifier for the route
	Name                    string              `json:"name"`                       // Name of the route
	Protocols               []string            `json:"protocols"`                  // List of protocols for the route (e.g., "http", "https")
	Methods                 []string            `json:"methods"`                    // HTTP methods (e.g., "GET", "POST")
	Hosts                   []string            `json:"hosts"`                      // List of hostnames for the route
	Paths                   []string            `json:"paths"`                      // List of URL paths for the route
	Headers                 map[string][]string `json:"headers"`                    // HTTP headers associated with the route
	HttpsRedirectStatusCode int                 `json:"https_redirect_status_code"` // Redirect status code for HTTPS
	RegexPriority           int                 `json:"regex_priority"`             // Regex priority for matching routes
	StripPath               bool                `json:"strip_path"`                 // Whether to strip the path before forwarding the request
	PathHandling            string              `json:"path_handling"`              // Path handling mode (e.g., "v0")
	PreserveHost            bool                `json:"preserve_host"`              // Whether to preserve the original host header
	RequestBuffering        bool                `json:"request_buffering"`          // Whether to buffer the request body
	ResponseBuffering       bool                `json:"response_buffering"`         // Whether to buffer the response body
	Tags                    []string            `json:"tags"`                       // Tags associated with the route
	Service                 ServiceID           `json:"service"`                    // The service linked to the route, with the service ID
}

// Routes represents the response for a list of routes in Kong API
type Routes struct {
	Data   []Route `json:"data"`   // List of routes
	Offset string  `json:"offset"` // Pagination offset
}

// request path
var pathComponents []string

// list(GET)
// 1. all routes associated with a service in a workspace  		/{workspace}/services/{serviceName}/routes
// 2. all routes in a workspace 								/{workspace}/routes
// 3. all routes associated with a service 						/services/{serviceName}/routes
// 4. all routes 												/routes
func ListAllRoutes(apiEndpoint, workspace string, serviceName string) (*http.Response, error) {
	return nil, nil
}

// create(POST)
// 1. a route associated with a service in a workspace  		/{workspace}/services/{serviceName}/routes
// 2. a route in a workspace 									/{workspace}/routes
// 3. a route associated with a service 						/services/{serviceName}/routes
// 4. a route 													/routes
func (r *Route) CreateRoute(apiEndpoint, workspace string, serviceName string) (*http.Response, error) {

	return nil, nil
}

// delete(DELETE)
// 1. a route 													/routes/{routeName}
// 2. a route associated with a service							/services/{serviceName}/routes/{routeName}
// 3. a route in a workspace									/{workspace}/routes/{routeName}
// 4. a route associated with a service in a workspace			/{workspace}/services/{serviceName}/routes/{routeName}
func (r *Route) DeleteRoute(apiEndpoint, workspace string, serviceName string) (*http.Response, error) {

	return nil, nil
}

// update (PATCH)
// 1. a route 													/routes/{routeName}
// 2. a route associated with a service							/services/{serviceName}/routes/{routeName}
// 3. a route in a workspace									/{workspace}/routes/{routeName}
// 4. a route associated with a service in a workspace			/{workspace}/services/{serviceName}/routes/{routeName}
func (r *Route) UpdateRoute(apiEndpoint, workspace string, serviceName string) (*http.Response, error) {
	return nil, nil
}

// batch create
func (r *Routes) BatchCreateRoutes(apiEndpoint, workspace string, serviceName string) ([]*http.Response, []error) {
	return nil, nil
}

// batch delete
func (r *Routes) BatchDeleteServices(apiEndpoint, workspace string, serviceName string) ([]*http.Response, []error) {
	return nil, nil
}

// batch update
func (r *Routes) BatchUpdateServices(apiEndpoint, workspace string, serviceName string) ([]*http.Response, []error) {
	return nil, nil
}

func batchExcuteRoutes(apiEndpoint string, workspace string, serviceName string, routes []*Route, operation string) ([]*http.Response, []error) {
	var wg sync.WaitGroup
	var responses []*http.Response
	var errs []error

	ch := make(chan struct {
		response *http.Response
		err      error
	}, len(routes))

	for _, route := range routes {
		wg.Add(1)
		go func(route *Route) {
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
