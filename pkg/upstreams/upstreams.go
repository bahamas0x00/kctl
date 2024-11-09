package upstreams

type Upstream struct {
	Name              string            `json:"name"`
	Algorithm         string            `json:"algorithm"`
	HashOn            string            `json:"hash_on"`
	HashFallback      string            `json:"hash_fallback"`
	HashOnCookiePath  string            `json:"hash_on_cookie_path"`
	Slots             int               `json:"slots"`
	Healthchecks      Healthchecks      `json:"healthchecks"`
	Tags              []string          `json:"tags"`
	HostHeader        string            `json:"host_header"`
	ClientCertificate ClientCertificate `json:"client_certificate"`
	UseSrvName        bool              `json:"use_srv_name"`
}

type Healthchecks struct {
	Passive   PassiveHealthcheck `json:"passive"`
	Active    ActiveHealthcheck  `json:"active"`
	Threshold int                `json:"threshold"`
}

type PassiveHealthcheck struct {
	Type      string       `json:"type"`
	Healthy   HealthStatus `json:"healthy"`
	Unhealthy HealthStatus `json:"unhealthy"`
}

type ActiveHealthcheck struct {
	Type                   string                   `json:"type"`
	Concurrency            int                      `json:"concurrency"`
	Headers                Headers                  `json:"headers"`
	Timeout                int                      `json:"timeout"`
	HttpPath               string                   `json:"http_path"`
	HttpsSni               string                   `json:"https_sni"`
	HttpsVerifyCertificate bool                     `json:"https_verify_certificate"`
	Healthy                HealthStatusWithInterval `json:"healthy"`
	Unhealthy              HealthStatusWithInterval `json:"unhealthy"`
}

type HealthStatus struct {
	HttpStatuses []int `json:"http_statuses"`
	Successes    int   `json:"successes"`
	Timeouts     int   `json:"timeouts"`
	HttpFailures int   `json:"http_failures"`
	TcpFailures  int   `json:"tcp_failures"`
}

type HealthStatusWithInterval struct {
	HttpStatuses []int `json:"http_statuses"`
	Successes    int   `json:"successes"`
	Interval     int   `json:"interval"`
	Timeouts     int   `json:"timeouts"`
	HttpFailures int   `json:"http_failures"`
	TcpFailures  int   `json:"tcp_failures"`
}

type Headers struct {
	XMyHeader      []string `json:"x-my-header"`
	XAnotherHeader []string `json:"x-another-header"`
}

type ClientCertificate struct {
	ID string `json:"id"`
}
