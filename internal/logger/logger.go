package logger

import (
	"os"
	"sync"

	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugar *zap.SugaredLogger
var once sync.Once

func GetInstance() *zap.SugaredLogger {
	once.Do(
		func() {
			config := config.GetInstance()
			zapConfig := zap.NewProductionEncoderConfig()
			zapConfig.EncodeTime = zapcore.ISO8601TimeEncoder
			fileEncoder := zapcore.NewJSONEncoder(zapConfig)
			consoleEncoder := zapcore.NewConsoleEncoder(zapConfig)
			logFile, _ := os.OpenFile(config.LogFileLocation, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			writer := zapcore.AddSync(logFile)

			var defaultLogLevel zapcore.Level

			switch config.ZapDefaultLogLevel {
			case "Debug":
				defaultLogLevel = zapcore.DebugLevel
			default:
				defaultLogLevel = zapcore.WarnLevel
			}

			core := zapcore.NewTee(
				zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
				zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
			)
			sugar = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()
			sugar.Debug("Zap Logger Initialized")
		},
	)

	return sugar
}
