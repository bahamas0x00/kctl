package routes

// ServiceID represents the Service ID structure in Kong API
type ServiceID struct {
	ID string `json:"id"`
}

// Route represents a Route in Kong API
type Route struct {
	ID                       string            `json:"id"`                        // Unique identifier for the route
	Name                     string            `json:"name"`                      // Name of the route
	Protocols                []string          `json:"protocols"`                 // List of protocols for the route (e.g., "http", "https")
	Methods                  []string          `json:"methods"`                   // HTTP methods (e.g., "GET", "POST")
	Hosts                    []string          `json:"hosts"`                     // List of hostnames for the route
	Paths                    []string          `json:"paths"`                     // List of URL paths for the route
	Headers                  map[string][]string `json:"headers"`                 // HTTP headers associated with the route
	HttpsRedirectStatusCode  int               `json:"https_redirect_status_code"`// Redirect status code for HTTPS
	RegexPriority            int               `json:"regex_priority"`            // Regex priority for matching routes
	StripPath                bool              `json:"strip_path"`                // Whether to strip the path before forwarding the request
	PathHandling             string            `json:"path_handling"`             // Path handling mode (e.g., "v0")
	PreserveHost             bool              `json:"preserve_host"`             // Whether to preserve the original host header
	RequestBuffering         bool              `json:"request_buffering"`         // Whether to buffer the request body
	ResponseBuffering        bool              `json:"response_buffering"`        // Whether to buffer the response body
	Tags                     []string          `json:"tags"`                      // Tags associated with the route
	Service                 ServiceID         `json:"service"`                   // The service linked to the route, with the service ID
}

// Routes represents the response for a list of routes in Kong API
type Routes struct {
	Data   []Route `json:"data"`   // List of routes
	Offset string  `json:"offset"` // Pagination offset
}


