package logger

import (
	"fmt"
	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

var defaultLogger *log.Logger

func newStandardLogger() *log.Logger {
	formatter := new(log.TextFormatter)

	formatter.DisableColors = false
	formatter.DisableTimestamp = true
	formatter.FullTimestamp = false

	standardLogger := log.New()
	standardLogger.SetOutput(os.Stdout)
	standardLogger.SetLevel(log.InfoLevel)
	standardLogger.SetFormatter(formatter)

	if viper.GetBool("debug") == true {
		standardLogger.SetLevel(log.DebugLevel)
	}

	if viper.GetBool("no-tty") == true {
		formatter := new(log.TextFormatter)

		formatter.DisableColors = true
		formatter.DisableTimestamp = false

		standardLogger.SetFormatter(formatter)
	}

	return standardLogger
}

func GetDefaultLogger() *log.Logger {
	if defaultLogger == nil {
		defaultLogger = newStandardLogger()
	}

	if viper.GetBool("debug") == true {
		defaultLogger.SetLevel(log.DebugLevel)
	}

	if viper.GetBool("no-tty") == true {
		formatter := new(log.TextFormatter)

		formatter.PadLevelText = true
		formatter.DisableColors = false
		formatter.DisableTimestamp = false
		formatter.FullTimestamp = true

		defaultLogger.SetFormatter(formatter)
	}

	return defaultLogger
}

func Debug(args ...interface{}) {
	if viper.GetBool("no-tty") == false {
		deb := color.New(color.FgHiBlack)
		GetDefaultLogger().Debug(deb.Sprint(args...))
	} else {
		GetDefaultLogger().Debug(args...)
	}
}

func Info(args ...interface{}) {
	GetDefaultLogger().Info(args...)
}

func Success(args ...interface{}) {
	if viper.GetBool("no-tty") == false {
		success := color.New(color.FgGreen)
		GetDefaultLogger().Infof(success.Sprintf("✔ %s ", fmt.Sprint(args...)))
	} else {
		GetDefaultLogger().Infof("✔ %s ", fmt.Sprint(args...))
	}
}

func Warning(args ...interface{}) {
	if viper.GetBool("no-tty") == false {
		warn := color.New(color.FgYellow)
		GetDefaultLogger().Warn(warn.Sprintf("⚠ %s ", fmt.Sprint(args...)))
	} else {
		GetDefaultLogger().Warnf("⚠ %s ", fmt.Sprint(args...))
	}
}

func Error(args ...interface{}) {
	if viper.GetBool("no-tty") == false {
		err := color.New(color.FgRed)
		GetDefaultLogger().Error(err.Sprintf("✗ %s ", fmt.Sprint(args...)))
	} else {
		GetDefaultLogger().Errorf("✗ %s", fmt.Sprint(args...))
	}
}

func Fatal(args ...interface{}) {
	if viper.GetBool("no-tty") == false {
		err := color.New(color.FgHiWhite, color.BgRed)
		GetDefaultLogger().Fatal(err.Sprintf("✗ %s ", fmt.Sprint(args...)))
	} else {
		GetDefaultLogger().Fatalf("✗ %s ", fmt.Sprint(args...))
	}
}

func Exit(code int) {
	GetDefaultLogger().Exit(code)
}

func GetLogWithContext(context log.Fields) *log.Entry {
	contextLogger := GetDefaultLogger().WithFields(context)
	return contextLogger
}
