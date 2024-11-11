package targets

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/bahamas0x00/kctl/pkg/common"
)

type Target struct {
	ID        string  `json:"id"`
	UpdatedAt float64 `json:"updated_at"`
	Weight    int     `json:"weight"`
	CreatedAt float64 `json:"created_at"`
	Upstream  struct {
		ID string `json:"id"`
	} `json:"upstream"`
	Tags   []interface{} `json:"tags"`
	Target string        `json:"target"`
}

type Targets struct {
	Data []Target    `json:"data"`
	Next interface{} `json:"next"`
}

// list(GET)
// 1. all targets associated with an upstream in a workspace  		/{workspace}/upstreams/{upstream_name}/targets
// 2. all targets associated with an upstream 						/upstreams/{upstream_name}/targets
func ListAllTargets(apiEndpoint, workspace string, upstreamName string) (*http.Response, error) {
	var pathComponents []string
	if !common.IsStringSet(upstreamName) {
		return nil, fmt.Errorf("must set associated upstream name")
	}

	if common.IsStringSet(workspace) {
		pathComponents = append(pathComponents, workspace, "upstreams", upstreamName, "targets")
	} else {
		pathComponents = append(pathComponents, "upstreams", upstreamName, "targets")
	}
	return common.SendRequest("GET", apiEndpoint, pathComponents, nil)
}

// create(POST)
// 1. all targets associated with an upstream in a workspace  		/{workspace}/upstreams/{upstream_name}/targets
// 2. all targets associated with an upstream 						/upstreams/{upstream_name}/targets
func (t *Target) CreateTarget(apiEndpoint, workspace, upstreamName string) (*http.Response, error) {
	var pathComponents []string
	if !common.IsStringSet(upstreamName) {
		return nil, fmt.Errorf("must set associated upstream name")
	}

	if common.IsStringSet(workspace) {
		pathComponents = append(pathComponents, workspace, "upstreams", upstreamName, "targets")
	} else {
		pathComponents = append(pathComponents, "upstreams", upstreamName, "targets")
	}
	return common.SendRequest("POST", apiEndpoint, pathComponents, t)
}

// delete(DELETE)
// 1. all targets associated with an upstream in a workspace  		/{workspace}/upstreams/{upstream_name}/targets/{target}
// 2. all targets associated with an upstream 						/upstreams/{upstream_name}/targets/{target}
func (t *Target) DeleteTarget(apiEndpoint, workspace, upstreamName string) (*http.Response, error) {
	var pathComponents []string
	if !common.IsStringSet(upstreamName) {
		return nil, fmt.Errorf("must set associated upstream name")
	}

	if common.IsStringSet(workspace) {
		pathComponents = append(pathComponents, workspace, "upstreams", upstreamName, "targets", t.Target)
	} else {
		pathComponents = append(pathComponents, "upstreams", upstreamName, "targets", t.Target)
	}
	return common.SendRequest("DELETE", apiEndpoint, pathComponents, nil)
}

// update (PATCH)
// 1. all targets associated with an upstream in a workspace  		/{workspace}/upstreams/{upstream_name}/targets/{target_name}
// 2. all targets associated with an upstream 						/upstreams/{upstream_name}/targets/{target_name}
func (t *Target) UpdateTarget(apiEndpoint, workspace, upstreamName string) (*http.Response, error) {
	var pathComponents []string
	if !common.IsStringSet(upstreamName) {
		return nil, fmt.Errorf("must set associated upstream name")
	}

	if common.IsStringSet(workspace) {
		pathComponents = append(pathComponents, workspace, "upstreams", upstreamName, "targets", t.Target)
	} else {
		pathComponents = append(pathComponents, "upstreams", upstreamName, "targets", t.Target)
	}
	return common.SendRequest("PATCH", apiEndpoint, pathComponents, t)
}

// batch create
func (t *Targets) BatchCreateTargets(apiEndpoint, workspace, upstreamName string) ([]*http.Response, []error) {
	return batchExecuteTargets(apiEndpoint, workspace, upstreamName, *t, "create")
}

// batch delete
func (t *Targets) BatchDeleteTargets(apiEndpoint, workspace, upstreamName string) ([]*http.Response, []error) {
	return batchExecuteTargets(apiEndpoint, workspace, upstreamName, *t, "delete")
}

// batch update
func (t *Targets) BatchUpdateTargets(apiEndpoint, workspace, upstreamName string) ([]*http.Response, []error) {
	return batchExecuteTargets(apiEndpoint, workspace, upstreamName, *t, "update")
}

func batchExecuteTargets(apiEndpoint string, workspace string, upstreamName string, targets Targets, operation string) ([]*http.Response, []error) {
	var wg sync.WaitGroup
	var responses []*http.Response
	var errs []error

	ch := make(chan struct {
		response *http.Response
		err      error
	}, len(targets.Data))

	for _, target := range targets.Data {
		wg.Add(1)
		go func(target Target) {
			defer wg.Done()
			var resp *http.Response
			var err error

			switch operation {
			case "create":
				resp, err = target.CreateTarget(apiEndpoint, workspace, upstreamName)
			case "update":
				resp, err = target.UpdateTarget(apiEndpoint, workspace, upstreamName)
			case "delete":
				resp, err = target.DeleteTarget(apiEndpoint, workspace, upstreamName)
			default:
				err = fmt.Errorf("invalid operation type: %s", operation)
			}

			ch <- struct {
				response *http.Response
				err      error
			}{response: resp, err: err}
		}(target)
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
