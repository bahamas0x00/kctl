package upstreams

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/bahamas0x00/kctl/pkg/common"
)

type Upstream struct {
	HashFallbackHeader   interface{} `json:"hash_fallback_header"`
	HashFallbackQueryArg interface{} `json:"hash_fallback_query_arg"`
	ID                   string      `json:"id"`
	HashOnHeader         interface{} `json:"hash_on_header"`
	HashOnQueryArg       interface{} `json:"hash_on_query_arg"`
	Slots                int         `json:"slots"`
	CreatedAt            int         `json:"created_at"`
	UseSrvName           bool        `json:"use_srv_name"`
	Algorithm            string      `json:"algorithm"`
	UpdatedAt            int         `json:"updated_at"`
	HashOnURICapture     interface{} `json:"hash_on_uri_capture"`
	Healthchecks         struct {
		Active struct {
			Timeout   int    `json:"timeout"`
			HTTPPath  string `json:"http_path"`
			Unhealthy struct {
				Timeouts     int   `json:"timeouts"`
				HTTPFailures int   `json:"http_failures"`
				HTTPStatuses []int `json:"http_statuses"`
				Interval     int   `json:"interval"`
				TCPFailures  int   `json:"tcp_failures"`
			} `json:"unhealthy"`
			HTTPSVerifyCertificate bool `json:"https_verify_certificate"`
			Headers                struct {
				Sdf []string `json:"sdf"`
			} `json:"headers"`
			Concurrency int    `json:"concurrency"`
			Type        string `json:"type"`
			Healthy     struct {
				HTTPStatuses []int `json:"http_statuses"`
				Interval     int   `json:"interval"`
				Successes    int   `json:"successes"`
			} `json:"healthy"`
			HTTPSSni interface{} `json:"https_sni"`
		} `json:"active"`
		Threshold int `json:"threshold"`
		Passive   struct {
			Type      string `json:"type"`
			Unhealthy struct {
				Timeouts     int   `json:"timeouts"`
				HTTPFailures int   `json:"http_failures"`
				HTTPStatuses []int `json:"http_statuses"`
				TCPFailures  int   `json:"tcp_failures"`
			} `json:"unhealthy"`
			Healthy struct {
				Successes    int   `json:"successes"`
				HTTPStatuses []int `json:"http_statuses"`
			} `json:"healthy"`
		} `json:"passive"`
	} `json:"healthchecks"`
	ClientCertificate      interface{} `json:"client_certificate"`
	HostHeader             string      `json:"host_header"`
	HashOnCookie           interface{} `json:"hash_on_cookie"`
	Name                   string      `json:"name"`
	Tags                   interface{} `json:"tags"`
	HashFallbackURICapture interface{} `json:"hash_fallback_uri_capture"`
	HashOn                 string      `json:"hash_on"`
	HashOnCookiePath       string      `json:"hash_on_cookie_path"`
	HashFallback           string      `json:"hash_fallback"`
}

type Upstreams struct {
	Next interface{} `json:"next"`
	Data []Upstream  `json:"data"`
}

// list
// 1. all upstreams 						/upstreams
// 2. all upstreams in a workspace			/{workspace}/upstreams
func ListAllUpstreams(apiEndpoint, workspace string) (*http.Response, error) {
	var pathComponents []string
	if common.IsStringSet(workspace) {
		pathComponents = append(pathComponents, workspace, "upstreams")
	} else {
		pathComponents = append(pathComponents, "upstreams")
	}
	return common.SendRequest("GET", apiEndpoint, pathComponents, nil)
}

// create
// 1. a upstream							/upstreams
// 2. a upstream in a workspace 			/{workspace}/upstreams
func (u *Upstream) CreateUpstream(apiEndpoint, workspace string) (*http.Response, error) {
	var pathComponents []string
	if common.IsStringSet(workspace) {
		pathComponents = append(pathComponents, workspace, "upstreams")
	} else {
		pathComponents = append(pathComponents, "upstreams")
	}
	return common.SendRequest("POST", apiEndpoint, pathComponents, u)
}

// delete
// 1. a upstream									/upstreams/{upstream_name}
// 2. a upstream in a workspace					/{workspace}/upstreams/{upstream_name}
func (u *Upstream) DeleteUpstream(apiEndpoint, workspace string) (*http.Response, error) {
	var pathComponents []string
	if common.IsStringSet(workspace) {
		pathComponents = append(pathComponents, workspace, "upstreams", u.Name)
	} else {
		pathComponents = append(pathComponents, "upstreams", u.Name)
	}
	return common.SendRequest("DELETE", apiEndpoint, pathComponents, nil)
}

// update
// 1. a upstream									/upstreams/{upstreamName}
// 2. a upstream in a workspace					/{workspace}/upstreams/{upstream_name}
func (u *Upstream) UpdateUpstream(apiEndpoint, workspace string) (*http.Response, error) {
	var pathComponents []string
	if common.IsStringSet(workspace) {
		pathComponents = append(pathComponents, workspace, "upstreams", u.Name)
	} else {
		pathComponents = append(pathComponents, "upstreams", u.Name)
	}
	return common.SendRequest("PATCH", apiEndpoint, pathComponents, u)
}

// batch create upstreams
func (u *Upstreams) BatchCreateUpstreams(apiEndpoint, workspace string) ([]*http.Response, []error) {
	return batchExecuteUpstreams(apiEndpoint, workspace, *u, "create")
}

// batch delete upstreams
func (s *Upstreams) BatchDeleteUpstreams(apiEndpoint, workspace string) ([]*http.Response, []error) {
	return batchExecuteUpstreams(apiEndpoint, workspace, *s, "delete")
}

// batch update upstreams
func (s *Upstreams) BatchUpdateUpstreams(apiEndpoint, workspace string) ([]*http.Response, []error) {
	return batchExecuteUpstreams(apiEndpoint, workspace, *s, "update")
}

func batchExecuteUpstreams(apiEndpoint, workspace string, upstreams Upstreams, operation string) ([]*http.Response, []error) {
	var wg sync.WaitGroup
	var responses []*http.Response
	var errs []error

	ch := make(chan struct {
		response *http.Response
		err      error
	}, len(upstreams.Data))

	for _, upstream := range upstreams.Data {
		wg.Add(1)
		go func(upstream Upstream) {
			defer wg.Done()
			var resp *http.Response
			var err error

			switch operation {
			case "create":
				resp, err = upstream.CreateUpstream(apiEndpoint, workspace)
			case "update":
				resp, err = upstream.UpdateUpstream(apiEndpoint, workspace)
			case "delete":
				resp, err = upstream.DeleteUpstream(apiEndpoint, workspace)
			default:
				err = fmt.Errorf("invalid operation type: %s", operation)
			}

			ch <- struct {
				response *http.Response
				err      error
			}{response: resp, err: err}
		}(upstream)
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
