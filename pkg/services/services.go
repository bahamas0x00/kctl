package services

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
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

		{
	  "data": [
	    {
	      "host": "example.internal",
	      "id": "49fd316e-c457-481c-9fc7-8079153e4f3c",
	      "name": "example-service",
	      "path": "/",
	      "port": 80,
	      "protocol": "http"
	    }
	  ],
	  "offset": "string"
	}
*/
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

var once sync.Once
var httpClient *http.Client

func init() {
	once.Do(func() {
		httpClient = &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
	})
}

// path: /serivces

// list all services
func ListAllServices(ApiEndpoint string) (*Services, error) {
	urlPath, err := url.JoinPath(ApiEndpoint, "services")
	if err != nil {
		return nil, fmt.Errorf("error parse request url: %v", err)
	}

	httpRequest, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		return nil, err
	}
	httpRequest.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("error request getting all services: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var services Services

	err = json.Unmarshal(body, &services)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling Json: %v", err)
	}
	return &services, nil

}

// create a new service
func (s *Service) CreateNewService(ApiEndpoint string) (*Service, error) {
	urlPath, err := url.JoinPath(ApiEndpoint, "services")
	if err != nil {
		return nil, fmt.Errorf("error parse request url: %v", err)
	}

	serviceData, err := json.Marshal(s)
	if err != nil {
		return nil, fmt.Errorf("invalid service data: %v", err)
	}

	httpRequest, err := http.NewRequest("POST", urlPath, bytes.NewBuffer(serviceData))
	if err != nil {
		return nil, err
	}
	httpRequest.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("error request create new service: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var service Service

	err = json.Unmarshal(body, &service)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling Json: %v", err)
	}
	return &service, nil
}

// Delete a service
func (s *Service) DeleteService(ApiEndpoint string) (bool, error) {
	urlPath, err := url.JoinPath(ApiEndpoint, "services", s.Name)
	if err != nil {
		return false, fmt.Errorf("error parse request url: %v", err)
	}

	serviceData, err := json.Marshal(s)
	if err != nil {
		return false, fmt.Errorf("invalid service data: %v", err)
	}

	httpRequest, err := http.NewRequest("DELETE", urlPath, bytes.NewBuffer(serviceData))
	if err != nil {
		return false, err
	}
	httpRequest.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(httpRequest)
	if err != nil {
		return false, fmt.Errorf("error request delete service: %v", err)
	}

	if resp.StatusCode != 204 {
		return false, fmt.Errorf("failed to delete service, code: %v", resp.StatusCode)
	}

	return true, nil

}

// update a service
func (s *Service) UpdateAService(ApiEndpoint string) (*Service, error) {
	urlPath, err := url.JoinPath(ApiEndpoint, "services", s.Name)
	if err != nil {
		return nil, fmt.Errorf("error parse request url: %v", err)
	}

	serviceData, err := json.Marshal(s)
	if err != nil {
		return nil, fmt.Errorf("invalid service data: %v", err)
	}

	httpRequest, err := http.NewRequest("PATCH", urlPath, bytes.NewBuffer(serviceData))
	if err != nil {
		return nil, err
	}
	httpRequest.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("error request update service: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var service Service

	err = json.Unmarshal(body, &service)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling Json: %v", err)
	}
	return &service, nil
}
