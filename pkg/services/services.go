package services

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/bahamas0x00/kctl/pkg/common"
)

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

type Services []Service

// request path

// list
// 1. all services 								/services
// 2. all services in a workspace				/{workspace}/services
func ListAllServices(apiEndpoint, workspace string) (*http.Response, error) {
	var pathComponents []string
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
	var pathComponents []string
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
	var pathComponents []string
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
	var pathComponents []string
	if common.IsStringSet(workspace) {
		pathComponents = append(pathComponents, workspace, "services", s.Name)
	} else {
		pathComponents = append(pathComponents, "services", s.Name)
	}
	return common.SendRequest("PATCH", apiEndpoint, pathComponents, s)
}

// batch create services
func (s *Services) BatchCreateServices(apiEndpoint, workspace string) ([]*http.Response, []error) {
	return batchExecuteServices(apiEndpoint, workspace, *s, "create")
}

// batch delete services
func (s *Services) BatchDeleteServices(apiEndpoint, workspace string) ([]*http.Response, []error) {
	return batchExecuteServices(apiEndpoint, workspace, *s, "delete")
}

// batch update services
func (s *Services) BatchUpdateServices(apiEndpoint, workspace string) ([]*http.Response, []error) {
	return batchExecuteServices(apiEndpoint, workspace, *s, "update")
}

func batchExecuteServices(apiEndpoint string, workspace string, services Services, operation string) ([]*http.Response, []error) {
	var wg sync.WaitGroup
	var responses []*http.Response
	var errs []error

	ch := make(chan struct {
		response *http.Response
		err      error
	}, len(services))

	for _, service := range services {
		wg.Add(1)
		go func(service Service) {
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
