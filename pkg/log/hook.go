package log

import (
	"fmt"
	"os"

	"github.com/ZC-A/notesBE/pkg/eventbus"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// setDefaultConfig
func setDefaultConfig() {
	viper.SetDefault(LevelConfigPath, "info")
	viper.SetDefault(PathConfigPath, "")
}

// 从string到AtomicLevel的转换
func setLogLevel(level string) {
	switch level {
	case "debug":
		loggerLevel.SetLevel(zap.DebugLevel)
	case "info":
		loggerLevel.SetLevel(zap.InfoLevel)
	case "warning":
		loggerLevel.SetLevel(zap.WarnLevel)
	case "error":
		loggerLevel.SetLevel(zap.ErrorLevel)
	case "fatal":
		loggerLevel.SetLevel(zap.FatalLevel)
	default:
		loggerLevel.SetLevel(zap.InfoLevel)
	}
}

// 初始化日志配置
func initLogConfig() {
	var (
		encoder zapcore.Encoder
		err     error
	)

	// 配置日志级别
	setLogLevel(viper.GetString(LevelConfigPath))

	// 日志路径及轮转配置
	var writeSyncer zapcore.WriteSyncer
	if viper.GetString(PathConfigPath) == "" {
		writeSyncer = zapcore.Lock(os.Stdout)
	} else {
		if syncer, err = NewReopenableWriteSyncer(viper.GetString(PathConfigPath)); err != nil {
			fmt.Printf("failed to create syncer for->[%s]", err)
			return
		}
		writeSyncer = syncer
	}
	// 配置日志格式
	encoder = zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})

	DefaultLogger = &Logger{
		logger: zap.New(
			zapcore.NewCore(encoder, writeSyncer, loggerLevel),
			zap.AddCaller(), zap.AddCallerSkip(2),
		),
	}
}

// init
func init() {
	if err := eventbus.EventBus.Subscribe(eventbus.EventSignalConfigPreParse, setDefaultConfig); err != nil {
		fmt.Printf(
			"failed to subscribe event->[%s] for log module for default config, maybe log module won't working.",
			eventbus.EventSignalConfigPreParse,
		)
	}

	if err := eventbus.EventBus.Subscribe(eventbus.EventSignalConfigPostParse, initLogConfig); err != nil {
		fmt.Printf(
			"failed to subscribe event->[%s] for log module for new config, maybe log module won't working.",
			eventbus.EventSignalConfigPostParse,
		)
	}
}
