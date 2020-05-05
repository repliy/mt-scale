package syslog

import (
	"fmt"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

var logger = log.New()

// Fields logrus.Fields
type Fields log.Fields

func init() {
	logger.Formatter = &log.JSONFormatter{}
	logger.Level = log.DebugLevel
}

// Debug args ...interface{}
func Debug(args ...interface{}) {
	if logger.Level >= log.DebugLevel {
		entry := logger.WithFields(log.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Debug(args)
	}
}

// DebugWithFields l interface{}, f Fields
func DebugWithFields(l interface{}, f Fields) {
	if logger.Level >= log.DebugLevel {
		entry := logger.WithFields(log.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Debug(l)
	}
}

// Info args ...interface{}
func Info(args ...interface{}) {
	if logger.Level >= log.InfoLevel {
		entry := logger.WithFields(log.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Info(args...)
	}
}

// InfoWithFields l interface{}, f Fields
func InfoWithFields(l interface{}, f Fields) {
	if logger.Level >= log.InfoLevel {
		entry := logger.WithFields(log.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Info(l)
	}
}

// Warn args ...interface{}
func Warn(args ...interface{}) {
	if logger.Level >= log.WarnLevel {
		entry := logger.WithFields(log.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Warn(args...)
	}
}

// WarnWithFields l interface{}, f Fields
func WarnWithFields(l interface{}, f Fields) {
	if logger.Level >= log.WarnLevel {
		entry := logger.WithFields(log.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Warn(l)
	}
}

// Error args ...interface{}
func Error(args ...interface{}) {
	if logger.Level >= log.ErrorLevel {
		entry := logger.WithFields(log.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Error(args...)
	}
}

// ErrorWithFields l interface{}, f Fields
func ErrorWithFields(l interface{}, f Fields) {
	if logger.Level >= log.ErrorLevel {
		entry := logger.WithFields(log.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Error(l)
	}
}

// Fatal args ...interface{}
func Fatal(args ...interface{}) {
	if logger.Level >= log.FatalLevel {
		entry := logger.WithFields(log.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Fatal(args...)
	}
}

// FatalWithFields l interface{}, f Fields
func FatalWithFields(l interface{}, f Fields) {
	if logger.Level >= log.FatalLevel {
		entry := logger.WithFields(log.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Fatal(l)
	}
}

// Panic args ...interface{}
func Panic(args ...interface{}) {
	if logger.Level >= log.PanicLevel {
		entry := logger.WithFields(log.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Panic(args...)
	}
}

// PanicWithFields l interface{}, f Fields
func PanicWithFields(l interface{}, f Fields) {
	if logger.Level >= log.PanicLevel {
		entry := logger.WithFields(log.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Panic(l)
	}
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}
