package trace

import (
	"context"
	"fmt"

	"github.com/ZC-A/notesBE/pkg/eventbus"
	"github.com/ZC-A/notesBE/pkg/log"
	"github.com/spf13/viper"
)

func setDefaultConfig() {
	viper.SetDefault(KeysConfigPath, nil)
	viper.SetDefault(OtlpHostConfigPath, "127.0.0.1")
	viper.SetDefault(OtlpPortConfigPath, "4317")
	viper.SetDefault(OtlpTokenConfigPath, "")
	viper.SetDefault(OtlpTypeConfigPath, "grpc")

	viper.SetDefault(ServiceNameConfigPath, "notesBE")
	viper.SetDefault(EnableConfigPath, true)
}

// InitConfig
func InitConfig() {

	Enable = viper.GetBool(EnableConfigPath)

	for key, value := range configLabels {
		log.Debugf(context.TODO(), "key->[%s] value->[%s] now is added to labels", key, value)
		labels[key] = value
	}

	otlpHost = viper.GetString(OtlpHostConfigPath)
	otlpPort = viper.GetString(OtlpPortConfigPath)
	otlpToken = viper.GetString(OtlpTokenConfigPath)
	log.Infof(context.TODO(), "trace will Otlp to host->[%s] port->[%s] token->[%s]", otlpHost, otlpPort, otlpToken)

	OtlpType = viper.GetString(OtlpTypeConfigPath)
	log.Infof(context.TODO(), "trace will Otlp as %s type", OtlpType)

	ServiceName = viper.GetString(ServiceNameConfigPath)
	log.Infof(context.TODO(), "trace will Otlp service name:%s", ServiceName)
}

// init
func init() {
	if err := eventbus.EventBus.Subscribe(eventbus.EventSignalConfigPreParse, setDefaultConfig); err != nil {
		fmt.Printf(
			"failed to subscribe event->[%s] for trace module for default config, maybe http module won't working.",
			eventbus.EventSignalConfigPreParse,
		)
	}

	if err := eventbus.EventBus.Subscribe(eventbus.EventSignalConfigPostParse, InitConfig); err != nil {
		fmt.Printf(
			"failed to subscribe event->[%s] for trace module for new config, maybe http module won't working.",
			eventbus.EventSignalConfigPostParse,
		)
	}
}
