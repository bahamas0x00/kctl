package services

import (
	"github.com/bahamas0x00/kctl/pkg/common"
)

type Service struct {
	Name              string   `json:"name"`
	Retries           int      `json:"retries"`
	Protocol          string   `json:"protocol"`
	Host              string   `json:"host"`
	Port              int      `json:"port"`
	Path              string   `json:"path"`
	ConnectTimeout    int      `json:"connect_timeout"`
	WriteTimeout      int      `json:"write_timeout"`
	ReadTimeout       int      `json:"read_timeout"`
	Tags              []string `json:"tags"`
	ClientCertificate struct {
		ID string `json:"id"`
	} `json:"client_certificate"`
	TlsVerify      bool     `json:"tls_verify"`
	TlsVerifyDepth *int     `json:"tls_verify_depth"`
	CaCertificates []string `json:"ca_certificates"`
	Enabled        bool     `json:"enabled"`
}

type Services struct {
	Data   []Service `json:"data"`
	Offset string    `json:"offset"`
}

// path: /serivces

// list all services
func ListAllServices(apiEndpoint string) (*common.HttpResponse, error) {
	return common.SendRequest("GET", apiEndpoint, []string{"services"}, nil)
}

// create a new service
func (s *Service) CreateNewService(apiEndpoint string) (*common.HttpResponse, error) {
	return common.SendRequest("POST", apiEndpoint, []string{"services"}, s)
}

// Delete a service
func (s *Service) DeleteService(apiEndpoint string) (*common.HttpResponse, error) {
	return common.SendRequest("DELETE", apiEndpoint, []string{"services", s.Name}, nil)
}

// update a service
func (s *Service) UpdateAService(apiEndpoint string) (*common.HttpResponse, error) {
	return common.SendRequest("PATCH", apiEndpoint, []string{"services", s.Name}, s)
}

// list all services in a workspace
func ListAllServicesInWorkspace(apiEndpoint, workspace string) (*common.HttpResponse, error) {
	return common.SendRequest("GET", apiEndpoint, []string{workspace, "services"}, nil)
}

// create a new service in a workspace
func (s *Service) CreateNewServiceInWorkspace(apiEndpoint, workspace string) (*common.HttpResponse, error) {
	return common.SendRequest("POST", apiEndpoint, []string{workspace, "services"}, s)
}

// Delete a service in workspace
func (s *Service) DeleteServiceInWorkspace(apiEndpoint, workspace string) (*common.HttpResponse, error) {
	return common.SendRequest("DELETE", apiEndpoint, []string{workspace, "services", s.Name}, nil)
}

// Update a Service in a workspace
func (s *Service) UpdateAServiceInWorkspace(apiEndpoint, workspace string) (*common.HttpResponse, error) {
	return common.SendRequest("PATCH", apiEndpoint, []string{workspace, "services", s.Name}, s)
}
