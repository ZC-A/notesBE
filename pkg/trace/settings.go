package trace

const (
	KeysConfigPath      = "trace.labels"
	OtlpHostConfigPath  = "trace.otlp.host"
	OtlpPortConfigPath  = "trace.otlp.port"
	OtlpTokenConfigPath = "trace.otlp.token"
	// OtlpTypeConfigPath 上报模式，http，grpc
	OtlpTypeConfigPath = "trace.otlp.type"

	ServiceNameConfigPath = "trace.service_name"
	EnableConfigPath      = "trace.enable"
)

var (
	OtlpType string

	otlpHost, otlpPort, otlpToken string
	configLabels                  map[string]string

	// 监控相关内容
	labels map[string]string

	ServiceName string
	Enable      bool
)
