package syslog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

func initLogger(logpath string, loglevel string) *zap.Logger {
	hook := lumberjack.Logger{
		Filename:   logpath, // ⽇志⽂件路径
		MaxSize:    1,       // megabytes
		MaxBackups: 3,       // 最多保留3个备份
		MaxAge:     7,       // days
		Compress:   true,    // 是否压缩 disabled by default
	}
	w := zapcore.AddSync(&hook)
	var level zapcore.Level
	switch loglevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		w,
		level,
	)
	logger := zap.New(core)
	logger.Info(loglevel + "init success")
	return logger
}

var (
	debugLogger *zap.Logger
	infoLogger  *zap.Logger
	errorLogger *zap.Logger
)

func init() {
	debugLogger = initLogger("/log/debug.log", "debug")
	infoLogger = initLogger("/log/info.log", "info")
	errorLogger = initLogger("/log/error.log", "error")
}

// Debug args ...interface{}
func Debug(info string, arg ...interface{}) {
	debugLogger.Debug(info, zap.Any("key", arg))
}

// Info args ...interface{}
func Info(info string, arg ...interface{}) {
	infoLogger.Info(info, zap.Any("key", arg))
}

// Error args ...interface{}
func Error(err error, arg ...interface{}) {
	errorLogger.Error(err.Error(), zap.Any("key", arg))
}
