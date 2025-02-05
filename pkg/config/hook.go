package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/ZC-A/notesBE/pkg/eventbus"
	"github.com/spf13/viper"
)

// InitConfig 初始化配置
func InitConfig() {
	if CustomConfigFilePath != "" {
		// Use config file from the flag.
		viper.SetConfigFile(CustomConfigFilePath)
	} else {
		fmt.Println(111)
		// Find home directory.
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(cwd)
		viper.SetConfigName(AppName)
		viper.SetConfigType("yaml")
	}

	viper.SetEnvPrefix(AppName)
	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// If a config file is found, read it in.
	// 在配置文件读取前，需要先通知全世界做好准备
	eventbus.EventBus.Publish(eventbus.EventSignalConfigPreParse)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("loading config file: %s failed,error:%s\n", viper.ConfigFileUsed(), err)
	} else {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	// 配置读取后，通知全世界reload读取新的配置
	eventbus.EventBus.Publish(eventbus.EventSignalConfigPostParse)
}
