package log

import (
	"go.uber.org/zap"
)

const (
	LevelConfigPath = "logger.level"
	PathConfigPath  = "logger.path"
)

var (
	DefaultLogger *Logger
	loggerLevel   = zap.NewAtomicLevel()
	syncer        *ReopenableWriteSyncer
)
