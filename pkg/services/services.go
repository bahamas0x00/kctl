package services

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/bahamas0x00/kctl/pkg/common"
)

/*
{
  "name": "my-service",
  "retries": 5,
  "protocol": "http",
  "host": "example.com",
  "port": 80,
  "path": "/some_api",
  "connect_timeout": 6000,
  "write_timeout": 6000,
  "read_timeout": 6000,
  "tags": [
    "user-level"
  ],
  "client_certificate": {
    "id": "4e3ad2e4-0bc4-4638-8e34-c84a417ba39b"
  },
  "tls_verify": true,
  "tls_verify_depth": null,
  "ca_certificates": [
    "4e3ad2e4-0bc4-4638-8e34-c84a417ba39b"
  ],
  "enabled": true
}
*/
// Service represents a single service in Kong Gateway.
type Service struct {
	Name              string   `json:"name"`            // Service name
	Retries           int      `json:"retries"`         // Number of retries on failure
	Protocol          string   `json:"protocol"`        // Protocol used by the service (e.g., http, https)
	Host              string   `json:"host"`            // Host address of the service
	Port              int      `json:"port"`            // Port the service is listening on
	Path              string   `json:"path"`            // Path for the service
	ConnectTimeout    int      `json:"connect_timeout"` // Timeout for establishing connections (in seconds)
	WriteTimeout      int      `json:"write_timeout"`   // Timeout for writing data to the service (in seconds)
	ReadTimeout       int      `json:"read_timeout"`    // Timeout for reading data from the service (in seconds)
	Tags              []string `json:"tags"`            // Tags associated with the service
	ClientCertificate struct {
		ID string `json:"id"` // ID of the client certificate (optional)
	} `json:"client_certificate"`
	TlsVerify      bool     `json:"tls_verify"`       // Whether to verify the TLS certificate
	TlsVerifyDepth *int     `json:"tls_verify_depth"` // Optional field for TLS verification depth
	CaCertificates []string `json:"ca_certificates"`  // List of CA certificates
	Enabled        bool     `json:"enabled"`          // Whether the service is enabled or not
}

// Services represents a response containing a list of services and pagination information.
type Services struct {
	Data   []Service `json:"data"`   // List of services
	Offset string    `json:"offset"` // Pagination offset for the next set of results
}

// request path
var pathComponents []string

// list
// 1. all services 								/services
// 2. all services in a workspace				/{workspace}/services
func ListAllServices(apiEndpoint, workspace string) (*http.Response, error) {
	if common.IsStringSet(workspace) {
		pathComponents = append(pathComponents, workspace, "services")
	} else {
		pathComponents = append(pathComponents, "services")
	}
	return common.SendRequest("GET", apiEndpoint, pathComponents, nil)
}

// create
// 1. a service									/services
// 2. a service in a workspace					/{workspace}/services
func (s *Service) CreateService(apiEndpoint, workspace string) (*http.Response, error) {
	if common.IsStringSet(workspace) {
		pathComponents = append(pathComponents, workspace, "services")
	} else {
		pathComponents = append(pathComponents, "services")
	}
	return common.SendRequest("POST", apiEndpoint, pathComponents, s)
}

// delete
// 1. a service									/services/{service_name}
// 2. a service in a workspace					/{workspace}/services/{service_name}
func (s *Service) DeleteService(apiEndpoint, workspace string) (*http.Response, error) {
	if common.IsStringSet(workspace) {
		pathComponents = append(pathComponents, workspace, "services", s.Name)
	} else {
		pathComponents = append(pathComponents, "services", s.Name)
	}
	return common.SendRequest("DELETE", apiEndpoint, pathComponents, nil)
}

// update a service or update a service in a workspace
// 1. a service									/services/{serviceName}
// 2. a service in a workspace					/{workspace}/services/{serviceName}
func (s *Service) UpdateService(apiEndpoint, workspace string) (*http.Response, error) {
	if common.IsStringSet(workspace) {
		pathComponents = append(pathComponents, workspace, "services", s.Name)
	} else {
		pathComponents = append(pathComponents, "services", s.Name)
	}
	return common.SendRequest("PATCH", apiEndpoint, pathComponents, s)
}

// batch create services
func (s *Services) BatchCreateServices(apiEndpoint, workspace string) ([]*http.Response, []error) {
	services := common.ConvertToPointers(s.Data)
	return batchExcuteServices(apiEndpoint, workspace, services, "create")
}

// batch delete services
func (s *Services) BatchDeleteServices(apiEndpoint, workspace string) ([]*http.Response, []error) {
	services := common.ConvertToPointers(s.Data)
	return batchExcuteServices(apiEndpoint, workspace, services, "delete")
}

// batch update services
func (s *Services) BatchUpdateServices(apiEndpoint, workspace string) ([]*http.Response, []error) {
	services := common.ConvertToPointers(s.Data)
	return batchExcuteServices(apiEndpoint, workspace, services, "update")
}

func batchExcuteServices(apiEndpoint string, workspace string, services []*Service, operation string) ([]*http.Response, []error) {
	var wg sync.WaitGroup
	var responses []*http.Response
	var errs []error

	ch := make(chan struct {
		response *http.Response
		err      error
	}, len(services))

	for _, service := range services {
		wg.Add(1)
		go func(service *Service) {
			defer wg.Done()
			var resp *http.Response
			var err error

			switch operation {
			case "create":
				resp, err = service.CreateService(apiEndpoint, workspace)
			case "update":
				resp, err = service.UpdateService(apiEndpoint, workspace)
			case "delete":
				resp, err = service.DeleteService(apiEndpoint, workspace)
			default:
				err = fmt.Errorf("invalid operation type: %s", operation)
			}

			ch <- struct {
				response *http.Response
				err      error
			}{response: resp, err: err}
		}(service)
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
