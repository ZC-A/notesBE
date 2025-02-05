package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ZC-A/notesBE/pkg/config"
	"github.com/ZC-A/notesBE/pkg/http"
	"github.com/ZC-A/notesBE/pkg/log"
	"github.com/ZC-A/notesBE/pkg/service"
	"github.com/ZC-A/notesBE/pkg/trace"
	"github.com/google/gops/agent"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "run",
	Short: "start notesBE module",
	Long:  `start notesBE module`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			serviceList     []service.Service
			ctx, cancelFunc = context.WithCancel(context.Background())
			sc              = make(chan os.Signal, 1)
		)
		config.InitConfig()
		id := uuid.New()

		type contextKey string
		const uuidKey contextKey = "uuid"
		ctx = context.WithValue(ctx, uuidKey, id.String())

		// 启动 gops
		if err := agent.Listen(agent.Options{}); err != nil {
			log.Warnf(ctx, "%s", err.Error())
		}

		// 初始化启动任务
		serviceList = []service.Service{
			&trace.Service{},
			&http.Service{},
		}
		log.Infof(ctx, "services started.")

		// 注册信号（重载配置文件 & 停止）
		signal.Notify(sc, syscall.SIGUSR1, syscall.SIGTERM, syscall.SIGINT)
	LOOP:
		for {
			for _, service := range serviceList {
				service.Reload(ctx)
			}
			log.Infof(ctx, "reload done")
			switch <-sc {
			case syscall.SIGUSR1:
				// 触发配置重载动作
				config.InitConfig()
				log.Debugf(ctx, "SIGUSR1 signal got, will reload server")
			case syscall.SIGTERM, syscall.SIGINT:
				log.Debugf(ctx, "shutdown signal got, will shutdown server")
				log.Warnf(ctx, "shutdown signal process done")
				break LOOP
			}
		}
		log.Debugf(ctx, "loop break, wait for all service exit.")
		for i := len(serviceList) - 1; i >= 0; i-- {
			log.Warnf(ctx, "close service:%s", serviceList[i].Type())
			serviceList[i].Close()
			log.Warnf(ctx, "waiting for service:%s", serviceList[i].Type())

			serviceList[i].Wait()
			log.Warnf(ctx, "waiting for service:%s done", serviceList[i].Type())
		}
		cancelFunc()
		log.Debugf(ctx, "all service exit, server exit now.")
		os.Exit(0)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// init 加载默认配置
func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(
		&config.CustomConfigFilePath, "config", "", "config file (default is ./config.yaml)",
	)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
