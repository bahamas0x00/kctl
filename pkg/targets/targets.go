package targets

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

// create(POST)

// delete(DELETE)

// update (PATCH)

// batch create

// batch delete

// batch update
