package cmd

import (
	"fmt"

	"github.com/ZC-A/notesBE/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "show current config",
	Run: func(cmd *cobra.Command, args []string) {
		config.InitConfig()

		output, err := yaml.Marshal(viper.AllSettings())
		if err != nil {
			fmt.Printf("failed to marshal config for->[%s]", err)
			return
		}
		fmt.Printf("%s", output)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
