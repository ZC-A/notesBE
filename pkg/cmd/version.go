package cmd

import (
	"fmt"

	"github.com/ZC-A/notesBE/pkg/config"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version ",
	Long:  `show version and exit`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s", config.Version)
		if cmd.Flag("detail").Changed {
			fmt.Printf("-%s", config.CommitHash)
		}
		fmt.Printf("\n")
	},
}

// init 加载默认配置
func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	versionCmd.Flags().BoolP("detail", "d", false, "show detail version info")
}
