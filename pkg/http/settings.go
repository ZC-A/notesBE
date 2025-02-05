package http

import "time"

const (
	IPAddressConfigPath           = "http.address"
	PortConfigPath                = "http.port"
	WriteTimeOutConfigPath        = "http.write_timeout"
	ReadTimeOutConfigPath         = "http.read_timeout"
	SingleflightTimeoutConfigPath = "http.singleflight_timeout"
	SlowQueryThresholdConfigPath  = "http.slow_query_threshold"
	HelloWorldPathConfigPath      = "http.path.hello_world"
)

var (
	IPAddress string
	Port      int

	WriteTimeout        time.Duration
	ReadTimeout         time.Duration
	SingleflightTimeout time.Duration
	SlowQueryThreshold  time.Duration
)
