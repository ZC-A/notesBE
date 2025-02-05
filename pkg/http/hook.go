package http

import (
	"context"
	"fmt"

	"github.com/ZC-A/notesBE/pkg/eventbus"
	"github.com/ZC-A/notesBE/pkg/log"
	"github.com/spf13/viper"
)

func setDefaultConfig() {
	viper.SetDefault(IPAddressConfigPath, "127.0.0.1")
	viper.SetDefault(PortConfigPath, 10205)
	viper.SetDefault(WriteTimeOutConfigPath, "30s")
	viper.SetDefault(ReadTimeOutConfigPath, "3s")
	viper.SetDefault(SlowQueryThresholdConfigPath, "3s")
	viper.SetDefault(SingleflightTimeoutConfigPath, "1m")

}

// LoadConfig
func LoadConfig() {

	IPAddress = viper.GetString(IPAddressConfigPath)
	Port = viper.GetInt(PortConfigPath)
	WriteTimeout = viper.GetDuration(WriteTimeOutConfigPath)
	ReadTimeout = viper.GetDuration(ReadTimeOutConfigPath)
	SingleflightTimeout = viper.GetDuration(SingleflightTimeoutConfigPath)
	SlowQueryThreshold = viper.GetDuration(SlowQueryThresholdConfigPath)

	log.Debugf(context.TODO(), "reload success new config address->[%s] port->[%d] "+
		"going to reload the service.",
		IPAddress, Port)
}

// init
func init() {
	if err := eventbus.EventBus.Subscribe(eventbus.EventSignalConfigPreParse, setDefaultConfig); err != nil {
		fmt.Printf(
			"failed to subscribe event->[%s] for http module for default config, maybe http module won't working.",
			eventbus.EventSignalConfigPreParse,
		)
	}

	if err := eventbus.EventBus.Subscribe(eventbus.EventSignalConfigPostParse, LoadConfig); err != nil {
		fmt.Printf(
			"failed to subscribe event->[%s] for http module for new config, maybe http module won't working.",
			eventbus.EventSignalConfigPostParse,
		)
	}
}
